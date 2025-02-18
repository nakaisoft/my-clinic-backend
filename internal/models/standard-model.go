package models

import "time"

// StandardModel contém campos comuns para todas as entidades.
type StandardModel struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
