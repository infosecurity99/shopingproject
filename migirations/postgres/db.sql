CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    firstname VARCHAR(30) NOT NULL,
    lastname VARCHAR(30) NOT NULL,
    email VARCHAR(30) NOT NULL,
    phone VARCHAR(30) NOT NULL
);
CREATE TABLE orders (
    id UUID PRIMARY KEY NOT NULL,
    amount INT NOT NULL,
    userid UUID REFERENCES users(id),
    createat TIMESTAMP NOT NULL 
);
CREATE TABLE products (
    id UUID PRIMARY KEY NOT NULL,
    price INT NOT NULL,
    name VARCHAR(30) NOT NULL
);
CREATE TABLE order_products (
    id UUID PRIMARY KEY NOT NULL,
    quantity INT NOT NULL,
    price INT NOT NULL, 
    productid UUID REFERENCES products(id),
    orderid UUID REFERENCES orders(id)
);
