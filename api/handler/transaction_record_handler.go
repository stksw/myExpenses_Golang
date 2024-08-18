package handler

import (
	"fmt"
	"myExpenses/database"
	"myExpenses/models"
	"net/http"

	"gorm.io/gorm"
)

type TransactionRecordHandler interface {
	ListTransactionRecord(w http.ResponseWriter, r *http.Request)
	CreateTransactionRecord(w http.ResponseWriter, r *http.Request)
	UpdateTransactionRecord(w http.ResponseWriter, r *http.Request)
}

type transactionRecordHandler struct {
	Handler
}

func NewTransactionRecordHandler(db gorm.DB) *transactionRecordHandler {
	return &transactionRecordHandler{}
}

func (h *transactionRecordHandler) ListTransactionRecord(w http.ResponseWriter, r *http.Request) {
	var tr []models.TransactionRecord
	result := database.DB.Find(&tr)
	if result.Error != nil {
		fmt.Println("Error retrieving records:", result.Error)
		return
	}
	h.RenderResponse(w, r, tr)

}

func (h *transactionRecordHandler) CreateTransactionRecord(w http.ResponseWriter, r *http.Request) {
	tr := models.TransactionRecord{
		Type:     "Deposit",
		Recorder: "LegalOn Technologies",
		Amount:   5000,
	}
	if err := database.DB.Create(&tr).Error; err != nil {
		fmt.Println("Error insert records:", err)
	}
}

func (h *transactionRecordHandler) UpdateTransactionRecord(w http.ResponseWriter, r *http.Request) {

}
