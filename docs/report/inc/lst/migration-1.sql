--
-- Initialize database with basic tables
--
CREATE TABLE IF NOT EXISTS migrations (
  id VARCHAR PRIMARY KEY
);

SELECT EXISTS (
  SELECT id FROM migrations WHERE id = :'MIGRATION_ID'
) as migrated \gset

\if :migrated
  \echo 'migration' :MIGRATION_ID 'already exists, skipping'
\else
  \echo 'migration' :MIGRATION_ID 'does not exist'

  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

  CREATE TABLE users_info (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4()
    , username      VARCHAR(64)  UNIQUE NOT NULL
    , email         VARCHAR(64)  UNIQUE
    , fullname      VARCHAR(256) NOT NULL
    , password_hash VARCHAR(256) NOT NULL
    , flags         NUMERIC      CHECK (flags > 0)
    , blocked       TIMESTAMP
    , created_at    TIMESTAMP DEFAULT NOW()
    , updated_at    TIMESTAMP DEFAULT NOW()
  );

  INSERT INTO migrations(id) VALUES (:'MIGRATION_ID');
\endif

COMMIT;
