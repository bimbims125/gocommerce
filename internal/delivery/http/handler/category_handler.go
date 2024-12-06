package handler

import (
	"encoding/json"
	"gocommerce/internal/entity"
	"gocommerce/internal/usecase"
	"gocommerce/internal/utils"
	"net/http"
)

type CategoryHandler struct {
	usecase *usecase.CategoryUseCase
}

func NewCategoryHandler(usecase *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{usecase: usecase}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.usecase.CreateCategory(r.Context(), &category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.usecase.GetCategories(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, categories)
}
