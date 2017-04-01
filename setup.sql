--CREATE DATABASE photo;
CREATE TABLE photographers (
  id integer PRIMARY KEY,
  name varchar(40) NOT NULL,
  surname varchar(40) NOT NULL,
  vk_url varchar(40),
  instagram_url varchar(40)
);
CREATE TABLE users (
  id integer PRIMARY KEY,
  email varchar(40) NOT NULL
);
CREATE TABLE tags (
  id integer PRIMARY KEY,
  Name varchar(40) NOT NULL
);
CREATE TABLE attachments (
  id integer PRIMARY KEY,
  description varchar(40)
);
