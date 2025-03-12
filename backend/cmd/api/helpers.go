package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/url"
	"startups/internal/validator"
	"strconv"
)

type envelope map[string]any

func parseFlags(cfg *config) {
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

    flag.Parse()
}

func (app *application) readInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
    s := qs.Get(key)

    if s == "" {
        return defaultValue
    }

    number, err := strconv.Atoi(s)
    if err != nil {
        v.AddError(key, "must be an integer value")
    }

    return number
}

func (app *application) readString(qs url.Values, key string, defaultValue string) string {
    s := qs.Get(key)

    if s == "" {
        return defaultValue
    }

    return s
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
    js, err := json.Marshal(data)
    if err != nil {
        return err
    }

    js = append(js, '\n')

    for key, value := range headers {
        w.Header()[key] = value
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)

    return nil
}
