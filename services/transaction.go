package services

import (
	"context"
	"log"
	"strconv"
	"time"

	// "database/sql"
	"dbf_api/models"
	"dbf_api/repositories"
	"dbf_api/schemas"
	"dbf_api/utils"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	// queries *models.Queries
   transactionRepository *repositories.TransactionRepository
}

func NewService(transRepo *repositories.TransactionRepository) *Service {
	return &Service{transactionRepository: transRepo}
}

func (s *Service) RegisterHandlers() chi.Router {
	r := chi.NewRouter()

	r.Get("/", s.ListTransactions)
	r.Post("/", s.CreateRecord)
	r.Get("/{id}", s.GetById)
	r.Put("/{id}", s.UpdateRecord)
	r.Delete("/{id}", s.DeleteRecord)

	return r
}

func (s *Service) ListTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := s.transactionRepository.ListTransactions(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(transactions)
	return
}

func (s *Service) GetById(w http.ResponseWriter, r *http.Request) {
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

func (s *Service) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}

	params := schemas.CreateTransactionParams{
		Name: transaction.Name,
		Cost: transaction.Cost,
		Time: transaction.Time,
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

func (s *Service) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateTransactionParams
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
    }
    nilTimeValue := time.Time{}
    log.Println("nil time value", nilTimeValue)

    err = json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        http.Error(w, "Unexpeced error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    params.ID = parsed_id

    if params.Cost != 0 {
        params.UpdateCost = true
    }
    if params.Name != "" {
        params.UpdateName = true
    }
    if params.Time == nilTimeValue {
        params.UpdateTime = true
    }

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

func (s *Service) DeleteRecord(w http.ResponseWriter, r *http.Request) {
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
