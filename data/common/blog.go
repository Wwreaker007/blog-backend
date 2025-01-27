package common

type Blog struct{
	BlogData
	ID			uint64		`json:"blog_id"`
	CreatedOn	string		`json:"created_on"`
	UpdatedOn	string		`json:"updated_on"`
}

type BlogData struct{
	Title		string 		`json:"title"`
	Content		string		`json:"content"`
	Tags		[]TAG		`json:"tags"`
}