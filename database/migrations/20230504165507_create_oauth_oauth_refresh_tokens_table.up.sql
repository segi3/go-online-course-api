CREATE TABLE oauth_refresh_tokens (
    `id` INT NOT NULL AUTO_INCREMENT,
    `oauth_access_token_id` INT NULL,
    `user_id` INT NULL,
    `token` VARCHAR(255) NOT NULL,
    `scope` VARCHAR(255) NOT NULL,
    `expired_at` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    UNIQUE KEY oauth_refresh_tokens_client_id_unique(`token`),

    CONSTRAINT FK_oauth_refresh_tokens_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_oauth_refresh_tokens_oauth_client_id FOREIGN KEY (`oauth_access_token_id`) REFERENCES oauth_access_tokens(`id`) ON DELETE SET NULL
)