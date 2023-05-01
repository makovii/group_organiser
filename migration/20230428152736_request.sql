-- +goose Up
CREATE TABLE requests (
  id SERIAL PRIMARY KEY,
	"from" bigint,
	"to" bigint,
	status_id bigint,
	type_id   bigint,
  CONSTRAINT fk_status_id
    FOREIGN KEY(status_id) 
      REFERENCES statuses(id),
  CONSTRAINT fk_type_id
    FOREIGN KEY(type_id) 
      REFERENCES types(id)
  
);

-- +goose Down
DROP TABLE requests;