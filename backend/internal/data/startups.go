package data

import "database/sql"

type Startup struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	AuthorID    int64  `json:"author_id"`
	Author      Author `json:"author"`
	Views       int64  `json:"views"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Pitch       string `json:"pitch"`
}

type StartupModel struct {
	DB *sql.DB
}
