CREATE TABLE IF NOT EXISTS playlist_songs (
    playlist_id UUID REFERENCES playlists(playlist_id) NOT NULL,
    song_id UUID REFERENCES songs(song_id) NOT NULL
);