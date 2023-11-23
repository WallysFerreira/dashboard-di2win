FROM postgres:latest

COPY ./dashboard.sql /dashboard.sql
COPY ./init.sh /docker-entrypoint-initdb.d/init.sh

ENV POSTGRES_HOST_AUTH_METHOD trust
EXPOSE 5432
