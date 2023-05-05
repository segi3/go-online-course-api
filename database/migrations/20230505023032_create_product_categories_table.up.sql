CREATE TABLE product_categories (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT FK_product_categories_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)