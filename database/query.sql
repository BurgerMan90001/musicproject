

-- name: NewSong :one
INSERT INTO songs
(album_id, song_name, duration, creation_date, song_url, song_image)
VALUES($1, $2, $3, $4, $5, $6) RETURNING song_id;

-- name: NewGenre :one 
INSERT INTO genres
(genre_name)
VALUES($1) RETURNING genre_id;

-- name: NewAlbum :one
INSERT INTO albums
(album_name, creation_date, artist_id)
VALUES($1, $2, $3) RETURNING album_id;

-- name: NewPlaylist :one
INSERT INTO playlists
(user_id, playlist_name)
VALUES($1, $2, $3) RETURNING playlist_id;

-- name: NewArtist :one
INSERT INTO artists
(artist_name)
VALUES($1) RETURNING artist_id;


-- name: GetSongByID :one
SELECT 
    songs.song_name, 
    songs.streams,
	songs.duration,
    songs.creation_date,
    songs.song_image,
    songs.song_url,
    albums.album_name,
    artists.artist_name
FROM songs
LEFT JOIN albums ON albums.album_id = songs.album_id
LEFT JOIN artist_songs ON artist_songs.song_id = songs.song_id
LEFT JOIN artists ON artists.artist_id = artist_songs.artist_id
WHERE songs.song_id=$1;


-- name: GetPlaylistSongsByID :many
SELECT 
    playlists.user_id,
    playlists.playlist_name,
    playlist_songs.playlist_id,

    songs.song_name, 
    songs.streams,
	songs.duration, 
    songs.creation_date,
    songs.song_image,
    songs.song_url,

    albums.album_name,
    artists.artist_name

FROM songs
INNER JOIN playlist_songs ON playlist_songs.song_id = songs.song_id
INNER JOIN playlists ON playlists.playlist_id = playlist_songs.playlist_id

LEFT JOIN albums ON albums.album_id = songs.album_id 

LEFT JOIN artist_songs ON artist_songs.song_id = songs.song_id
LEFT JOIN artists ON artists.artist_id = artist_songs.artist_id
WHERE playlists.playlist_id=$1 LIMIT $2;

-- name: GetSongsByGenre :many
SELECT
genre_name,
songs.song_id, 
song_name, 
streams,
duration,
songs.creation_date,
albums.album_name

FROM genres
INNER JOIN song_genres ON genres.genre_id=song_genres.genre_id 
INNER JOIN songs ON song_genres.song_id=songs.song_id
LEFT JOIN albums ON songs.album_id=albums.album_id 
WHERE genres.genre_name=$1 LIMIT $2;


-- name: GetSongGenres :many
SELECT genre_name
FROM genres
WHERE song_id=$1 LIMIT $2;

-- name: GetArtistSongs :many
SELECT 
    artists.artist_name,
    songs.song_name
FROM artists
INNER JOIN artist_songs ON artist_songs.song_id = songs.song_id
INNER JOIN songs ON songs.song_id = artist_songs.song_id
WHERE artists_songs=$1;


-- name: GetRandomSongs :many
SELECT * FROM songs 
ORDER BY RANDOM() LIMIT $1;


-- name: PutSongGenre :exec
INSERT INTO song_genres
(genre_id, song_id)
VALUES($1, $2);

-- name: PutPlaylistSong :exec
INSERT INTO playlist_songs
(playlist_id, song_id)
VALUES($1, $2);

-- name: PutRating :exec
INSERT INTO song_ratings 
(song_id, user_id, rating_value) 
VALUES($1, $2, $3);

-- name: PutArtistSong :exec
INSERT INTO artist_songs
(artist_id, song_id)
VALUES($1, $2);


-- name: DeleteSongByID :exec
DELETE FROM songs WHERE song_id=$1;

-- name: GetAggregatedRating :one
SELECT SUM(rating_value) / COUNT(rating_value)
FROM song_ratings 
WHERE song_id=$1;

