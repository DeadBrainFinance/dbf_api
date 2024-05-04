package services

import (
	"context"
	"log"
	"strconv"
	"time"

	"dbf_api/models"
	"dbf_api/repositories"
	"dbf_api/schemas"
	"dbf_api/utils"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TransactionService struct {
   transactionRepository *repositories.TransactionRepository
}

func NewTransactionService(transRepo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: transRepo}
}

func (s *TransactionService) RegisterHTTPEndpoints(r chi.Router) {
    r.Get("/transactions", s.ListTransactions)
    r.Post("/transactions", s.CreateTransaction)
    r.Get("/transactions/{id}", s.GetById)
    r.Put("/transactions/{id}", s.UpdateRecord)
    r.Delete("/transactions/{id}", s.DeleteRecord)
}

func (s *TransactionService) ListTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := s.transactionRepository.ListTransactions(context.Background())
    log.Println(transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(transactions)
	return
}

func (s *TransactionService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
	}
	transaction, err := s.transactionRepository.GetByID(context.Background(), parsed_id)
	if err != nil {
		if transaction.Name == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(transaction)
	return
}

func (s *TransactionService) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}

    log.Println(transaction)

	params := schemas.CreateTransactionParams{
		Name: transaction.Name,
		Cost: transaction.Cost,
		Time: transaction.Time,
        CategoryID: transaction.CategoryID,
	}


	err = s.transactionRepository.CreateTransaction(context.Background(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	msg := "Transaction created successfully"
	utils.WriteResponse(w, http.StatusOK, msg)
	return
}

func (s *TransactionService) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateTransactionParams
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
    }
    nilTimeValue := time.Time{}

    err = json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        http.Error(w, "Unexpeced error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    params.ID = parsed_id
    log.Println(params.CategoryID)

    if params.Cost != 0 {
        params.UpdateCost = true
    }
    if params.Name != "" {
        params.UpdateName = true
    }
    if params.Time != nilTimeValue {
        params.UpdateTime = true
    }
    if params.CategoryID != 0 {
        params.UpdateCategoryID = true
    }
    log.Println(params)

    err = s.transactionRepository.PartialUpdateTransaction(context.Background(), params)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Transaction updated successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

func (s *TransactionService) DeleteRecord(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    parsed_id, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
        return
    }

    err = s.transactionRepository.DeleteTransaction(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Transaction deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}
