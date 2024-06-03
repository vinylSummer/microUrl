-- +goose Up
CREATE TABLE IF NOT EXISTS url_bindings (
    id INTEGER PRIMARY KEY,
    short_url TEXT NOT NULL UNIQUE,
    long_url TEXT NOT NULL UNIQUE,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS url_bindings;