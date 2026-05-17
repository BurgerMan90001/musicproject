CREATE TABLE IF NOT EXISTS songs (
    song_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    song_name VARCHAR(100) NOT NULL, 
    streams INT NOT NULL DEFAULT 0,
    creation_date VARCHAR(100) NOT NULL DEFAULT 'Unknown',
    
    album_id UUID REFERENCES albums(album_id),

    song_cover_url TEXT,
    song_audio_url TEXT NOT NULL
);