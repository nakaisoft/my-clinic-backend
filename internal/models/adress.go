package models

import "time"

// PersonalAddress representa um endereço do paciente.
type PersonalAddress struct {
	ID    string `gorm:"type:string;primaryKey" json:"id"`

	PatientID    string `gorm:"index" json:"patient_id"` // Chave estrangeira para PatientModel
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`

	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}
