CREATE TABLE `transactions` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `commerce_id` INT NOT NULL,
    `branch_id` INT NOT NULL,
    `campaign_id` INT NOT NULL,
    `amount` FLOAT NOT NULL,
    `points_earned` FLOAT,
    `cashback_earned` FLOAT,
    `created_at` DATETIME NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
    FOREIGN KEY (`commerce_id`) REFERENCES `commerces`(`id`),
    FOREIGN KEY (`branch_id`) REFERENCES `branches`(`id`),
    FOREIGN KEY (`campaign_id`) REFERENCES `campaigns`(`id`)
);
