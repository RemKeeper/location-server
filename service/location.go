package service

import (
	"encoding/json"
	"net/http"
)

// LocationService 位置服务
type LocationService struct{}

// NewLocationService 创建位置服务实例
func NewLocationService() *LocationService {
	return &LocationService{}
}

// GetLocation 获取位置信息
func (s *LocationService) GetLocation(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Location service",
		"status":  "active",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
