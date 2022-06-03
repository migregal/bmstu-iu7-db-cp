--
-- Create database roles for statistics, backend admins and users
--
SELECT EXISTS (
  SELECT id FROM migrations WHERE id = :'MIGRATION_ID'
) as migrated \gset

\if :migrated
  \echo 'migration' :MIGRATION_ID 'already exists, skipping'
\else
  \echo 'migration' :MIGRATION_ID 'does not exist'

  \set STAT_USERNAME    `echo $STAT_USERNAME`
  \set STAT_PASSWORD    `echo $STAT_PASSWORD`
  \set ADMIN_USERNAME   `echo $ADMIN_USERNAME`
  \set ADMIN_PASSWORD   `echo $ADMIN_PASSWORD`
  \set REGULAR_USERNAME `echo $REGULAR_USERNAME`
  \set REGULAR_PASSWORD `echo $REGULAR_PASSWORD`

  \set exit_error false

  SELECT (:'STAT_USERNAME' = '') as is_empty \gset
  \if :is_empty
    \warn 'STAT_USERNAME is empty'
    \set exit_error true
  \endif

  SELECT (:'STAT_PASSWORD' = '') as is_empty \gset
  \if :is_empty
    \warn 'STAT_PASSWORD is empty'
    \set exit_error true
  \endif

  SELECT (:'ADMIN_USERNAME' = '') as is_empty \gset
  \if :is_empty
    \warn 'ADMIN_USERNAME is empty'
    \set exit_error true
  \endif

  SELECT (:'ADMIN_PASSWORD' = '') as is_empty \gset
  \if :is_empty
    \warn 'ADMIN_PASSWORD is empty'
    \set exit_error true
  \endif

  SELECT (:'REGULAR_USERNAME' = '') as is_empty \gset
  \if :is_empty
    \warn 'REGULAR_USERNAME is empty'
    \set exit_error true
  \endif

  SELECT (:'REGULAR_PASSWORD' = '') as is_empty \gset
  \if :is_empty
    \warn 'REGULAR_PASSWORD is empty'
    \set exit_error true
  \endif

  \if :exit_error
    DO $$
    BEGIN
    RAISE EXCEPTION 'all required environment variables must not be empty';
    END;
    $$;
  \endif

  CREATE ROLE users;

  CREATE USER :STAT_USERNAME
    WITH ENCRYPTED PASSWORD :'STAT_PASSWORD'
    IN ROLE users;

  GRANT SELECT (id, created_at, updated_at)
    ON TABLE users_info, models, weights_info
    TO :STAT_USERNAME;

  CREATE USER :REGULAR_USERNAME
    WITH ENCRYPTED PASSWORD :'REGULAR_PASSWORD'
    IN ROLE users;

  GRANT SELECT, INSERT (
      id
      , username
      , email
      , fullname
      , password_hash
      , flags
      , blocked
    )
    ON TABLE users_info
    TO :REGULAR_USERNAME;

  GRANT SELECT, INSERT, UPDATE (id, title, owner_id)
    ON TABLE models
    TO :REGULAR_USERNAME;

  GRANT DELETE
    ON TABLE models
    TO :REGULAR_USERNAME;

  GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLE
      structures,
      layers,
      neurons,
      neuron_links,
      weights_info,
      neuron_offsets,
      link_weights
    TO :REGULAR_USERNAME;

  CREATE ROLE administrators;
  GRANT USAGE ON SCHEMA public TO administrators;
  GRANT SELECT, INSERT, UPDATE, DELETE
    ON ALL TABLES
    IN SCHEMA public
    TO administrators;

  ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLES TO administrators;

  CREATE USER :ADMIN_USERNAME
    WITH
    CREATEDB
    CREATEROLE
    ENCRYPTED PASSWORD :'ADMIN_PASSWORD'
    IN ROLE administrators;

  INSERT INTO migrations(id) VALUES (:'MIGRATION_ID');
\endif

COMMIT;
