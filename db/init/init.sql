CREATE DATABASE vul_db;

DROP SCHEMA IF EXISTS vul_schema CASCADE;
CREATE SCHEMA vul_schema;

DROP TABLE IF EXISTS vul_schema.users;
CREATE TABLE vul_schema.users
(
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  profile_link VARCHAR(255) NOT NULL
);
