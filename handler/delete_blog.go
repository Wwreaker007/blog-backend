package handler

import (
	"blog-backend/data/contracts"
	"blog-backend/types"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type DeleteBlogHandler struct{
	bsm		types.BlogServiceManager
}

func NewDeleteBlogHandler(service types.BlogServiceManager) *DeleteBlogHandler{
	return &DeleteBlogHandler{
		bsm: 	service,
	}
}

func (h *DeleteBlogHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.DeleteBlogRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := h.bsm.DeleteBlogByID(ctx, request)
	json.NewEncoder(w).Encode(response)
}