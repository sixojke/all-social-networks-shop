CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(32) UNIQUE,
    password VARCHAR(256),
    email VARCHAR(255) UNIQUE,
    balance NUMERIC(10, 2),
    last_visit_at TIMESTAMP NOT NULL,
    registered_at TIMESTAMP DEFAULT now()
);

CREATE TABLE verification(
    verified bool DEFAULT false,
    code VARCHAR(32),

    user_id BIGINT REFERENCES users(id)
);

CREATE TABLE sessions(
    refresh_token VARCHAR(256) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    
    user_id BIGINT REFERENCES users(id)
);

CREATE TABLE suppliers(
    id SERIAL PRIMARY KEY,
    total_profit NUMERIC(10, 2) NOT NULL,
    rating NUMERIC(1, 1) NOT NULL,
    total_reviews INTEGER NOT NULL,
    total_orders INTEGER NOT NULL, 
    avg_order_fulfillment_time INTEGER NOT NULL,

    user_id BIGINT REFERENCES users(id)
);

CREATE TABLE buyers(
    id SERIAL PRIMARY KEY,
    total_expenses NUMERIC(10, 2),
    total_orders INTEGER NOT NULL,

    user_id BIGINT REFERENCES users(id)
);

CREATE TABLE reviews(
    id SERIAL PRIMARY KEY,
    text TEXT,
    rating NUMERIC(1, 1),
    created_at TIMESTAMP DEFAULT now(),

    buyer_id INTEGER REFERENCES buyers(id),
    supplier_id INTEGER REFERENCES suppliers(id)
);

CREATE TABLE category(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    img_path VARCHAR(255)
);

CREATE TABLE subcategory(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    hold_time SMALLINT NOT NULL,

    category_id INTEGER REFERENCES category(id)
);

CREATE TABLE filters(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    quantity INTEGER NOT NULL,
    quantity_sales BIGINT NOT NULL,
    description TEXT,
    uploaded_at TIMESTAMP DEFAULT now(), 

    category_id INTEGER REFERENCES category(id),
    subcategory_id INTEGER REFERENCES subcategory(id)
);

CREATE TABLE orders(
    id BIGSERIAL PRIMARY KEY,
    amount NUMERIC(10, 2) NOT NULL,
    status VARCHAR(32) NOT NULL,
    devilery_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT now(),

    product_id BIGINT REFERENCES products(id),
    buyer_id INTEGER REFERENCES buyers(id)
);

CREATE TABLE project_reviews(
    id SERIAL PRIMARY KEY,
    rating DOUBLE PRECISION NOT NULL,
    text TEXT,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE project_info(
    id SERIAL PRIMARY KEY,
    total_rating DOUBLE PRECISION NOT NULL,
    total_reviews INTEGER NOT NULL
);

INSERT INTO project_info (total_rating, total_reviews) VALUES (0.0, 0);







INSERT INTO category (name, img_path) VALUES ('Twitter', 'path/to/img');
INSERT INTO category (name, img_path) VALUES ('Facebook', 'path/to/img');

INSERT INTO subcategory (name, hold_time, category_id) VALUES ('accounts', 60, 2);

INSERT INTO products (name, price, quantity, quantity_sales, description, category_id, subcategory_id) VALUES ('Facebook account', 99, 15000, 500, 'Хорошие аккаунты', 2, 1);