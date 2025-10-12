-- Enable citext extension if it doesn't already exist
CREATE EXTENSION IF NOT EXISTS citext;

-- Create users table if it doesn't already exist
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);