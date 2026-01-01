-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY NOT NULL,
  username VARCHAR(26) NOT NULL UNIQUE,
  email VARCHAR(32) NOT NULL UNIQUE,
  password CHAR(64) NOT NULL
);


CREATE INDEX idx_users_username on users(username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
