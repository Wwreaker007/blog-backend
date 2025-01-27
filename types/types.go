package types

import (
	"blog-backend/data/common"
	"blog-backend/data/contracts"
	"context"
)

type BlogRepositoryManager interface{
	CreateBlog(context.Context, common.BlogData) (error)
	DeleteBlog(context.Context, uint64) (error)
	GetBlogByID(context.Context, uint64) (common.Blog, error)
	GetBlogByTags(context.Context, []common.TAG) ([]common.Blog, error)
	UpdateBlog(context.Context, common.Blog) (common.Blog, error)
}

type BlogServiceManager interface{
	CreateBlog(context.Context, contracts.CreateBlogRequest) contracts.CreateBlogResponse
	GetBlogByID(context.Context, contracts.GetBlogByIDRequest) contracts.GetBlogByIDResponse
	GetBlogByTags(context.Context, contracts.GetBlogByTagsRequest) contracts.GetBlogByTagsResponse
	DeleteBlogByID(context.Context, contracts.DeleteBlogRequest) contracts.DeleteBlogResponse
	UpdateBlogByID(context.Context, contracts.UpdateBlogRequest) contracts.UpdateBlogResponse
}