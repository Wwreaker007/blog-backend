package bsm

import (
	"blog-backend/data/common"
	"blog-backend/data/contracts"
	"blog-backend/types"
	"context"
	"log"
)

type BlogService struct{
	db		types.BlogRepositoryManager
}

func NewBlogService(client types.BlogRepositoryManager) *BlogService{
	return &BlogService{
		db:		client,
	}
}

func (b *BlogService) CreateBlog(ctx context.Context, r contracts.CreateBlogRequest) contracts.CreateBlogResponse{
	blogData := common.BlogData{
		Title:		r.Title,
		Content:	r.Content,
		Tags:		r.Tags,
	}
	err := b.db.CreateBlog(ctx, blogData)
	if err != nil{
		log.Println(err)
		return contracts.CreateBlogResponse{
			Status: "FAILED: " + err.Error(),
		} 
	}
	return contracts.CreateBlogResponse{
		Status: "SUCCESS",
	}
}

func (b *BlogService) GetBlogByID(ctx context.Context, r contracts.GetBlogByIDRequest) contracts.GetBlogByIDResponse{
	blog, err := b.db.GetBlogByID(ctx, r.ID)
	if err != nil{
		log.Println(err)
		return contracts.GetBlogByIDResponse{
			Status: "FAILED: " + err.Error(),
		}
	}

	return contracts.GetBlogByIDResponse{
		Status: "SUCCESS",
		Article: blog,
	}
}

func (b *BlogService) GetBlogByTags(ctx context.Context, r contracts.GetBlogByTagsRequest) contracts.GetBlogByTagsResponse{
	blogs, err := b.db.GetBlogByTags(ctx, r.Tags)
	if err != nil{
		log.Println(err)
		return contracts.GetBlogByTagsResponse{
			Status: "FAILED: " + err.Error(),
		}
	}
	return	contracts.GetBlogByTagsResponse{
		Status: "SUCCESS",
		Articles: blogs,
	}
}

func (b *BlogService) DeleteBlogByID(ctx context.Context, r contracts.DeleteBlogRequest) contracts.DeleteBlogResponse{
	err := b.db.DeleteBlog(ctx, r.ID)
	if err != nil{
		log.Println(err)
		return contracts.DeleteBlogResponse{
			Status: "FAILED: " + err.Error(),
		}
	}
	return contracts.DeleteBlogResponse{
		Status: "SUCCESS",
	}

}

func (b *BlogService) UpdateBlogByID(ctx context.Context, r contracts.UpdateBlogRequest) contracts.UpdateBlogResponse{
	blog, err := b.db.GetBlogByID(ctx, r.ID)
	if err != nil{
		log.Println(err)
		return contracts.UpdateBlogResponse{
			Status: "FAILED: " + err.Error(),
		}
	}
	blog.Title = r.BlogData.Title
	blog.Content = r.BlogData.Content
	blog.Tags = r.BlogData.Tags

	updatedBlog, err := b.db.UpdateBlog(ctx, blog)
	if err != nil{
		log.Println(err)
		return contracts.UpdateBlogResponse{
			Status: "FAILED: " + err.Error(),
		}
	}

	return contracts.UpdateBlogResponse{
		Status: "SUCCESS",
		UpdatedArticle: updatedBlog,
	}
}