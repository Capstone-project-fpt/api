DROP TRIGGER IF EXISTS topic_references_search_tsvector_trigger ON topic_references;

DROP FUNCTION IF EXISTS update_topic_references_search_tsvector();

ALTER TABLE topic_references DROP COLUMN IF EXISTS topic_references_tvs;

DROP INDEX IF EXISTS topic_references_tvs_idx;