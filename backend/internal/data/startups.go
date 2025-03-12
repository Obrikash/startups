package data

import (
	"context"
	"database/sql"
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
