--
-- Initialize database with models info
--
SELECT EXISTS (
  SELECT id FROM migrations WHERE id = :'MIGRATION_ID'
) as migrated \gset

\if :migrated
  \echo 'migration'   :MIGRATION_ID 'already exists, skipping'
\else
  \echo 'migration' :MIGRATION_ID 'does not exist'

  CREATE TABLE models (
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4()
    , title      VARCHAR(64) UNIQUE NOT NULL
    , created_at TIMESTAMP   DEFAULT NOW()
    , updated_at TIMESTAMP   DEFAULT NOW()
    , owner_id   UUID        NOT NULL
    , FOREIGN KEY (owner_id)
      REFERENCES users_info(id)
      ON DELETE CASCADE
  );

  CREATE TABLE structures (
    id         UUID        PRIMARY KEY DEFAULT uuid_generate_v4()
    , title    VARCHAR(64) NOT NULL
    , model_id UUID        NOT NULL
    , FOREIGN KEY (model_id)
      REFERENCES models(id)
      ON DELETE CASCADE
  );

  CREATE TABLE layers (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4()
    , layer_id        INT
    , limit_func      VARCHAR(64) NOT NULL
    , activation_func VARCHAR(64) NOT NULL
    , structure_id    UUID NOT NULL
    , FOREIGN KEY (structure_id)
      REFERENCES structures(id)
      ON DELETE CASCADE
  );

  CREATE TABLE neurons (
    id          UUID  PRIMARY KEY DEFAULT uuid_generate_v4()
    , neuron_id INT NOT NULL
    , layer_id  UUID  NOT NULL
    , FOREIGN KEY (layer_id)
      REFERENCES layers(id)
      ON DELETE CASCADE
  );

  CREATE TABLE neuron_links (
    id        UUID PRIMARY KEY DEFAULT uuid_generate_v4()
    , link_id INT  NOT NULL
    , from_id UUID
    , to_id   UUID NOT NULL
    , FOREIGN KEY (from_id)
      REFERENCES neurons(id)
      ON DELETE CASCADE
    , FOREIGN KEY (to_id)
      REFERENCES neurons(id)
      ON DELETE CASCADE
  );

  CREATE TABLE weights_info (
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4()
    , name         VARCHAR(64) NOT NULL
    , created_at   TIMESTAMP   DEFAULT NOW()
    , updated_at   TIMESTAMP   DEFAULT NOW()
    , structure_id UUID        NOT NULL
    , FOREIGN KEY (structure_id)
      REFERENCES structures(id)
      ON DELETE CASCADE
  );

  CREATE TABLE neuron_offsets (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4()
    , value           REAL NOT NULL
    , weights_info_id UUID NOT NULL
    , neuron_id       UUID NOT NULL
    , FOREIGN KEY (weights_info_id)
      REFERENCES weights_info(id)
      ON DELETE CASCADE
    , FOREIGN KEY (neuron_id)
      REFERENCES neurons(id)
      ON DELETE CASCADE
  );

  CREATE TABLE link_weights (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4()
    , value           REAL NOT NULL
    , weights_info_id UUID NOT NULL
    , link_id         UUID NOT NULL
    , FOREIGN KEY (weights_info_id)
      REFERENCES weights_info(id)
      ON DELETE CASCADE
    , FOREIGN KEY (link_id)
      REFERENCES neuron_links(id)
      ON DELETE CASCADE
  );

  INSERT INTO migrations(id) VALUES (:'MIGRATION_ID');
\endif

COMMIT;
