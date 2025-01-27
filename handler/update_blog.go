package handler

import (
	"blog-backend/data/contracts"
	"blog-backend/types"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type UpdateBlogHandler struct{
	bsm		types.BlogServiceManager
}

func NewUpdateBlogHandler(service types.BlogServiceManager) *UpdateBlogHandler{
	return &UpdateBlogHandler{
		bsm: 	service,
	}
}

func (h *UpdateBlogHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.UpdateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := h.bsm.UpdateBlogByID(ctx, request)
	json.NewEncoder(w).Encode(response)
}