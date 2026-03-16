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

// 通过ID和时间获取设备位置信息
func GetLocationByDevIdAndTime(deviceId string, startTime int64, endTime int64) ([]model.LocationRecord, error) {
	query := `
		SELECT id, device_id, timestamp, latitude, longitude, altitude, signal_strength, speed, satellites, created_at
		FROM location_records
		WHERE device_id = ? AND timestamp BETWEEN ? AND ?
		ORDER BY timestamp DESC;
	`
	rows, err := locationDb.Query(query, deviceId, startTime, endTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []model.LocationRecord
	for rows.Next() {
		var record model.LocationRecord
		err := rows.Scan(
			&record.ID,
			&record.DeviceID,
			&record.Timestamp,
			&record.Latitude,
			&record.Longitude,
			&record.Altitude,
			&record.SignalStrength,
			&record.Speed,
			&record.Satellites,
			&record.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
