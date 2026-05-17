CREATE TABLE IF NOT EXISTS song_genres (
    genre_id UUID REFERENCES genres(genre_id) NOT NULL,
    song_id UUID REFERENCES songs(song_id) NOT NULL
);