
-- +migrate Up
CREATE TABLE todo (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users(id),
  Text TEXT NOT NULL,
  done BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate Down
Drop table todo;