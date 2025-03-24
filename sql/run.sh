#!sh

set -e

if [ "$SQLDEF_ACTION" = "apply" ]; then
    postgresdef --user=manager --password=${POSTGRES_PASSWORD:-password} --host=${POSTGRES_HOST:-db} --port=${POSTGRES_PORT:-5432} ${POSTGRES_DATABASE:-local} < schema.sql
elif [ "$SQLDEF_ACTION" = "apply-enable-drop-table" ]; then
    postgresdef --user=manager --password=${POSTGRES_PASSWORD:-password} --host=${POSTGRES_HOST:-db} --port=${POSTGRES_PORT:-5432} ${POSTGRES_DATABASE:-local} < schema.sql --enable-drop-table
else
    postgresdef --user=manager --password=${POSTGRES_PASSWORD:-password} --host=${POSTGRES_HOST:-db} --port=${POSTGRES_PORT:-5432} ${POSTGRES_DATABASE:-local} < schema.sql --dry-run
fi