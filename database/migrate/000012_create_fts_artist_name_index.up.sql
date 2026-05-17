CREATE INDEX CONCURRENTLY IF NOT EXISTS fts_artist_name 
ON artists 
USING GIN ((to_tsvector('english', artist_name)));