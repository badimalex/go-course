#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER docker;
    CREATE DATABASE url_shortener;
    GRANT ALL PRIVILEGES ON DATABASE url_shortener TO docker;

    CREATE TABLE urls (
        short_url character varying(50) PRIMARY KEY,
        destination character varying(255) NOT NULL,
        creation_time timestamp without time zone NOT NULL
    );


    CREATE UNIQUE INDEX urls_pkey ON urls(short_url text_ops);
EOSQL
