CREATE TABLE IF NOT EXISTS artists (
    artist_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    artist_name VARCHAR(100) NOT NULL UNIQUE,
    artist_avatar_url TEXT
);