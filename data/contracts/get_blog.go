package contracts

import "blog-backend/data/common"

type GetBlogByIDRequest struct{
	ID			uint64			`json:"blog_id"`
}

type GetBlogByIDResponse struct{
	Status		string			`json:"status"`
	Article		common.Blog		`json:"article,omitempty"`
}

type GetBlogByTagsRequest struct{
	Tags		[]common.TAG	`json:"tags"`
}

type GetBlogByTagsResponse struct{
	Status		string			`json:"status"`
	Articles	[]common.Blog	`json:"articles,omitempty"`
}