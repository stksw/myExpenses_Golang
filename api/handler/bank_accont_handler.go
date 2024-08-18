package handler

import (
	"encoding/json"
	"myExpenses/database"
	"myExpenses/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type BankAccountHandler interface {
	ListBankAccount(w http.ResponseWriter, r *http.Request)
	CreateBankAccount(w http.ResponseWriter, r *http.Request)
	UpdateBankAccount(w http.ResponseWriter, r *http.Request)
}

type bankAccountHandler struct {
	Handler
}

func NewBankAccountHandler(db gorm.DB) *bankAccountHandler {
	return &bankAccountHandler{}
}

// CreateBankAccount creates a new bank account
func (h *bankAccountHandler) CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	var account models.BankAccount
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&account).Error; err != nil {
		http.Error(w, "Failed to create bank account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// GetBankAccount retrieves a bank account by ID
func (h *bankAccountHandler) GetBankAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var account models.BankAccount

	if err := database.DB.First(&account, id).Error; err != nil {
		http.Error(w, "Bank account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// UpdateBankAccount updates a bank account by ID
func (h *bankAccountHandler) UpdateBankAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var account models.BankAccount

	if err := database.DB.First(&account, id).Error; err != nil {
		http.Error(w, "Bank account not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&account).Error; err != nil {
		http.Error(w, "Failed to update bank account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// DeleteBankAccount deletes a bank account by ID
func (h *bankAccountHandler) DeleteBankAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var account models.BankAccount

	if err := database.DB.First(&account, id).Error; err != nil {
		http.Error(w, "Bank account not found", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&account).Error; err != nil {
		http.Error(w, "Failed to delete bank account", http.StatusInternalServerError)
		return
	}

	h.RenderResponse(w, r, http.StatusNoContent)
}
