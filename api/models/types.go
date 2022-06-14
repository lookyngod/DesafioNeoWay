package models

type CPFCNPJ struct {
	IsCPF   bool   //true para CPF, false para CNPJ
	CPFCNPJ string `json:"cpfcnpj,omitempty"`
}
