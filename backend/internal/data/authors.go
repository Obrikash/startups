package data

import "database/sql"

type Author struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Image    string    `json:"image"`
	Bio      string    `json:"bio"`
	Startups []Startup `json:"startups"`
}

type AuthorModel struct {
	DB *sql.DB
}
