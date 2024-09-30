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

type BalanceSheetService struct {
   balancesheetRepository *repositories.BalanceSheetRepository
}

func NewBalanceSheetService(transRepo *repositories.BalanceSheetRepository) *BalanceSheetService {
	return &BalanceSheetService{balancesheetRepository: transRepo}
}

func (s *BalanceSheetService) RegisterHTTPEndpoints(r chi.Router) {
    r.Get("/balancesheets", s.ListBalanceSheets)
    r.Post("/balancesheets", s.CreateBalanceSheet)
    r.Get("/balancesheets/{id}", s.GetById)
    r.Put("/balancesheets/{id}", s.UpdateRecord)
    r.Delete("/balancesheets/{id}", s.DeleteRecord)
}

func (s *BalanceSheetService) ListBalanceSheets(w http.ResponseWriter, r *http.Request) {
	balancesheets, err := s.balancesheetRepository.ListBalanceSheets(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(balancesheets)
	return
}

func (s *BalanceSheetService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
	}
	balanceSheet, err := s.balancesheetRepository.GetByID(context.Background(), parsed_id)
	if err != nil {
		if balanceSheet.Month == 0 {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(balanceSheet)
	return
}

func (s *BalanceSheetService) CreateBalanceSheet(w http.ResponseWriter, r *http.Request) {
	var balancesheet models.BalanceSheet
	err := json.NewDecoder(r.Body).Decode(&balancesheet)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}

	params := schemas.CreateBalanceSheetParams{
		Month: balancesheet.Month,
		Year: balancesheet.Year,
		Allocation: balancesheet.Allocation,
		Paid: balancesheet.Paid,
		Remaining: balancesheet.Remaining,
        CategoryID: balancesheet.CategoryID,
	}

	err = s.balancesheetRepository.CreateBalanceSheet(context.Background(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	msg := "BalanceSheet created successfully"
	utils.WriteResponse(w, http.StatusOK, msg)
	return
}

func (s *BalanceSheetService) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateBalanceSheetParams
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
    }

    err = json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    params.ID = parsed_id

    if params.Month != 0 {
        params.UpdateMonth = true
    }
    if params.Year != 0 {
        params.UpdateYear = true
    }
    if params.Allocation != 0 {
        params.UpdateAllocation = true
    }
    if params.Paid != 0 {
        params.UpdatePaid = true
    }
    if params.Remaining != 0 {
        params.UpdateRemaining = true
    }
    if params.CategoryID != 0 {
        params.UpdateCategories = true
    }

    err = s.balancesheetRepository.PartialUpdateBalanceSheet(context.Background(), params)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "BalanceSheet updated successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

func (s *BalanceSheetService) DeleteRecord(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    parsed_id, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
        return
    }

    err = s.balancesheetRepository.DeleteBalanceSheet(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "BalanceSheet deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}
