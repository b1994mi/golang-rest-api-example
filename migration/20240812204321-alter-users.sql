
-- +migrate Up
ALTER TABLE
  `users`
ADD COLUMN
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
-- +migrate Down
