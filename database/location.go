package database

import "github.com/syumai/workers/_templates/cloudflare/deploy-go/model"

func InsertLocationData(record model.LocationRecord) error {
	// 插入数据到 D1 数据库
	query := `
		INSERT INTO location_records (device_id,timestamp,latitude,longitude,altitude,signal_strength,speed,satellites,created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	_, err := locationDb.Exec(
		query,
		record.DeviceID,
		record.Timestamp,
		record.Latitude,
		record.Longitude,
		record.Altitude,
		record.SignalStrength,
		record.Speed,
		record.Satellites,
		record.CreatedAt,
	)
	return err
}
