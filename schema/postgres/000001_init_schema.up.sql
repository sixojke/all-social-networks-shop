CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(32),
    password VARCHAR(256),
    email VARCHAR(255),
    balance NUMERIC(10, 2),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE suppliers(
    id SERIAL PRIMARY KEY,
    rating NUMERIC(1, 1) NOT NULL,
    total_reviews INTEGER NOT NULL,
    total_orders INTEGER NOT NULL, 
    avg_order_fulfillment_time NOT NULL,

    user_id BIGINT REFERENCES users(id)
);

CREATE TABLE buyers(
    id SERIAL PRIMARY KEY,
    total_expenses NUMERIC(10, 2),
    total_orders INTEGER NOT NULL,

    user_id BIGINT REFERENCES users(id),
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
    name VARCHAR(255) NOT NULL
);

CREATE TABLE subcategory(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    hold_time SMALLINT NOT NULL,

    category_id INTEGER REFERENCES category(id)
)

CREATE TABLE filters(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    quantity INTEGER NOT NULL,
    img_path VARCHAR(255),
    description TEXT,
    uploaded_at TIMESTAMP DEFAULT now(), 

    category_id INTEGER REFERENCES category(id),
    filter_id INTEGER REFERENCES filters(id)
);

CREATE TABLE orders(
    id BIGSERIAL PRIMARY KEY,
    status VARCHAR(32) NOT NULL
    devilery_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT now(),

    product_id BIGINT REFERENCES products(id),
    buyer_id INTEGER REFERENCES buyers(id)
);