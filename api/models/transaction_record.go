package models

import "gorm.io/gorm"

type TransactionRecord struct {
	gorm.Model
	Type     string `json:"type"`
	Recorder string `json:"recorder"`
	Amount   uint32 `json:"amount"`
}
