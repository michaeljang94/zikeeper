CREATE DATABASE IF NOT EXISTS zikeeper;

USE zikeeper;
CREATE TABLE IF NOT EXISTS users(
    id varchar(36) DEFAULT (UUID()), 
    name varchar(255),
    score int,
    PRIMARY KEY (id),
    CONSTRAINT uc_name UNIQUE (name)
);