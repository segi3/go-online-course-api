CREATE TABLE forgot_passwords (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NULL,
    `valid` BOOLEAN NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `expired_at` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT FK_forgot_passwords_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE SET NULL
)