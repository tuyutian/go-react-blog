package models

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Model struct {
	ID        uint           `gorm:"primaryKey;authIncrement" json:"id" `
	CreatedAt time.Time      ` json:"created_at" `
	UpdatedAt time.Time      ` json:"updated_at" `
	DeletedAt gorm.DeletedAt `gorm:"index"  json:"deleted_at" `
}

func (m *Model) IDtoString() string {
	return strconv.Itoa(int(m.ID))
}
