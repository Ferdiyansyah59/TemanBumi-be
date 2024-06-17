package entity

import "time"

type Article struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Description  string    `json:"description"`
	Excerpt   string    `json:"excerpt"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
}
