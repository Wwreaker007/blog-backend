package handler

import (
	"blog-backend/data/contracts"
	"blog-backend/types"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type CreateBlogHandler struct{
	bsm		types.BlogServiceManager
}

func NewCreateBlogHandler(service types.BlogServiceManager) *CreateBlogHandler{
	return &CreateBlogHandler{
		bsm: 	service,
	}
}

func (h *CreateBlogHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.CreateBlogRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := h.bsm.CreateBlog(ctx, request)
	json.NewEncoder(w).Encode(response)
}