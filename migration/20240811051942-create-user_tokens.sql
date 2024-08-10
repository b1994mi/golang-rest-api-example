
-- +migrate Up
CREATE TABLE
  `user_tokens` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `user_id` varchar(255) NOT NULL,
    `token` varchar(255) NOT NULL,
    `exp_date_str` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `user_tokens_users` (`user_id`),
    CONSTRAINT `user_tokens_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT
  ) ENGINE = InnoDB;
-- +migrate Down
