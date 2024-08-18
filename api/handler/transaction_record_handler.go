package handler

import (
	"encoding/json"
	"fmt"
	"myExpenses/database"
	"myExpenses/models"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	// ctx := r.Context()
	body := models.TransactionRecord{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tr := models.TransactionRecord{
		Type:     body.Type,
		Recorder: body.Recorder,
		Amount:   body.Amount,
	}
	if err := database.DB.Create(&tr).Error; err != nil {
		fmt.Println("Error insert records:", err)
		http.Error(w, "Failed to create transaction record", http.StatusInternalServerError)
		return
	}
	h.RenderResponse(w, r, "")
}

func (h *transactionRecordHandler) UpdateTransactionRecord(w http.ResponseWriter, r *http.Request) {
	// idを取得、リクエストボディからデータをデコードする
	id := chi.URLParam(r, "id")
	body := models.TransactionRecord{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 更新対象のレコードを検索する
	var existingRecord models.TransactionRecord
	if err := database.DB.First(&existingRecord, id).Error; err != nil {
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}

	// フィールドを更新する
	existingRecord.Type = body.Type
	existingRecord.Recorder = body.Recorder
	existingRecord.Amount = body.Amount

	// レコードをデータベースに保存する
	if err := database.DB.Save(&existingRecord).Error; err != nil {
		fmt.Println("Error updating record:", err)
		http.Error(w, "Failed to update transaction record", http.StatusInternalServerError)
		return
	}

	h.RenderResponse(w, r, "Transaction record updated successfully")
}
