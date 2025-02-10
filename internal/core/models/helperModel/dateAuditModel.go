package helperModel

import (
	"time"

	"gorm.io/gorm"
)

type DateAuditModel struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at is null")
}
