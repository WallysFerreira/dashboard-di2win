#!/bin/bash

createdb -U postgres database
psql -U postgres -c "CREATE ROLE admin"
pg_restore -U postgres -d database ../dashboard.sql
