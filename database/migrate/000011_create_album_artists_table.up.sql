CREATE TABLE IF NOT EXISTS album_artists (
    album_id UUID REFERENCES albums(album_id) NOT NULL,
    artist_id UUID REFERENCES artists(artist_id) NOT NULL
);