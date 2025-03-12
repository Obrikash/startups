package data

import "database/sql"

type Models struct {
    Authors AuthorModel
    Startups StartupModel
}

func NewModels(db *sql.DB) Models {
    return Models{
        Authors: AuthorModel{},
        Startups: StartupModel{},
    }
}
