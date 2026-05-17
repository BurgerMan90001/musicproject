CREATE TABLE IF NOT EXISTS artist_songs (
    artist_id UUID REFERENCES artists(artist_id) NOT NULL,
    song_id UUID REFERENCES songs(song_id) NOT NULL
);
