CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL PRIMARY KEY,
    lastname VARCHAR(128),
    firstname VARCHAR(128),
    username VARCHAR(256) NOT NULL,
    password VARCHAR(256),
    passport VARCHAR(9),
    phone_number VARCHAR(13),
    age INTEGER,
    role VARCHAR(10)
);