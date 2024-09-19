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

type DebtService struct {
   debtRepository *repositories.DebtRepository
}

func NewDebtService(debtRepo *repositories.DebtRepository) *DebtService {
	return &DebtService{debtRepository: debtRepo}
}

func (s *DebtService) RegisterHTTPEndpoints(r chi.Router) {
    r.Get("/debts", s.ListDebts)
    r.Post("/debts", s.CreateDebt)
    r.Get("/debts/{id}", s.GetById)
    r.Put("/debts/{id}", s.UpdateRecord)
    r.Delete("/debts/{id}", s.DeleteRecord)
}

func (s *DebtService) ListDebts(w http.ResponseWriter, r *http.Request) {
	debts, err := s.debtRepository.ListDebts(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(debts)
	return
}

func (s *DebtService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
	}
	debt, err := s.debtRepository.GetByID(context.Background(), parsed_id)
	if err != nil {
		if debt.Name == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(debt)
	return
}

func (s *DebtService) CreateDebt(w http.ResponseWriter, r *http.Request) {
	var debt models.Debt
	err := json.NewDecoder(r.Body).Decode(&debt)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}


	params := schemas.CreateDebtParams{
		Name: debt.Name,
		Lender: debt.Lender,
		Borrower: debt.Borrower,
        InterestRate: debt.InterestRate,
        BorrowedAmount: debt.BorrowedAmount,
        PaidAmount: debt.PaidAmount,
        LendDate: debt.LendDate,
	}


	err = s.debtRepository.CreateDebt(context.Background(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	msg := "Debt created successfully"
	utils.WriteResponse(w, http.StatusOK, msg)
	return
}

func (s *DebtService) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateDebtParams
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

    if params.Lender != "" {
        params.UpdateLender = true
    }
    if params.Name != "" {
        params.UpdateName = true
    }
    if params.Borrower != "" {
        params.UpdateBorrower = true
    }
    if params.LendDate != nilTimeValue {
        params.UpdateLendDate = true
    }
    if params.PaidAmount != 0 {
        params.UpdatePaidAmount = true
    }
    if params.BorrowedAmount != 0 {
        params.UpdateBorrowedAmount = true
    }
    if params.InterestRate != 0 {
        params.UpdateInterestRate = true
    }

    err = s.debtRepository.PartialUpdateDebt(context.Background(), params)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Debt updated successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

func (s *DebtService) DeleteRecord(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    parsed_id, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
        return
    }

    err = s.debtRepository.DeleteDebt(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Debt deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}
