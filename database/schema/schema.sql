CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS genres (
    genre_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    genre_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS artists (
    artist_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    artist_name VARCHAR(100) NOT NULL UNIQUE,
    artist_avatar_url TEXT
);

CREATE TABLE IF NOT EXISTS albums (
    album_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    album_name VARCHAR(100) NOT NULL,
    album_cover_url TEXT,
    creation_date VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS songs (
    song_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    song_name VARCHAR(100) NOT NULL, 
    streams INT NOT NULL DEFAULT 0,
    duration INT NOT NULL, 
    creation_date VARCHAR(100) NOT NULL DEFAULT 'Unknown',
    
    album_id UUID REFERENCES albums(album_id),

    song_cover_url TEXT,
    song_audio_url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS song_ratings (
    user_id UUID,
    song_id UUID REFERENCES albums(album_id) NOT NULL,
    rating_value INT NOT NULL
);

CREATE TABLE IF NOT EXISTS playlists (
    playlist_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    user_id UUID,
    playlist_cover_url TEXT,
    playlist_name VARCHAR(100) NOT NULL
);

-- link table playlists songs
CREATE TABLE IF NOT EXISTS playlist_songs (
    playlist_id UUID REFERENCES playlists(playlist_id) NOT NULL,
    song_id UUID REFERENCES songs(song_id) NOT NULL
);


-- link table artists, songs
CREATE TABLE IF NOT EXISTS artist_songs (
    artist_id UUID REFERENCES artists(artist_id) NOT NULL,
    song_id UUID REFERENCES songs(song_id) NOT NULL
);

-- link table songs, genres
CREATE TABLE IF NOT EXISTS song_genres (
    genre_id UUID REFERENCES genres(genre_id) NOT NULL,
    song_id UUID REFERENCES songs(song_id) NOT NULL
);

-- link table albums artists
CREATE TABLE IF NOT EXISTS album_artists (
    album_id UUID REFERENCES albums(album_id) NOT NULL,
    artist_id UUID REFERENCES artists(artist_id) NOT NULL
);

CREATE INDEX IF NOT EXISTS fts_artist_name 
ON artists 
USING GIN ((to_tsvector('english', artist_name)));



