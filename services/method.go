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

type MethodService struct {
   methodRepository *repositories.MethodRepository
}

func NewMethodService(methodRepo *repositories.MethodRepository) *MethodService {
	return &MethodService{methodRepository: methodRepo}
}

func (s *MethodService) RegisterHTTPEndpoints(r chi.Router) {
    r.Get("/methods", s.ListMethods)
    r.Post("/methods", s.CreateMethod)
    r.Get("/methods/{id}", s.GetById)
    r.Put("/methods/{id}", s.UpdateRecord)
    r.Delete("/methods/{id}", s.DeleteMethod)
}

func (s *MethodService) ListMethods(w http.ResponseWriter, r *http.Request) {
	methods, err := s.methodRepository.ListMethods(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(methods)
	return
}

func (s *MethodService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpeced error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
	}
	method, err := s.methodRepository.GetByID(context.Background(), parsed_id)
	if err != nil {
		if method.Name == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(method)
	return
}

func (s *MethodService) CreateMethod(w http.ResponseWriter, r *http.Request) {
	var method models.Method
	err := json.NewDecoder(r.Body).Decode(&method)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}

	err = s.methodRepository.CreateMethod(context.Background(), method.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	msg := "Method created successfully"
	utils.WriteResponse(w, http.StatusOK, msg)
	return
}

func (s *MethodService) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateMethodParams
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

    err = s.methodRepository.PartialUpdateMethod(context.Background(), params)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Method updated successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

func (s *MethodService) DeleteMethod(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    parsed_id, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
        return
    }

    err = s.methodRepository.DeleteMethod(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Method deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}
