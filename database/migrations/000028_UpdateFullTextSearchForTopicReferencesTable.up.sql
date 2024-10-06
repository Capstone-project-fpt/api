CREATE EXTENSION IF NOT EXISTS unaccent;

ALTER TABLE "topic_references" ADD COLUMN IF NOT EXISTS "topic_references_tvs" tsvector;

CREATE OR REPLACE FUNCTION update_topic_references_search_tsvector() RETURNS trigger AS $$
BEGIN
  NEW.topic_references_tvs := to_tsvector('simple', unaccent(lower(NEW.name)));
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER topic_references_search_tsvector_trigger
BEFORE INSERT OR UPDATE ON topic_references
FOR EACH ROW EXECUTE FUNCTION update_topic_references_search_tsvector();

UPDATE topic_references
SET topic_references_tvs = to_tsvector('simple', unaccent(lower(name)));

CREATE INDEX topic_references_tvs_idx ON topic_references USING GIN(topic_references_tvs);