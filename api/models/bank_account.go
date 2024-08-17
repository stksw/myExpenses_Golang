package models

import "gorm.io/gorm"

type BankAccount struct {
	gorm.Model
	BankName string `json:"bankName"`
	Branch   string `json:"branch"`
	Type     string `json:"type"`
	Number   string `json:"number"`
	Holder   string `json:"holder"`
}
