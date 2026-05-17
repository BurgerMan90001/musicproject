

-- name: NewSong :one
INSERT INTO songs
(song_name, album_id, creation_date, song_audio_url)
VALUES($1, $2, $3, $4) RETURNING song_id;

-- name: NewGenre :one 
INSERT INTO genres
(genre_name)
VALUES($1) ON CONFLICT DO NOTHING
RETURNING genre_id;

-- name: NewAlbum :one
INSERT INTO albums
(album_name, creation_date, album_cover_url)
VALUES($1, $2, $3) RETURNING album_id;

-- name: NewPlaylist :one
INSERT INTO playlists
(user_id, playlist_name, playlist_cover_url)
VALUES($1, $2, $3) RETURNING playlist_id;

-- name: NewArtist :one
INSERT INTO artists
(artist_name, artist_avatar_url)
VALUES($1, $2) RETURNING artist_id;


-- name: GetPlaylistSongsById :many
SELECT
songs.song_id,
song_name,
artist_list,
genre_list,
streams,
album_id,
creation_date, 
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artist_list, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genre_list, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id)
INNER JOIN playlist_songs ON playlist_songs.song_id = songs.song_id
LEFT JOIN playlists ON playlists.playlist_id = playlist_songs.playlist_id
WHERE playlist_songs.playlist_id = $1
LIMIT $2;

-- name: GetAlbumById :one
SELECT * FROM albums
WHERE album_id=$1;

-- name: GetAlbumSongs :many
SELECT 
songs.song_id,
song_name,
artist_list,
genre_list,
album_id,
streams,
creation_date,
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artist_list, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genre_list, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id)
WHERE songs.album_id = $1
LIMIT $2;

-- name: GetPlaylists :many
SELECT
playlist_id,
user_id,
playlist_cover_url,
playlist_name
FROM playlists
LIMIT $1;

-- name: GetPlaylistById :many
SELECT
playlist_id,
user_id,
playlist_cover_url,
playlist_name
FROM playlists
WHERE playlist_id=$1;

-- name: GetSongsByGenre :many
SELECT
songs.song_id,
song_name,
artist_list,
genre_list,
streams,
creation_date,  
album_id,
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artist_list, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genre_list, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
    WHERE genre_name=$1
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id);


-- name: GetSongs :many
SELECT
songs.song_id,
song_name,
artist_list,
genre_list,
streams,
creation_date AS creation_date,
album_id,
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artist_list, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genre_list, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id)
LIMIT $1;

-- name: GetSongById :one
SELECT
songs.song_id,
song_name,
artist_list,
genre_list,
streams,
songs.creation_date AS creation_date,  
album_id,
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artist_list, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    WHERE artist_songs.song_id=$1
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genre_list, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
	WHERE song_genres.song_id=$1
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id);

-- name: GetSongGenres :many
SELECT 
genres.genre_id,
genres.genre_name
FROM songs
INNER JOIN song_genres ON song_genres.song_id = songs.song_id
INNER JOIN genres ON song_genres.genre_id = genres.genre_id
WHERE songs.song_id=$1
GROUP BY songs.song_id;

-- name: GetArtistSongs :many
SELECT
songs.song_id,
song_name,
artist_list,
genre_list,
streams,
creation_date,  
album_id,
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artist_list, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genre_list, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id)
WHERE artists.artist_id=$1
LIMIT $2;

-- name: GetGenres :many
SELECT genre_id, genre_name
FROM genres
LIMIT $1;

SELECT 
artist_id,
artist_name,
artist_avatar_url
FROM artists
LIMIT $1;
-- name: GetGenreByName :one
SELECT genre_id, genre_name
FROM genres
WHERE genre_name=$1;

-- name: GetGenreById :one
SELECT genre_id, genre_name
FROM genres
WHERE genre_id=$1;

-- name: PutSongAudio :exec
UPDATE songs SET song_audio_url = $1
WHERE song_id=$2;

-- name: PutSongCover :exec
UPDATE songs SET song_cover_url = $1
WHERE song_id=$2;

-- name: GetArtistById :one
SELECT 
artist_name,
artist_avatar_url
FROM artists
WHERE artist_id = $1;

-- name: GetArtistByName :one
SELECT 
artist_id,
artist_avatar_url
FROM artists
WHERE artist_name=$1;

-- name: PutSongGenres :exec
INSERT INTO song_genres
(genre_id, song_id)
VALUES(
    unnest(@genre_ids::UUID[]),
    $1
);

-- name: PutPlaylistSong :exec
INSERT INTO playlist_songs
(playlist_id, song_id)
VALUES($1, $2);

-- name: PutRating :exec
INSERT INTO song_ratings 
(song_id, user_id, rating_value) 
VALUES($1, $2, $3);

-- name: PutSongArtists :exec
INSERT INTO artist_songs
(song_id,artist_id)
VALUES(
    $1,
    unnest(@artist_ids::UUID[])  
);

-- name: PutAlbumArtist :exec
INSERT INTO album_artists
(album_id, artist_id)
VALUES($1, $2);

-- name: DeleteSongById :exec
DELETE FROM songs WHERE song_id=$1;

-- name: GetAggregatedRating :one
SELECT SUM(rating_value) / COUNT(rating_value)
FROM song_ratings 
WHERE song_id=$1;

