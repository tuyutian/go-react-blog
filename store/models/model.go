package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Model struct {
	ID        uint      `gorm:"primaryKey;authIncrement" json:"id" `
	CreatedAt time.Time ` json:"created_at" `
	UpdatedAt time.Time ` json:"updated_at" `
}

func (m *Model) IDtoString() string {
	return strconv.Itoa(int(m.ID))
}
