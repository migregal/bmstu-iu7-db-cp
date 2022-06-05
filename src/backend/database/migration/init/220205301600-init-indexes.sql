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

  CREATE UNIQUE INDEX idx_struct_model ON structures(model_id);

  CREATE INDEX idx_layer_struct ON layers(structure_id);

  CREATE INDEX idx_neurons_layer ON neurons(layer_id);

  CREATE INDEX idx_link_from_neurons ON neuron_links(from_id);

  CREATE INDEX idx_link_to_neurons ON neuron_links(to_id);

  CREATE INDEX idx_weights_struct ON weights_info(structure_id);

  CREATE INDEX idx_off_weights ON neuron_offsets(weights_info_id);

  CREATE INDEX idx_lw_weights ON link_weights(weights_info_id);

  INSERT INTO migrations(id) VALUES (:'MIGRATION_ID');
\endif

COMMIT;
