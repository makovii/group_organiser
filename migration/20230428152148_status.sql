-- +goose Up
CREATE TABLE statuses (
    id int NOT NULL,
    status text,
    PRIMARY KEY(id)
);

INSERT INTO statuses ("id", "status")
VALUES (1, 'wait'), (2, 'accept'), (3, 'reject'), (4, 'cancel');

-- +goose Down
DROP TABLE statuses;