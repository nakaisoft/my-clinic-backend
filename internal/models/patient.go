package models

import "time"

// PatientModel representa um paciente no sistema.
type PatientModel struct {
	ID        string            `gorm:"type:string;primaryKey" json:"id"`
	Name      string            `json:"name"`
	Email     string            `gorm:"not null" json:"email"`
	Cpf       string            `gorm:"unique" json:"cpf,omitempty"`
	Rg        string            `gorm:"unique" json:"rg,omitempty"`
	Phone     string            `json:"phone"`
	BirthDate string            `json:"birth_date,omitempty"`
	Gender    string            `json:"gender,omitempty"`
	CreatedAt time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	Addresses []PersonalAddress `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"addresses,omitempty"`
}
