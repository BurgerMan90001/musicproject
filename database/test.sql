

INSERT INTO genres 
(genre_id, genre_name)
VALUES
('50104fc9-e0bc-40e0-ad48-f5347cd60896','Pop'),
('55b4817b-98aa-4615-bbd7-cf173d0845bd','Rap');

INSERT INTO albums
(album_id, album_name, creation_date)
VALUES('f7eed090-032b-4175-b014-7a1bd3667751','Fireman', '2048-20-32');

INSERT INTO songs
(song_id,song_name, duration, creation_date, song_audio_url, album_id)
VALUES
('f7d84ef8-75dc-4331-b1e3-5ada0770c4c0','Mickey', 123, '2006-03-23', '', 'f7eed090-032b-4175-b014-7a1bd3667751'),
('f594d596-8a24-4bfe-8660-71e0f79c0a91','Mickey 2', 400, '2006-03-24', '', NULL);

INSERT INTO artists
(artist_id, artist_name)
VALUES
('1928aca9-3f41-41a5-8103-8fb0c013aea0', 'Drake'),
('c05ca8f4-f969-4485-b38f-993513b4fbff', 'Mickey Mouse');

INSERT INTO artist_songs
(artist_id, song_id)
VALUES
('1928aca9-3f41-41a5-8103-8fb0c013aea0', 'f594d596-8a24-4bfe-8660-71e0f79c0a91'),
('1928aca9-3f41-41a5-8103-8fb0c013aea0', 'f7d84ef8-75dc-4331-b1e3-5ada0770c4c0'),
('c05ca8f4-f969-4485-b38f-993513b4fbff', 'f7d84ef8-75dc-4331-b1e3-5ada0770c4c0');

INSERT INTO song_genres
(genre_id, song_id)
VALUES
('50104fc9-e0bc-40e0-ad48-f5347cd60896', 'f7d84ef8-75dc-4331-b1e3-5ada0770c4c0'),
('55b4817b-98aa-4615-bbd7-cf173d0845bd', 'f7d84ef8-75dc-4331-b1e3-5ada0770c4c0'),
('50104fc9-e0bc-40e0-ad48-f5347cd60896', 'f594d596-8a24-4bfe-8660-71e0f79c0a91');

INSERT INTO album_artists
(album_id, artist_id)
VALUES('f7eed090-032b-4175-b014-7a1bd3667751', '1928aca9-3f41-41a5-8103-8fb0c013aea0');

INSERT INTO playlists
(playlist_id,user_id, playlist_name)
VALUES
('706a017c-a6b4-46c1-be1a-11c78dbbd4dc', 'a4218141-a605-4a1e-830b-8e95155a1eff','~my songs~');


INSERT INTO playlist_songs
(playlist_id, song_id)
VALUES
('706a017c-a6b4-46c1-be1a-11c78dbbd4dc','f7d84ef8-75dc-4331-b1e3-5ada0770c4c0'),
('706a017c-a6b4-46c1-be1a-11c78dbbd4dc', 'f594d596-8a24-4bfe-8660-71e0f79c0a91');
