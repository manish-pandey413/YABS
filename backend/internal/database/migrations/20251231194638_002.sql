-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
  id SERIAL PRIMARY KEY NOT NULL,
  owner_username TEXT NOT NULL,
  content TEXT NOT NULL,
  comments TEXT[],
  FOREIGN KEY(owner_username) REFERENCES users(username)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
