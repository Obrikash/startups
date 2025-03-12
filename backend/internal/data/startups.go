package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Startup struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"-"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	AuthorID      int64     `json:"author_id"`
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
	query := fmt.Sprintf(`SELECT count(*) OVER(), id, created_at, title, slug, author_id, views, description, category, image_url, pitch_markdown
    FROM startup
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
			&startup.AuthorID,
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
