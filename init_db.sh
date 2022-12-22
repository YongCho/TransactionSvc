#!/usr/bin/env bash

# This script is used in initializating the database with the necessary tables.

set -e

echo "DB_HOST: ${DB_HOST}"
echo "DB_PORT: ${DB_PORT}"

# Wait until the database is ready to accept connections.
while ! pg_isready -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER} -d ${DB_DBNAME}; do
    echo "Waiting for DB..."
    sleep 1
done

# Execute the .sql files using psql (Postgres CLI) utility.
script_dir=db_script
scripts=$(ls ${script_dir})
connstr="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DBNAME}"
for f in ${scripts}; do
    echo "Executing ${f}"
    ec=0
    psql ${connstr} --echo-all -v ON_ERROR_STOP=1 -f ${script_dir}/${f} || ec=$?
    if [ $ec -ne 0 ]; then
        echo "psql returned error: ${ec}"
        exit 1
    fi
done

echo "Executed all DB scripts"
