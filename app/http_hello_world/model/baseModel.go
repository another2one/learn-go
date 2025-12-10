package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// XTime 自定义时间类型
type XTime struct {
	time.Time
}

// MarshalJSON JSON序列化格式化
func (t XTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))), nil
}

// Value 数据库写入格式化
func (t XTime) Value() (driver.Value, error) {
	return t.Format("2006-01-02 15:04:05"), nil
}

// Scan 数据库读取格式化
func (t *XTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	ID        uint  `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt XTime `gorm:"autoCreateTime"`
	UpdatedAt XTime `gorm:"autoUpdateTime"`
	DeletedAt XTime `sql:"index"`
}
