#!/usr/bin/env bash
set -Eeo pipefail

docker_process_init_files() {
  psql=( docker_process_sql )

  echo
  local f
  for f; do
    export MIGRATION_ID="$(basename $f)"
    case "$f" in
      *.sh)
        if [ -x "$f" ]; then
          echo "$0: running $f"
          "$f"
        else
          echo "$0: sourcing $f"
          . "$f"
        fi
        ;;
      *.sql)    echo "$0: running $f"; docker_process_sql -f "$f"; echo ;;
      *.sql.gz) echo "$0: running $f"; gunzip -c "$f" | docker_process_sql; echo ;;
      *.sql.xz) echo "$0: running $f"; xzcat "$f" | docker_process_sql; echo ;;
      *)        echo "$0: ignoring $f" ;;
    esac
    echo
  done
}

docker_process_sql() {
  local query_runner=( \
    psql \
    -v ON_ERROR_STOP=1 \
    -v AUTOCOMMIT=off \
    -v MIGRATION_ID=$MIGRATION_ID \
    --no-psqlrc \
    "host=$POSTGRES_HOST \
     dbname=$POSTGRES_DB \
     user=$POSTGRES_USER \
     password=$POSTGRES_PASSWORD" \
  )

  PGHOST= PGHOSTADDR= "${query_runner[@]}" "$@"
}

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> \
              /dev/null && pwd )/init

echo $SCRIPT_DIR

docker_process_init_files $SCRIPT_DIR/*
