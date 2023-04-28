-- +goose Up
CREATE TABLE users (
  id int NOT NULL,
	"name" text,
	email text,
  "password" text,
  teams text[],
  requests text[],
	notifications text[],
  ban boolean,
  "role" bigint,
  PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;