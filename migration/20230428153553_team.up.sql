-- +goose Up
CREATE TABLE teams (
  id SERIAL PRIMARY KEY,
	"name" text,
	manager_id bigint,
	member_ids text[]
);

