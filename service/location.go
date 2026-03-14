package service

import (
	"encoding/json"
	"github.com/syumai/workers/_templates/cloudflare/deploy-go/database"
	"github.com/syumai/workers/_templates/cloudflare/deploy-go/model"
	"net/http"
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
