package contracts

import "blog-backend/data/common"

type UpdateBlogRequest struct{
	ID 			uint64				`json:"blog_id"`
	BlogData	common.BlogData		`json:"blog_data"`
}

type UpdateBlogResponse struct{
	Status 				string			`json:"status"`
	UpdatedArticle		common.Blog		`json:"updated_article"`
}