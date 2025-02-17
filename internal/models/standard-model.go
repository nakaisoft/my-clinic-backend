package models

type StandardModel struct {
	ID        string  `json:"id"` // ID único do paciente

	CreatedAt string `json:"created_at"` // Data de criação do registro
	UpdatedAt string `json:"updated_at"` // Data da última atualização do registro
}
