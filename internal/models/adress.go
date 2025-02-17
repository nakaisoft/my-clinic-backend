package models

// PersonalAddress representa o endereço pessoal de um usuário.
type PersonalAddress struct {
	CEP          string `json:"cep"`         // Código postal (CEP)
	Street       string `json:"logradouro"`  // Nome da rua/avenida
	Number       string `json:"number"`      // Número da residência
	Complement   string `json:"complemento"` // Complemento (ex: apartamento, bloco)
	Neighborhood string `json:"bairro"`      // Bairro
	City         string `json:"localidade"`  // Cidade
	State        string `json:"uf"`          // Estado (UF)
	Country      string `json:"pais"`        // País

	StandardModel
}
