package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Author struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"-"`
	ImageURL string    `json:"image_url"`
	Bio      string    `json:"bio"`
	Startups []Startup `json:"startups"`
}

type AuthorModel struct {
	DB *sql.DB
}

func (m AuthorModel) GetById(id int64) (*Author, error) {
	query := `SELECT id, name
        FROM author
        WHERE id = $1`

	var author Author
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&author.ID,
		&author.Name,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &author, nil
}
