package service

import (
	"encoding/json"
	"github.com/syumai/workers/_templates/cloudflare/deploy-go/database"
	"github.com/syumai/workers/_templates/cloudflare/deploy-go/model"
	"net/http"
	"strconv"
)

// LocationService 位置服务
type LocationService struct {
}

// NewLocationService 创建位置服务实例
func NewLocationService() *LocationService {
	return &LocationService{}
}

// SetLocation 获取位置信息
func (s *LocationService) SetLocation(w http.ResponseWriter, r *http.Request) {
	// 解析请求体中的位置信息
	//解析json
	var locationRecord model.LocationRecord
	if err := json.NewDecoder(r.Body).Decode(&locationRecord); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := database.InsertLocationData(locationRecord)
	if err != nil {
		http.Error(w, "Failed to save location data:"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Location data saved successfully"))
}

// GetLocationByDevIDAndTime  根据设备ID和时间范围获取位置信息
func (s *LocationService) GetLocationByDevIDAndTime(w http.ResponseWriter, r *http.Request) {
	devID := r.URL.Query().Get("dev_id")
	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")

	//时间转int64时间戳
	startTimeInt64, err := strconv.ParseInt(startTime, 10, 64)
	if err != nil {
		http.Error(w, "Invalid start_time format", http.StatusBadRequest)
		return
	}
	endTimeInt64, err := strconv.ParseInt(endTime, 10, 64)
	if err != nil {
		http.Error(w, "Invalid end_time format", http.StatusBadRequest)
		return
	}

	if devID == "" || startTime == "" || endTime == "" {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	records, err := database.GetLocationByDevIdAndTime(devID, startTimeInt64, endTimeInt64)
	if err != nil {
		http.Error(w, "Failed to retrieve location data:"+err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(records)
	if err != nil {
		http.Error(w, "Failed to marshal response data:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
