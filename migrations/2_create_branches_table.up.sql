CREATE TABLE `branches` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `commerce_id` INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `address` TEXT,
    FOREIGN KEY (`commerce_id`) REFERENCES `commerces`(`id`)
);
