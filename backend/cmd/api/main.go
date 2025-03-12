package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"startups/internal/data"
	"time"

    _ "github.com/lib/pq"
)

type config struct {
	port int
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *slog.Logger
    models data.Models
}

func main() {
	logger := newLogger()
	var cfg config
	parseFlags(&cfg)

    db, err := openDB(cfg)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    app := &application{
        config: cfg,
        logger: logger,
        models: data.NewModels(db),
    }

    err = app.serve()
    if err != nil {
        panic(err)
    }
}

func newLogger() *slog.Logger {

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
        ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
            if a.Key == slog.TimeKey {
                a.Key = "date"
                a.Value = slog.Int64Value(time.Now().Unix())
            }
            return a
        },
	})

	return slog.New(logHandler)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
