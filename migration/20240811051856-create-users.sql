
-- +migrate Up
CREATE TABLE
  `users` (
    `id` varchar(255) NOT NULL,
    `wallet` decimal NOT NULL DEFAULT 0,
    `first_name` varchar(255) DEFAULT NULL,
    `last_name` varchar(255) DEFAULT NULL,
    `phone_number` varchar(255) NOT NULL,
    `address` text,
    `pin` varchar(255) NOT NULL,
    `is_deleted` tinyint NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_index_phone_number` (`phone_number`)
  ) ENGINE = InnoDB;
-- +migrate Down
