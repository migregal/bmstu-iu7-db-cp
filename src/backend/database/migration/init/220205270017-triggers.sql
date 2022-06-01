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

  CREATE FUNCTION user_info_preupdate()
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

  CREATE FUNCTION model_preupdate()
  RETURNS trigger AS
  $$
  BEGIN
      NEW.UPDATED_AT = NOW();
      RETURN NEW;
  END;
  $$
  LANGUAGE 'plpgsql';

  CREATE TRIGGER updt_model BEFORE UPDATE
  ON models
  FOR ROW
  EXECUTE PROCEDURE model_preupdate();

  CREATE FUNCTION weights_info_preupdate()
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

  CREATE EXTENSION PLPYTHON3U;

  CREATE FUNCTION remove_model_from_cache_py()
  RETURNS TRIGGER
  AS $$
      import os
      import tarantool

      host = os.getenv('CACHE_DB_HOST')
      if host is None:
        return

      port = os.getenv('CACHE_DB_PORT')
      if host is None:
        return

      user = os.getenv('CACHE_DB_USERNAME')
      if user is None:
        return

      pwd = os.getenv('CACHE_DB_PASSWORD')
      if pwd is None:
        return

      model_space = os.getenv('CACHE_DB_MODEL_SPACE')
      if pwd is None:
        return

      try:
        conn = tarantool.connect(host, port,
                                  user=user, password=pwd)
        conn.space(model_space).delete(TD["old"]["id"])
      except:
        pass
  $$ LANGUAGE PLPYTHON3U;

  CREATE TRIGGER rm_model_upd AFTER UPDATE
  ON models FOR ROW
  EXECUTE PROCEDURE remove_model_from_cache_py();

  CREATE TRIGGER rm_model_del AFTER DELETE
  ON models FOR ROW
  EXECUTE PROCEDURE remove_model_from_cache_py();

  CREATE TRIGGER rm_model_s_upd AFTER UPDATE
  ON structures FOR ROW
  EXECUTE PROCEDURE remove_model_from_cache_py();

  CREATE TRIGGER rm_model_s_del AFTER DELETE
  ON structures FOR ROW
  EXECUTE PROCEDURE remove_model_from_cache_py();

  CREATE TRIGGER rm_model_wi_upd AFTER UPDATE
  ON weights_info FOR ROW
  EXECUTE PROCEDURE remove_model_from_cache_py();

  CREATE TRIGGER rm_model_wi_del AFTER DELETE
  ON weights_info FOR ROW
  EXECUTE PROCEDURE remove_model_from_cache_py();

  CREATE FUNCTION remove_weight_from_cache_py()
  RETURNS TRIGGER
  AS $$
      import os
      import tarantool

      host = os.getenv('CACHE_DB_HOST')
      if host is None:
        return

      port = os.getenv('CACHE_DB_PORT')
      if host is None:
        return

      user = os.getenv('CACHE_DB_USERNAME')
      if user is None:
        return

      pwd = os.getenv('CACHE_DB_PASSWORD')
      if pwd is None:
        return

      weight_space = os.getenv('CACHE_DB_WEIGHT_SPACE')
      if pwd is None:
        return

      try:
        conn = tarantool.connect(host, port,
                                user=user, password=pwd)
        conn.space(weight_space).delete(TD["old"]["id"])
      except:
        pass
  $$ LANGUAGE PLPYTHON3U;

  CREATE TRIGGER rm_weight_del AFTER DELETE
  ON models FOR ROW
  EXECUTE PROCEDURE remove_weight_from_cache_py();

  CREATE TRIGGER rm_weight_s_upd AFTER UPDATE
  ON structures FOR ROW
  EXECUTE PROCEDURE remove_weight_from_cache_py();

  CREATE TRIGGER rm_weight_s_del AFTER DELETE
  ON structures FOR ROW
  EXECUTE PROCEDURE remove_weight_from_cache_py();

  CREATE TRIGGER rm_weight_wi_upd AFTER UPDATE
  ON weights_info FOR ROW
  EXECUTE PROCEDURE remove_weight_from_cache_py();

  CREATE TRIGGER rm_weight_wi_del AFTER DELETE
  ON weights_info FOR ROW
  EXECUTE PROCEDURE remove_weight_from_cache_py();

  INSERT INTO migrations(id) VALUES (:'MIGRATION_ID');
\endif

COMMIT;
