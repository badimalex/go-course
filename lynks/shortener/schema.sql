-- Table Definition ----------------------------------------------

CREATE TABLE urls (
    short_url character varying(50) PRIMARY KEY,
    destination character varying(255) NOT NULL,
    creation_time timestamp without time zone NOT NULL
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX urls_pkey ON urls(short_url text_ops);
