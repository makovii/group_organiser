-- +goose Up
CREATE TABLE types (
    id int NOT NULL,
    type text,
    PRIMARY KEY(id)
);

INSERT INTO types ("id", "type")
VALUES (1, 'registration'), (2, 'joinTeam'), (3, 'leaveTeam');

-- +goose Down
DROP TABLE types;