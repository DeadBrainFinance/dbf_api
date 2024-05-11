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

type InstallmentService struct {
   installmentRepository *repositories.InstallmentRepository
}

func NewInstallmentService(transRepo *repositories.InstallmentRepository) *InstallmentService {
	return &InstallmentService{installmentRepository: transRepo}
}

func (s *InstallmentService) RegisterHTTPEndpoints(r chi.Router) {
    r.Get("/installments", s.ListInstallments)
    r.Post("/installments", s.CreateInstallment)
    r.Get("/installments/{id}", s.GetById)
    r.Put("/installments/{id}", s.UpdateRecord)
    r.Delete("/installments/{id}", s.DeleteRecord)
}

func (s *InstallmentService) ListInstallments(w http.ResponseWriter, r *http.Request) {
	installments, err := s.installmentRepository.ListInstallments(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(installments)
	return
}

func (s *InstallmentService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
	}
	installment, err := s.installmentRepository.GetByID(context.Background(), parsed_id)
	if err != nil {
		if installment.Name == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(installment)
	return
}

func (s *InstallmentService) CreateInstallment(w http.ResponseWriter, r *http.Request) {
	var installment models.Installment
	err := json.NewDecoder(r.Body).Decode(&installment)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}

	params := schemas.CreateInstallmentParams{
		Name: installment.Name,
		TotalCost: installment.TotalCost,
		InterestRate: installment.InterestRate,
		PeriodNum: installment.PeriodNum,
		PaidCost: installment.PaidCost,
		CurrentPeriod: installment.CurrentPeriod,
		PeriodCost: installment.PeriodCost,
        AccountID: installment.AccountID,
	}

	err = s.installmentRepository.CreateInstallment(context.Background(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	msg := "Installment created successfully"
	utils.WriteResponse(w, http.StatusOK, msg)
	return
}

func (s *InstallmentService) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateInstallmentParams
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
    if params.TotalCost != 0 {
        params.UpdateTotalCost = true
    }
    if params.InterestRate != 0 {
        params.UpdateInterestRate = true
    }
    if params.PeriodNum != 0 {
        params.UpdatePeriodNum = true
    }
    if params.PaidCost != 0 {
        params.UpdatePaidCost = true
    }
    if params.CurrentPeriod != 0 {
        params.UpdateCurrentPeriod = true
    }
    if params.PeriodCost != 0 {
        params.UpdatePeriodCost = true
    }
    if params.AccountID != 0 {
        params.UpdateAccount = true
    }

    err = s.installmentRepository.PartialUpdateInstallment(context.Background(), params)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Installment updated successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

func (s *InstallmentService) DeleteRecord(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    parsed_id, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
        return
    }

    err = s.installmentRepository.DeleteInstallment(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Installment deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

