package db

import (
	"blog-backend/data/common"
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const (
	CREATE_BLOG_QUERY = "INSERT INTO blogs (title, content, tags) VALUES ($1, $2, $3)"
	DELETE_BLOG_QUERY = "DELETE FROM blogs WHERE id = $1"
	GET_BLOG_BY_ID_QUERY = "SELECT * FROM blogs WHERE id = $1"
	GET_BLOG_BY_TAGS = "SELECT * FROM blogs WHERE tags && $1::TEXT[]"
	UPDATE_BLOG_QUERY = "UPDATE blogs SET title = $1, content = $2, tags = $3 WHERE id = $4"
)

type BlogRepository struct{
	db		*sql.DB
}

func NewBlogRepository(client *sql.DB) *BlogRepository{
	return &BlogRepository{
		db: client,
	}
}

func (b *BlogRepository) CreateBlog(ctx context.Context, blog common.BlogData) error {
	_, err := b.db.Exec(CREATE_BLOG_QUERY, blog.Title, blog.Content, pq.Array(convertTagsToStrings(blog.Tags)))
	if err != nil {
		return err
	}
	return nil
}

func (b *BlogRepository) DeleteBlog(ctx context.Context, id uint64) error {
	_, err := b.db.Exec(DELETE_BLOG_QUERY, id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlogRepository) GetBlogByID(ctx context.Context, id uint64) (common.Blog, error) {
	var blog common.Blog
	var tags []string
	err := b.db.QueryRow(GET_BLOG_BY_ID_QUERY, id).Scan(&blog.ID, &blog.Title, &blog.Content, pq.Array(&tags), &blog.CreatedOn, &blog.UpdatedOn)
	if err != nil {
		return common.Blog{}, err
	}
	blog.Tags = convertStringsToTags(tags)
	return blog, nil
}

func (b *BlogRepository) GetBlogByTags(ctx context.Context, tags []common.TAG) ([]common.Blog, error) {
	rows, err := b.db.Query(GET_BLOG_BY_TAGS, pq.Array(tags))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []common.Blog
	for rows.Next() {
		var blog common.Blog
		var tempTags []string
		err = rows.Scan(&blog.ID, &blog.Title, &blog.Content, pq.Array(&tempTags), &blog.CreatedOn, &blog.UpdatedOn)
		if err != nil {
			return nil, err
		}
		blog.Tags = convertStringsToTags(tempTags)
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (b *BlogRepository) UpdateBlog(ctx context.Context, updatedBlog common.Blog) (common.Blog, error) {
	_, err := b.db.Exec(UPDATE_BLOG_QUERY, updatedBlog.Title, updatedBlog.Content, pq.Array(convertTagsToStrings(updatedBlog.Tags)), updatedBlog.ID)
	if err != nil {
		return common.Blog{}, err
	}
	return updatedBlog, nil
}

// Convert []TAG to []string
func convertTagsToStrings(tags []common.TAG) []string {
	stringTags := make([]string, len(tags))
	for i, tag := range tags {
		stringTags[i] = string(tag)
	}
	return stringTags
}

// Convert []string to []TAG
func convertStringsToTags(tags []string) []common.TAG {
	tagEnums := make([]common.TAG, len(tags))
	for i, tag := range tags {
		tagEnums[i] = common.TAG(tag)
	}
	return tagEnums
}