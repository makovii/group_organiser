-- +goose Up
CREATE TABLE types (
    id SERIAL PRIMARY KEY,
    type text
);

INSERT INTO types ("id", "type")
VALUES (1, 'registration'), (2, 'joinTeam'), (3, 'leaveTeam');

-- +goose Down
DROP TABLE types;