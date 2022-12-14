package domains

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID      uint64         `json:"id"      gorm:"->;primaryKey;autoIncrement"`
	Title   string         `json:"title"   gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Checked bool           `json:"checked" gorm:"default: false"`
	Created time.Time      `json:"created" gorm:"autoCreateTime"`
	Updated time.Time      `json:"updated" gorm:"autoUpdateTime"`
	Deleted gorm.DeletedAt `json:"-"       gorm:"index"`
}
