CREATE TABLE IF NOT EXISTS song_ratings (
    user_id UUID,
    song_id UUID REFERENCES albums(album_id) NOT NULL,
    rating_value INT NOT NULL
);
