-- +goose Up
CREATE TABLE teams (
  id int NOT NULL,
	"name" text,
	manager_id bigint,
	member_ids text[],
  PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE teams;