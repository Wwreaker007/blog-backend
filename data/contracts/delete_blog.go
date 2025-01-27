package contracts

type DeleteBlogRequest struct{
	ID 		uint64			`json:"blog_id"`
}

type DeleteBlogResponse struct{
	Status 		string		`json:"status"`
}