CREATE TABLE order_details (
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_id` INT NULL,
    `price` INT NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    INDEX idx_order_details_created_by(`created_by`),
    INDEX idx_order_details_updated_by(`updated_by`),
    CONSTRAINT FK_order_details_product_id FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_order_details_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_order_details_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)