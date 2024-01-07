-- +goose Up

CREATE TABLE audit_trail (
  `id` VARCHAR(255) NOT NULL,
  `model_name` varchar(100) NOT NULL,
  `model_key` varchar(100) NOT NULL,
  `action` varchar(50) NOT NULL,
  `log` TEXT NOT NULL DEFAULT "",
  `remark` TEXT NOT NULL DEFAULT "",
  `ip_address` varchar(100) NOT NULL,
  `token_id` varchar(100) NOT NULL,
  `created_by` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE user (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) NOT NULL,
  `account_no` varchar(50) NOT NULL,
  `first_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `phone` varchar(50) NOT NULL,
  `phone_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(100) NOT NULL,
  `is_active` TINYINT(1) NOT NULL DEFAULT 0,
  `last_login_at` timestamp NULL DEFAULT NULL,
  `default_language` varchar(50) NOT NULL DEFAULT "en",
  `default_timezone` varchar(50) NOT NULL DEFAULT "Asia/Jakarta",
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_deleted_at` (`deleted_at`),
  INDEX `idx_email` (`email`),
  INDEX `idx_is_active` (`is_active`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE user_address (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `country_id` bigint(20) unsigned NOT NULL,
  `province` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `address_line_1` TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  `address_line_2` TEXT COLLATE utf8mb4_unicode_ci NOT NULL,
  `postcode` varchar(20) NOT NULL,
  `is_primary` TINYINT(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_deleted_at` (`deleted_at`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_is_primary` (`is_primary`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- +goose Down

DROP TABLE audit_trail;
DROP TABLE user;
DROP TABLE user_address;