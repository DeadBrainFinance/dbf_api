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

type CategoryService struct {
   categoryRepository *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepo}
}

func (s *CategoryService) RegisterHTTPEndpoints(r chi.Router) {
    r.Get("/categories", s.ListCategories)
    r.Post("/categories", s.CreateCategory)
    r.Get("/categories/{id}", s.GetById)
    r.Put("/categories/{id}", s.UpdateRecord)
    r.Delete("/categories/{id}", s.DeleteCategory)
}

func (s *CategoryService) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := s.categoryRepository.ListCategories(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(categories)
	return
}

func (s *CategoryService) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsed_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusInternalServerError)
		log.Println("Can't parse id in the URL.")
		return
	}
	category, err := s.categoryRepository.GetByID(context.Background(), parsed_id)
	if err != nil {
		if category.Name == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	utils.WriteResponse(w, http.StatusOK, "")
	json.NewEncoder(w).Encode(category)
	return
}

func (s *CategoryService) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
		return
	}


	params := schemas.CreateCategoryParams{
		Name: category.Name,
		Description: category.Description,
	}

	err = s.categoryRepository.CreateCategory(context.Background(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	msg := "Category created successfully"
	utils.WriteResponse(w, http.StatusOK, msg)
	return
}

func (s *CategoryService) UpdateRecord(w http.ResponseWriter, r *http.Request) {
    var params schemas.PartialUpdateCategoryParams
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

    if params.Name != "" {
        params.UpdateName = true
    }
    if params.Description != "" {
        params.UpdateDescription = true
    }

    err = s.categoryRepository.PartialUpdateCategory(context.Background(), params)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Category updated successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

func (s *CategoryService) DeleteCategory(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    parsed_id, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        http.Error(w, "Unexpected error", http.StatusBadRequest)
        log.Println(err.Error())
        return
    }

    err = s.categoryRepository.DeleteCategory(context.Background(), parsed_id)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        log.Println(err.Error())
        return
    }

    msg := "Category deleted successfully"
    utils.WriteResponse(w, http.StatusOK, msg)
    return
}

