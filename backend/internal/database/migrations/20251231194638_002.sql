-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
  id SERIAL PRIMARY KEY NOT NULL,
  owner_id INT NOT NULL,
  content TEXT NOT NULL,
  FOREIGN KEY(owner_id) REFERENCES users(user_id)
);

ALTER TABLE posts
ADD COLUMN comments TEXT[]
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
