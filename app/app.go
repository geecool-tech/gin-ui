package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var _db *gorm.DB

// ApiRecord 结构体定义日志模型
type ApiRecord struct {
	ID           uint           `gorm:"primaryKey"` // 主键
	InstanceName string         `gorm:"size:255"`
	Method       string         `gorm:"size:10"`  // 请求方法 (GET, POST, etc.)
	Path         string         `gorm:"size:255"` // 请求路径
	Status       int            `gorm:"index"`    // HTTP 状态码
	Duration     int64          `gorm:"index"`    // 请求响应时间 (单位：毫秒)
	IP           string         `gorm:"size:50"`  // 请求来源 IP
	UserAgent    string         `gorm:"size:255"` // 用户代理信息 (浏览器等)
	Body         string         `gorm:"size:255"` // 请求体（如果需要记录）
	Headers      string         `gorm:"size:255"` // 请求头部（如果需要记录）
	ErrorMsg     string         `gorm:"size:255"`
	CreatedAt    time.Time      // 创建时间
	UpdatedAt    time.Time      // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"index"` // 软删除时间
}

func GetDb() *gorm.DB {
	var err error
	if _db == nil {
		_db, err = gorm.Open(sqlite.Open("gin-ui-data.db"), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database:", err)
		}
		_ = _db.AutoMigrate(&ApiRecord{})
	}
	return _db

}
