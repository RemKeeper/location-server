package database

import (
	"database/sql"
	_ "github.com/syumai/workers/cloudflare/d1"
)

var locationDb *sql.DB

func DataInit() {
	var err error
	// 连接到 D1 数据库
	locationDb, err = sql.Open("d1", "DB")
	if err != nil {
		panic("Failed to connect to D1 database: " + err.Error())
	}
}
