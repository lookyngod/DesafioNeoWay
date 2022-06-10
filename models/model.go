package models

import "gorm.io/gorm"

type CPFCNPJ struct {
	gorm.Model
	CPF  string `gorm:"type:varchar(11);unique_index"`
	CNPJ string `gorm:"type:varchar(14);unique_index"`
}
