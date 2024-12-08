#!sh

set -e

if [ "$SQLDEF_ACTION" = "apply" ]; then
    mysqldef -u manager -p${MYSQL_PASSWORD:-password} -h ${MYSQL_HOST:-db} -P ${MYSQL_PORT:-3306} ${MYSQL_DATABASE:-local} < schema.sql
elif [ "$SQLDEF_ACTION" = "apply-enable-drop-table" ]; then
    mysqldef -u manager -p${MYSQL_PASSWORD:-password} -h ${MYSQL_HOST:-db} -P ${MYSQL_PORT:-3306} ${MYSQL_DATABASE:-local} < schema.sql --enable-drop-table
else
    mysqldef -u manager -p${MYSQL_PASSWORD:-password} -h ${MYSQL_HOST:-db} -P ${MYSQL_PORT:-3306} ${MYSQL_DATABASE:-local} < schema.sql --dry-run
fi