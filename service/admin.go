package service

import (
	"encoding/json"
	"net/http"
)

// AdminService 管理服务
type AdminService struct{}

// NewAdminService 创建管理服务实例
func NewAdminService() *AdminService {
	return &AdminService{}
}

// HandleAdmin 管理员处理器
func (s *AdminService) HandleAdmin(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Admin panel",
		"status":  "authenticated",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
