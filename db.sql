CREATE TABLE users (
    ID UUID PRIMARY KEY NOT NULL,
    FirstName VARCHAR(30) NOT NULL,
    LastName VARCHAR(30) NOT NULL,
    Email VARCHAR(30) NOT NULL,
    Phone VARCHAR(30) NOT NULL
);
CREATE TABLE orders (
    ID UUID PRIMARY KEY NOT NULL,
    Amount INT NOT NULL,
    User_id UUID REFERENCES users(ID),
    create_at TIMESTAMP NOT NULL
);


CREATE TABLE products (
    ID UUID PRIMARY KEY NOT NULL,
    Price INT NOT NULL,
    Name VARCHAR(30) NOT NULL
);

CREATE TABLE order_products (
    ID UUID PRIMARY KEY NOT NULL,
    Quantity INT NOT NULL,
    Price INT NOT NULL, 
    Product_id UUID REFERENCES products(ID),
    Order_id UUID REFERENCES orders(ID)
);
