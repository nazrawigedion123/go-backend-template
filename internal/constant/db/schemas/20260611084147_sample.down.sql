-- down migration: drop_sample_table.sql

DROP TRIGGER IF EXISTS trigger_sample_updated_at ON sample;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP INDEX IF EXISTS idx_sample_email;
DROP INDEX IF EXISTS idx_sample_is_active;
DROP TABLE IF EXISTS sample;