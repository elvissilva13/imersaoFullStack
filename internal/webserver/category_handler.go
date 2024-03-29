package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/elvissilva13/imersaoFullStack/goapi/internal/service"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}

}

func (wch *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := wch.CategoryService.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)

}

func (wch *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, id)
}
