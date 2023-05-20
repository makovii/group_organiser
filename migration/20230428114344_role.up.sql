-- +goose Up
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role text
);

INSERT INTO roles ("id", "role")
VALUES (1, 'admin'), (2, 'manager'), (3, 'player') ;

