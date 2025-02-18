package models

// PatientModel representa um paciente no sistema.
type PatientModel struct {
	StandardModel

	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Cpf       string            `gorm:"unique" json:"cpf,omitempty"`
	Rg        string            `json:"rg,omitempty"`
	Phone     string            `json:"phone"`
	BirthDate string            `json:"birth_date,omitempty"`
	Gender    string            `json:"gender,omitempty"`

	Addresses []PersonalAddress `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"addresses,omitempty"`
}
