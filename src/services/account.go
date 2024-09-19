package services

import (
	"context"
	"log"
	"strconv"

	"dbf_api/models"
	"dbf_api/repositories"
	"dbf_api/schemas"
	"dbf_api/utils"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AccountService struct {
   accountRepository *repositories.AccountRepository
}

func NewAccountService(transRepo *repositories.AccountRepository) *AccountService {
	return &AccountService{accountRepository: transRepo}
}

func (s *AccountService) RegisterHTTPEndpoints(r chi.Router) {
    r.Get("/accounts", s.ListAccounts)
    r.Post("/accounts", s.CreateAccount)
    r.Get("/accounts/{id}", s.GetById)
    r.Put("/accounts/{id}", s.UpdateRecord)
    r.Delete("/accounts/{id}", s.DeleteRecord)
}

func (s *AccountService) ListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := s.accountRepository.ListAccounts(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(accounts)
	return
}

func (s *AccountService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
	}
	account, err := s.accountRepository.GetByID(context.Background(), parsed_id)
	if err != nil {
		if account.Name == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(account)
	return
}

func (s *AccountService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}

	params := schemas.CreateAccountParams{
		Name: account.Name,
		AccBalance: account.AccBalance,
		AccNum: account.AccNum,
		CardNum: account.CardNum,
		Pin: account.Pin,
		SecurityCode: account.SecurityCode,
		CreditLimit: account.CreditLimit,
        MethodID: account.MethodID,
	}

	err = s.accountRepository.CreateAccount(context.Background(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	msg := "Account created successfully"
	utils.WriteResponse(w, http.StatusOK, msg)
	return
}

func (s *AccountService) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateAccountParams
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
    }

    err = json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        http.Error(w, "Unexpeced error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    params.ID = parsed_id

    if params.Name != "" {
        params.UpdateName = true
    }
    if params.AccBalance != 0 {
        params.UpdateAccBalance = true
    }
    if params.AccNum != "" {
        params.UpdateAccNum = true
    }
    if params.CardNum != "" {
        params.UpdateCardNum = true
    }
    if params.Pin != "" {
        params.UpdatePin = true
    }
    if params.SecurityCode != "" {
        params.UpdateSecurityCode = true
    }
    if params.CreditLimit != 0 {
        params.UpdateCreditLimit = true
    }
    if params.MethodID != 0 {
        params.UpdateMethod = true
    }

    err = s.accountRepository.PartialUpdateAccount(context.Background(), params)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Account updated successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

func (s *AccountService) DeleteRecord(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    parsed_id, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
        return
    }

    err = s.accountRepository.DeleteAccount(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Account deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}
