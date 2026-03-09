CREATE TABLE IF NOT EXISTS public.urls (
    id         SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL UNIQUE,
    short_code   TEXT NOT NULL UNIQUE,
    created_at   TIMESTAMP DEFAULT NOW()
);
