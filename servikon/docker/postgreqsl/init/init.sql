-- DROP SCHEMA records;

ALTER ROLE "user" NOSUPERUSER CREATEDB NOCREATEROLE INHERIT LOGIN;

CREATE schema IF NOT EXISTS records AUTHORIZATION postgres;
SET search_path TO records;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    uid         UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    username    varchar(30) NOT NULL UNIQUE,  
    email       varchar(30) NOT NULL UNIQUE,
    pass        varchar(70) NOT NULL
);

INSERT INTO users (username, email, pass) VALUES('maxim', 'maxim@ya.ru', '$2a$08$EZNHnalSY.BEFrxijfm7E.ueVYHCGEKA5UPTS0hznAmgL2FN8ESUW');
INSERT INTO users (username, email, pass) VALUES('foma', 'foma@ya.ru', '$2a$08$EZNHnalSY.BEFrxijfm7E.ueVYHCGEKA5UPTS0hznAmgL2FN8EW2A');
INSERT INTO users (username, email, pass) VALUES('admin', 'admin@ya.ru', '$2a$08$EZNHnalSY.BEFrxijfm7E.sdWYHCGEKA5UPTS0hznAmgL2FN8ESUW');
