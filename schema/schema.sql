

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    email VARCHAR(255),
    password_hash VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS artists (
    id UUID PRIMARY KEY DEFAULT uuidv7()
);

CREATE TABLE IF NOT EXISTS songs (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
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
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    artist_id UUID,
    title VARCHAR(255),
    creation_date VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS ratings (
    song_id VARCHAR(255),
    user_id VARCHAR(255),
    value INT
);