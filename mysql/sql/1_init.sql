USE echo_local;

-- 1. usersテーブル作成

CREATE TABLE `users` (
    `id` BIGINT AUTO_INCREMENT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `birth_year` INT(11) NOT NULL,
    `gender` TINYINT(1) NOT NULL,
    `avatar` VARCHAR(255) NOT NULL,
    `is_delete` TINYINT(1) NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY(`id`)
);

-- 2. postsテーブル作成

CREATE TABLE `posts` (
    `id` BIGINT AUTO_INCREMENT NOT NULL,
    `user_id` BIGINT NOT NULL,
    `content` TEXT NOT NULL,
    `is_delete` TINYINT(1) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY(`id`)
);