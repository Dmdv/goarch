package database

import (
	"github.com/goarch/gosvc/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB ... Migrates db
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}

	return nil
}
