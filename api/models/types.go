package models

import "time"

type CPFCNPJ struct {
	IsCPF       bool      //true para CPF, false para CNPJ
	CPFCNPJ     string    `json:"cpfcnpj,omitempty"` // Formatação json para CPF/CNPJ
	Datacriacao time.Time // Timestamp da última inserção
}
