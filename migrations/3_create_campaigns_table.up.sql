CREATE TABLE `campaigns` (
     `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
     `commerce_id` INT NOT NULL,
     `branch_id` INT NOT NULL,
     `start_date` DATETIME NOT NULL,
     `end_date` DATETIME NOT NULL,
     `multiplier` FLOAT NOT NULL,
     FOREIGN KEY (`commerce_id`) REFERENCES `commerces`(`id`),
     FOREIGN KEY (`branch_id`) REFERENCES `branches`(`id`)
);