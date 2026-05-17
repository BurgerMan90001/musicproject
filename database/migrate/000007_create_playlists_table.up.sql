CREATE TABLE IF NOT EXISTS playlists (
    playlist_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    user_id UUID,
    playlist_cover_url TEXT,
    playlist_name VARCHAR(100) NOT NULL
);
