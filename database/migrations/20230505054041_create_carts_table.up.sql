CREATE TABLE carts (
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_id` INT NULL,
    `user_id` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT FK_carts_product_id FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_carts_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE SET NULL
)