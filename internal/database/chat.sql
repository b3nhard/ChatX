DROP DATABASE IF EXISTS chatty;

CREATE DATABASE chatty;

USE chatty;

CREATE TABLE users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    INDEX(id)
);

-- CREATE TABLE groups (
--    `id` INT PRIMARY KEY AUTO_INCREMENT,
--    `name` VARCHAR(255) NOT NULL
-- );

-- CREATE TABLE user_group (
--    `user_id` INT,
--    `group_id` INT,
--    PRIMARY KEY (user_id, group_id),
--    FOREIGN KEY (user_id) REFERENCES users(user_id),
--    FOREIGN KEY (group_id) REFERENCES groups(group_id)
-- );


-- CREATE TABLE messages (
--    `id` INT PRIMARY KEY AUTO_INCREMENT,
--    `sender_id` INT NOT NULL,
--    `receiver_id` INT, -- Nullable for group messages
--    `group_id` INT,    -- Nullable for one-to-one messages
--    `content` TEXT NOT NULL,
--    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--    FOREIGN KEY (sender_id) REFERENCES users(id),
--    FOREIGN KEY (receiver_id) REFERENCES users(id),
--    FOREIGN KEY (group_id) REFERENCES groups(id)
-- );


