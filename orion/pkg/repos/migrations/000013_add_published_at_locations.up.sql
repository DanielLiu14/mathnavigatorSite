ALTER TABLE locations
    ADD COLUMN published_at datetime AFTER deleted_at;