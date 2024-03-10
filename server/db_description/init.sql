CREATE TABLE users(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE ,
    username VARCHAR(30) NOT NULL UNIQUE ,
    password_hash VARCHAR(150) NOT NULL,
    role VARCHAR(10) NOT NULL
);

CREATE TABLE products(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    category VARCHAR(30) NOT NULL ,
    name VARCHAR(30) NOT NULL ,
    color VARCHAR(30),
    description VARCHAR(200) ,
    price INT NOT NULL
)

