# /bin/usr/sh
echo list all tables:
sqlite currency.dbf < tables.sql
echo list schema movs:
sqlite currency.dbf < schema.sql

