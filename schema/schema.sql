
/*  */
CREATE TABLE IF NOT EXISTS revoked_tokens (
    token VARCHAR(255) PRIMARY KEY,
    revocation_date timestamp default now()
);

CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuidv7(),
    email VARCHAR(255),
    password_hash VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS artists (
    artist_id UUID PRIMARY KEY DEFAULT uuidv7()
);

CREATE TABLE IF NOT EXISTS songs (
    song_id UUID PRIMARY KEY DEFAULT uuidv7(),
    album_id UUID,
    name VARCHAR(255), 
    genre VARCHAR(255),
    streams INT,
	duration INT,
	image VARCHAR(255),
    creation_date VARCHAR(255),
    src VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS albums (
    album_id UUID PRIMARY KEY DEFAULT uuidv7(),
    artist_id UUID,
    title VARCHAR(255),
    creation_date VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS genres {
    genre_id UUID PRIMARY KEY DEFAULT uuidv7(),
    genre VARCHAR(255)
}

CREATE TABLE IF NOT EXISTS ratings (
    song_id VARCHAR(255),
    user_id VARCHAR(255),
    value INT
);