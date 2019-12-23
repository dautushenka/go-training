package handlers

import (
	"go-training/proj/core/repository"
	"go-training/proj/server"
	"net/http"
)

type CategoryHandler struct {
	repository repository.CategoryRepository
}

func CreateCategoryHandler(repository repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repository}
}

func (h *CategoryHandler) GetList(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repository.GetCategoryList()
	if err != nil {
		panic(err)
	}

	server.WriteJSONResponse(w, categories, http.StatusOK)
}
