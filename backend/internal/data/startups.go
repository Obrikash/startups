package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Startup struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	AuthorID      int64     `json:"-"`
	Author        Author    `json:"author"`
	Views         int64     `json:"views"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	ImageURL      string    `json:"image_url"`
	PitchMarkdown string    `json:"pitch"`
}

type StartupModel struct {
	DB *sql.DB
}

func (s StartupModel) GetAll(title string, filters Filters) ([]*Startup, Metadata, error) {
	query := fmt.Sprintf(`SELECT count(*) OVER(), 
    startup.id, 
    startup.created_at, 
    startup.title, 
    startup.slug, 
    startup.author_id, 
    author.name AS author_name, 
    startup.views, 
    startup.description, 
    startup.category, 
    startup.image_url, 
    startup.pitch_markdown
    FROM startup
    INNER JOIN author ON startup.author_id = author.id
    WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
    ORDER BY %s %s, id ASC
    LIMIT $2 OFFSET $3`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.DB.QueryContext(ctx, query, title, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	startups := []*Startup{}

	for rows.Next() {
		var startup Startup

		err := rows.Scan(
			&totalRecords,
			&startup.ID,
			&startup.CreatedAt,
			&startup.Title,
			&startup.Slug,
			&startup.Author.ID,
            &startup.Author.Name,
			&startup.Views,
			&startup.Description,
			&startup.Category,
			&startup.ImageURL,
			&startup.PitchMarkdown,
		)
        if err != nil {
            return nil, Metadata{}, err
        }
        startups = append(startups, &startup)
	}

    if err = rows.Err(); err != nil {
        return nil, Metadata{}, err
    }

    metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

    return startups, metadata, nil
}

func (s StartupModel) Get(id int64) (*Startup, error) {
    if id < 1 {
        return nil, ErrRecordNotFound
    }

    query := `SELECT 
    startup.id, 
    startup.title, 
    startup.slug, 
    startup.created_at, 
    author.id AS author_id,
    author.name AS author_name,
    author.image_url AS author_image_url,
    author.bio AS author_bio,
    author.username AS author_username,
    startup.views,
    startup.description,
    startup.category,
    startup.image_url,
    startup.pitch_markdown
    FROM startup INNER JOIN author ON startup.author_id = author.id
    WHERE startup.id = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    startup := &Startup{}

    err := s.DB.QueryRowContext(ctx, query, id).Scan(
        &startup.ID,
        &startup.Title,
        &startup.Slug,
        &startup.CreatedAt,
        &startup.Author.ID,
        &startup.Author.Name,
        &startup.Author.ImageURL,
        &startup.Author.Bio,
        &startup.Author.Username,
        &startup.Views,
        &startup.Description,
        &startup.Category,
        &startup.ImageURL,
        &startup.PitchMarkdown,
    )

    if err != nil {
        switch {
            case errors.Is(err, sql.ErrNoRows):
                return nil, ErrRecordNotFound
            default:
                return nil, err
        }
    }

    return startup, nil
}

func (s StartupModel) UpdateViews(id int64) error {
    query := `UPDATE startup SET views = views + 1
    WHERE id = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    result, err := s.DB.ExecContext(ctx, query, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return ErrRecordNotFound
    }

    return nil
}
