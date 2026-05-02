
-- name: GetUserByID :one
SELECT email, password_hash 
FROM users 
WHERE user_id=$1;

-- name: GetUserByEmail :one
SELECT user_id, password_hash 
FROM users 
WHERE email=$1;

-- name: PutUser :one
INSERT INTO users (email, password_hash) 
VALUES($1, $2) 
RETURNING user_id;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE user_id=$1;

-- name: GetAggregatedRating :one
SELECT SUM(rating_value) / COUNT(rating_value)
FROM song_ratings 
WHERE song_id=$1;

-- name: PutRating :exec
INSERT INTO song_ratings 
(song_id, user_id, rating_value) 
VALUES($1, $2, $3);

-- name: GetSongByID :one
SELECT * FROM songs
WHERE song_id=$1;

-- name: GetSongRandom :many
SELECT * FROM songs 
ORDER BY RANDOM() LIMIT $1;

-- name: PutSong :one
INSERT INTO songs
(album_id, song_name, streams, duration, creation_date)
VALUES($1, $2, $3, $4, $5) RETURNING song_id;

-- name: DeleteSongByID :exec
DELETE FROM songs WHERE song_id=$1;

-- name: GetPlaylistByID :one
SELECT * FROM playlists
WHERE playlist_id=$1;

-- name: PutPlaylist :one
INSERT INTO playlists
(user_id)
VALUES($1) RETURNING playlist_id;

-- asdasd name: GetPlaylistSongs :many
-- SELECT * FROM playlists
-- WHERE 

-- name: PutPlaylistSong :exec
INSERT INTO playlist_songs
(playlist_id, song_id)
VALUES($1,$2);