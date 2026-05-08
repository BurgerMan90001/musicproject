

-- name: NewSong :one
INSERT INTO songs
(song_name, duration, album_id, creation_date,song_audio_url)
VALUES($1, $2, $3, $4, $5) RETURNING song_id;


-- name: PutSongCover :exec
UPDATE songs
SET song_cover_url=$1
WHERE song_id=$2;


-- INSERT INTO song_genres
-- (genre_id, song_id)
-- VALUES($1, $2);

-- INSERT INTO artist_songs
-- (artist_id, song_id)
-- VALUES($1, $2);

-- COMMIT;

-- type NewSongRequest struct {
-- 	// Required
-- 	Name     string   `json:"name"`
-- 	Artists  []string `json:"artists"`
-- 	Duration int      `json:"duration"`
-- 	// Optional
-- 	Genres []string `json:"genres"`
-- 	// YYYY-MM-DD format
-- 	CreationDate string `json:"creationDate"`

-- 	// Optional
-- 	AlbumID uuid.UUID `json:"albumId"`

-- 	// Optional
-- 	Image string `json:"image,omitempty"`
-- }

-- name: NewGenre :one 
INSERT INTO genres
(genre_name)
VALUES($1) RETURNING genre_id;

-- name: NewAlbum :one
INSERT INTO albums
(album_name, creation_date)
VALUES($1, $2) RETURNING album_id;

-- name: NewPlaylist :one
INSERT INTO playlists
(user_id, playlist_name)
VALUES($1, $2) RETURNING playlist_id;

-- name: NewArtist :one
INSERT INTO artists
(artist_name)
VALUES($1) RETURNING artist_id;


-- name: GetPlaylistSongsByID :many
SELECT
songs.song_id,
song_name,
artists,
genres,
streams,
duration,
songs.creation_date AS song_creation_date,  
albums.creation_date AS album_creation_date ,
albums.album_id,
albums.album_name,
song_cover_url, 	
song_audio_url,
playlists.playlist_name,
playlists.playlist_id,
playlists.user_id
FROM (SELECT string_agg(artists.artist_name, ',') AS artists, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genres, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id)
LEFT JOIN albums ON albums.album_id = a1.song_id
INNER JOIN playlist_songs ON playlist_songs.song_id = songs.song_id
LEFT JOIN playlists ON playlists.playlist_id = playlist_songs.playlist_id
WHERE playlist_songs.playlist_id = $1
LIMIT $2;

-- TODO name: GetSongsByGenre :many

-- name: GetSongs :many
SELECT
songs.song_id,
song_name,
artists,
genres,
streams,
duration,
songs.creation_date AS song_creation_date,  
albums.creation_date AS album_creation_date ,
albums.album_id,
albums.album_name,
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artists, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genres, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id)
LEFT JOIN albums ON albums.album_id = a1.song_id
LIMIT $1;

-- name: GetSongByID :one
SELECT
songs.song_id,
song_name,
artists,
genres,
streams,
duration,
songs.creation_date AS song_creation_date,  
albums.creation_date AS album_creation_date,
albums.album_id,
albums.album_name,
song_cover_url, 	
song_audio_url
FROM (SELECT string_agg(artists.artist_name, ',') AS artists, song_id 
    FROM artists
    INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
    WHERE artist_songs.song_id=$1
    GROUP BY song_id) a1
INNER JOIN
   (SELECT string_agg(genres.genre_name, ',') AS genres, song_id
    FROM genres
    INNER JOIN song_genres ON genres.genre_id = song_genres.genre_id
	WHERE song_genres.song_id=$1
  	GROUP BY song_id) a2
ON (a1.song_id = a2.song_id)
INNER JOIN songs ON (songs.song_id = a1.song_id)
LEFT JOIN albums ON albums.album_id = a1.song_id;

-- name: GetSongGenres :many
SELECT 
string_agg(genres.genre_name, ',') AS genres,
songs.song_id
FROM songs
INNER JOIN song_genres ON song_genres.song_id = songs.song_id
INNER JOIN genres ON song_genres.genre_id = genres.genre_id
WHERE songs.song_id=$1
GROUP BY songs.song_id LIMIT $2;

-- name: GetArtistSongs :many
SELECT 
    artists.artist_name,
    songs.song_name
FROM artists
INNER JOIN artist_songs ON artist_songs.artist_id = artists.artist_id
INNER JOIN songs ON songs.song_id = artist_songs.song_id
WHERE artists.artist_id=$1;

-- TODO name: GetAlubumSongs :many
SELECT *

FROM songs
INNER JOIN albums ON albums.album_id = songs.album_id
WHERE albums.album_id=$1;


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



-- TODO name: Search :many


-- name: DeleteSongByID :exec
DELETE FROM songs WHERE song_id=$1;

-- name: GetAggregatedRating :one
SELECT SUM(rating_value) / COUNT(rating_value)
FROM song_ratings 
WHERE song_id=$1;

