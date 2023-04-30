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

INSERT INTO users ("id", "name", "email", "password", "role")
VALUES (1, 'admin', 'admin2gmail.com', '1234', 1);

-- +goose Down
DROP TABLE users;