-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
	"name" text,
	email text,
  "password" text,
  teams text[],
  requests text[],
	notifications text[],
  ban boolean,
  "role" bigint
);

INSERT INTO users ("id", "name", "email", "password", "role")
VALUES (0, 'admin', 'admin@gmail.com', '1234', 1);

