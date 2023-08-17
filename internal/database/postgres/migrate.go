package postgres

import (
	"gorm.io/gorm"

	"github.com/hibiki-horimi/go-todo-api/internal/domain"
)

func Migrate(db *gorm.DB) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.AutoMigrate(&domain.Todo{}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
