CREATE TABLE IF NOT EXISTS quotes (
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    body TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);