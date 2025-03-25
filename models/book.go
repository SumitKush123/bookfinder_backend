package models

import "time"

type Books struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Title       string    `gorm:"column:title;not null"`
	Thumbnail   string    `gorm:"column:thumbnail"`
	WritterName string    `gorm:"column:writter_name"`
	Description string    `gorm:"column:description"`
	Category    string    `gorm:"column:category"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Books) TableName() string {
	return "books"
}

type BookReqs struct {
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	WritterName string    `json:"writter_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BooksResp struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	WritterName string    `json:"writter_name"`
	Description string    `json:"description"`
	Cart        bool      `json:"cart"`
	Relevance   float64   `json:"relevance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookWithCart struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	WritterName string    `json:"writter_name"`
	Cart        bool      `json:"cart"`
	Limit       int64     `json:"limit"`
	Page        int64     `json:"page"`
	TotalPage   int64     `json:"total_page"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SearchByInputParam struct {
	WritterName string `query:"writter_name"`
	Title       string `query:"title"`
	UserID      int64  `json:"user_id"`
	Limit       int64  `json:"limit"`
	Page        int64  `json:"page"`
}

type BulkInsertBook struct {
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	PublishedDate time.Time `json:"published_date"`
	Thumbnail     string    `json:"thumbnail"`
	AuthorDetails string    `gorm:"author_details"`
	Summary       string    `gorm:"summary"`
}

type BulkInsertBookReqs struct {
	File *multipart.FileHeader
}

type RecommendReqs struct {
	BookID int64 `json:"book_id"`
}
type BookRespData struct {
	Data []BooksResp
	Meta MetaPagination
}
