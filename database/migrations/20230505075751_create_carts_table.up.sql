CREATE TABLE carts (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id`INT  NULL,
    `product_id`INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    PRIMARY KEY ( `id` ),
    INDEX idx_carts_user_id ( `user_id` ) ,
    INDEX idx_carts_product_id ( `product_id` ) ,
    CONSTRAINT FK_carts_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_carts_product_id FOREIGN KEY (`product_id`) REFERENCES products(`id`)  ON DELETE SET NULL
) 