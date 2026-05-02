


CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuidv7(),
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(100),
    created_at DATE DEFAULT NOW() NOT NULL,
    avatar_url VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID REFERENCES users(user_id),
    role_name VARCHAR(100)
);


CREATE TABLE IF NOT EXISTS artists (
    artist_id UUID PRIMARY KEY DEFAULT uuidv7(),
    artist_name VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS genres (
    genre_id UUID PRIMARY KEY DEFAULT uuidv7(),
    genre VARCHAR(100)
);


CREATE TABLE IF NOT EXISTS albums (
    album_id UUID PRIMARY KEY DEFAULT uuidv7(),
    artist_id UUID REFERENCES artists(artist_id),
    album_name VARCHAR(100),
    genre_id UUID REFERENCES genres(genre_id),
    creation_date VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS songs (
    song_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    album_id UUID REFERENCES albums(album_id),

    song_name VARCHAR(100) NOT NULL, 
    -- genre VARCHAR(100) REFERENCES genres(genre_id),
    streams INT NOT NULL DEFAULT 0,
	duration INT NOT NULL, 
    creation_date VARCHAR(100) NOT NULL DEFAULT 'Unknown',

    song_image VARCHAR(100),
    song_url VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS song_ratings (
    user_id UUID REFERENCES users(user_id) NOT NULL,
    song_id UUID REFERENCES albums(album_id) NOT NULL,
    rating_value INT NOT NULL
);

CREATE TABLE IF NOT EXISTS playlists (
    playlist_id UUID PRIMARY KEY DEFAULT uuidv7() NOT NULL,
    user_id UUID REFERENCES users(user_id) NOT NULL,
    playlist_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS playlist_songs (
    playlist_id UUID REFERENCES playlists(playlist_id) NOT NULL,
    song_id UUID REFERENCES songs(song_id) NOT NULL
);


