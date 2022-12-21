#!/usr/bin/env bash

set -e

echo "DB_HOST: ${DB_HOST}"
echo "DB_PORT: ${DB_PORT}"

while ! pg_isready -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER} -d ${DB_DBNAME}; do
    echo "Waiting for DB..."
    sleep 1
done

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
