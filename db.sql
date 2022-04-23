-- CREATE DATABASE notice OWNER postgres;
CREATE TABLE notes(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    title TEXT,
    info TEXT
)