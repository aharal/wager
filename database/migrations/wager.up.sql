CREATE DATABASE IF NOT EXISTS wagerDB;
CREATE TABLE `wager`(
    `wager_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `total_wager_value` INT(10) UNSIGNED NOT NULL,
    `odds` INT(10) UNSIGNED NOT NULL,
    `selling_percentage` INT(3) UNSIGNED NOT NULL,
    `selling_price` INT(10,2) UNSIGNED NOT NULL,
    `current_selling_price` INT(10,2) UNSIGNED NOT NULL,
    `percentage_sold` INT(10) UNSIGNED NULL,
    `amount_sold` INT(10,2) UNSIGNED NULL,
    `placed_at` TIMESTAMP NULL,
    PRIMARY KEY(`wager_id`),
) ENGINE = InnoDB DEFAULT CHARSET = latin1;

CREATE TABLE `buy_wager`(
    `purchase_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `wager_id` INT(10) UNSIGNED NOT NULL,
    `buying_price` INT(10,2) UNSIGNED NULL,
    `bought_at` TIMESTAMP NULL,
    PRIMARY KEY(`purchase_id`),
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
