package models

// PatientModel representa um paciente no sistema.
type PatientModel struct {
	Name  string `json:"name"`          // Nome completo do paciente
	Email string `json:"email"`         // Endereço de e-mail do paciente
	Cpf   string `json:"cpf,omitempty"` // CPF do paciente (opcional)
	Rg    string `json:"rg,omitempty"`  // RG do paciente (opcional)
	Phone string `json:"phone"`         // Número de telefone do paciente

	Addresses []PersonalAddress `json:"addresses,omitempty"` // Lista de endereços do paciente (opcional)

	BirthDate string `json:"birth_date,omitempty"` // Data de nascimento do paciente (opcional)
	Gender    string `json:"gender,omitempty"`     // Gênero do paciente (opcional)

	StandardModel
}
