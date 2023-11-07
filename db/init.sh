#!/bin/bash

createdb -U postgres -T template0 database
pg_restore -U postgres -d database dashboard.sql
