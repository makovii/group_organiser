-- +goose Up
CREATE TABLE statuses (
    id SERIAL PRIMARY KEY,
    status text
);

INSERT INTO statuses ("id", "status")
VALUES (1, 'wait'), (2, 'accept'), (3, 'reject'), (4, 'cancel');

-- +goose Down
DROP TABLE statuses;