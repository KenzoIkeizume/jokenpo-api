#!/bin/sh
# wait-for-mysql.sh

until nc -z -v -w5 jokenpo-mysql 3306
do
    echo "Mysql is unavailable - sleeping"
    sleep 1
done

echo "Mysql is up - running"

exec "$@"
