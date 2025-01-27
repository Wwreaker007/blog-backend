package handler

import (
	"blog-backend/data/contracts"
	"blog-backend/types"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GetBlogByTagsHandler struct{
	bsm		types.BlogServiceManager
}

func NewGetBlogByTagsHandler(service types.BlogServiceManager) *GetBlogByTagsHandler{
	return &GetBlogByTagsHandler{
		bsm: 	service,
	}
}

func (h *GetBlogByTagsHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.GetBlogByTagsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := h.bsm.GetBlogByTags(ctx, request)
	json.NewEncoder(w).Encode(response)
}