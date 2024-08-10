
-- +migrate Up
create table
  `user_transactions` (
    `id` VARCHAR(255) not null,
    `user_id` varchar(255) not null,
    `handling_type` varchar(255) not null,
    `transaction_type` varchar(255) not null,
    `status` varchar(255) not null,
    `amount` varchar(255) not null,
    `remarks` varchar(255) null,
    `balance_before` varchar(255) not null,
    `balance_after` varchar(255) not null,
    `created_at` timestamp not null default CURRENT_TIMESTAMP,
    primary key (`id`),
    KEY `user_transactions_users` (`user_id`),
    CONSTRAINT `user_transactions_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT
  ) ENGINE = InnoDB;
-- +migrate Down
