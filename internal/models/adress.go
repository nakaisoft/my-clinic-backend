package models

// PersonalAddress representa um endereço do paciente.
type PersonalAddress struct {
	StandardModel
	PatientID    string `gorm:"index" json:"patient_id"` // Chave estrangeira para PatientModel
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}
