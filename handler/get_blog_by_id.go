package handler

import (
	"blog-backend/data/contracts"
	"blog-backend/types"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GetBlogByIDHandler struct{
	bsm		types.BlogServiceManager
}

func NewGetBlogByIDHandler(service types.BlogServiceManager) *GetBlogByIDHandler{
	return &GetBlogByIDHandler{
		bsm: 	service,
	}
}

func (h *GetBlogByIDHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.GetBlogByIDRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := h.bsm.GetBlogByID(ctx, request)
	json.NewEncoder(w).Encode(response)
}