CREATE TABLE users(
                      id SERIAL PRIMARY KEY NOT NULL UNIQUE ,
                      username VARCHAR(30) NOT NULL UNIQUE ,
                      password_hash VARCHAR(150) NOT NULL,
                      role VARCHAR(10) NOT NULL,
                      cart_id INT UNIQUE
);

CREATE TABLE products(
                         id SERIAL PRIMARY KEY NOT NULL UNIQUE,
                         category VARCHAR(30) NOT NULL ,
                         name VARCHAR(30) NOT NULL ,
                         color VARCHAR(30),
                         description VARCHAR(200) ,
                         price INT NOT NULL
);

CREATE TABLE cart(
                     id SERIAL PRIMARY KEY,
                     user_id INT UNIQUE
);

CREATE TABLE cart_products (
                               cart_id INT,
                               product_id INT,
                               PRIMARY KEY (cart_id, product_id),
                               FOREIGN KEY (cart_id) REFERENCES cart(id),
                               FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE order_from_cart (
                                 id SERIAL PRIMARY KEY,
                                 cart_id INT
);

CREATE TABLE order_product(
                              order_id INT,
                              product_id INT,
                              PRIMARY KEY (order_id, product_id),
                              FOREIGN KEY (order_id) REFERENCES order_from_cart(id),
                              FOREIGN KEY (product_id) REFERENCES products(id)
);