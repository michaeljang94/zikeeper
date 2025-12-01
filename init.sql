CREATE DATABASE IF NOT EXISTS zikeeper;

USE zikeeper;
CREATE TABLE IF NOT EXISTS users(
    id varchar(36) DEFAULT (UUID()), 
    name varchar(255),
    score int,
    user_name varchar(255),
    password varchar(255),
    PRIMARY KEY (id),
    CONSTRAINT uc_name UNIQUE (name)
);

INSERT INTO users
VALUES ("1cb4a8ea-3cfb-4283-b400-2e21b7668266", "muone", 0, "muone", "1234");

INSERT INTO users
VALUES ("3564dc6f-b8d1-422e-b02d-1465f7acdc75", "guksoo", 0, "guksoo", "1234");