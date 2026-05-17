CREATE TABLE IF NOT EXISTS albums (
    album_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    album_name VARCHAR(100) NOT NULL,
    album_cover_url TEXT,
    creation_date VARCHAR(100) NOT NULL
);
