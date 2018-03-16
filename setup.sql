CREATE DATABASE `zy_todo`;

USE `zy_todo`;

CREATE TABLE `task` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(30) NOT NULL,
  `done` TINYINT(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `task` VALUES (1, 'GO', 0);

CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(30) NOT NULL unique,
    `password` varchar(30) NOT NULL,
    PRIMARY KEY (`Id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `user` VALUES (1, 'hello', 'world');
INSERT INTO `user` VALUES (2, 'yu', 'password');