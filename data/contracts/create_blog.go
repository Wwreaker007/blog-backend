package contracts

import "blog-backend/data/common"

type CreateBlogRequest struct{
	Title		string 			`json:"title"`
	Content		string			`json:"content"`
	Tags		[]common.TAG	`json:"tags"`
}

type CreateBlogResponse struct{
	Status		string 		`json:"status"`
}