package services

import (
	"context"
	"log"
	"strconv"
	"time"

	// "database/sql"
	"dbf_api/models"
	"dbf_api/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	queries *models.Queries
}

func NewService(queries *models.Queries) *Service {
	return &Service{queries: queries}
}

type test_struct struct {
	Test string
}

func (s *Service) RegisterHandlers() chi.Router {
	r := chi.NewRouter()

	r.Get("/", s.ListAll)
	r.Post("/", s.CreateRecord)
	r.Get("/{id}", s.GetById)
	r.Put("/{id}", s.UpdateRecord)
	r.Delete("/{id}", s.DeleteRecord)

	return r
}

func (s *Service) ListAll(w http.ResponseWriter, r *http.Request) {
	transactions, err := s.queries.ListTransactions(context.Background())
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
	transaction, err := s.queries.GetTransaction(context.Background(), parsed_id)
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

	params := models.CreateTransactionParams{
		Name: transaction.Name,
		Cost: transaction.Cost,
		Time: transaction.Time,
	}

	_, err = s.queries.CreateTransaction(context.Background(), params)
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
    var params models.PartialUpdateTransactionParams
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

    _, err = s.queries.PartialUpdateTransaction(context.Background(), params)
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

    err = s.queries.DeleteTransaction(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Transaction deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}
