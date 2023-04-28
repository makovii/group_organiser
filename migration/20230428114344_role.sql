-- +goose Up
CREATE TABLE roles (
    id int NOT NULL,
    role text,
    PRIMARY KEY(id)
);

INSERT INTO roles ("id", "role")
VALUES (1, 'admin'), (2, 'manager'), (3, 'player') ;

-- +goose Down
DROP TABLE roles;