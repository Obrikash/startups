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
	Username string    `json:"-"`
	Email    string    `json:"-"`
	Image    string    `json:"-"`
	Bio      string    `json:"-"`
	Startups []Startup `json:"-"`
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
