--
-- Initialize database with basic triggers
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

    CREATE OR REPLACE FUNCTION user_info_preupdate()
    RETURNS trigger AS
    $$
    BEGIN
        NEW.UPDATED_AT = NOW();
        RETURN NEW;
    END;
    $$
    LANGUAGE 'plpgsql';

    CREATE TRIGGER updt_user_info BEFORE UPDATE
    ON users_info
    FOR ROW
    EXECUTE PROCEDURE user_info_preupdate();


    CREATE OR REPLACE FUNCTION model_info_preupdate()
    RETURNS trigger AS
    $$
    BEGIN
        NEW.UPDATED_AT = NOW();
        RETURN NEW;
    END;
    $$
    LANGUAGE 'plpgsql';

    CREATE TRIGGER updt_model_info BEFORE UPDATE
    ON models
    FOR ROW
    EXECUTE PROCEDURE model_info_preupdate();

    CREATE OR REPLACE FUNCTION weights_info_preupdate()
    RETURNS trigger AS
    $$
    BEGIN
        NEW.UPDATED_AT = NOW();
        RETURN NEW;
    END;
    $$
    LANGUAGE 'plpgsql';

    CREATE TRIGGER updt_weights_info BEFORE UPDATE
    ON weights_info
    FOR ROW
    EXECUTE PROCEDURE weights_info_preupdate();

    INSERT INTO migrations(id) VALUES (:'MIGRATION_ID');
\endif

COMMIT;
