package database

import "gorm.io/gorm"

type BankAccount struct {
	gorm.Model
	BankName string
	Branch   string
	Kind     string
	Number   string
	Holder   string
}

type TransactionRecord struct {
	gorm.Model
	Kind     string
	Recorder string
	Amount   uint32
}
