package model

type LocationRecord struct {
	ID             int     `json:"id"`
	DeviceID       string  `json:"device_id"`
	Timestamp      int64   `json:"timestamp"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Altitude       float64 `json:"altitude,omitempty"`
	SignalStrength int     `json:"signal_strength,omitempty"`
	Speed          float64 `json:"speed,omitempty"`
	Satellites     int     `json:"satellites,omitempty"`
	CreatedAt      int64   `json:"created_at"`
}
