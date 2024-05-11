CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(32) UNIQUE,
    password VARCHAR(256),
    email VARCHAR(255) UNIQUE,
    balance NUMERIC(10, 2),
    role VARCHAR(32) DEFAULT 'user',
    telegram_id BIGINT UNIQUE,
    last_visit_at TIMESTAMP NOT NULL,
    registered_at TIMESTAMP DEFAULT now()
);

CREATE TABLE two_fa(
    is_active BOOLEAN NOT NULL,
    secret_code VARCHAR(255),
    user_id BIGINT UNIQUE NOT NULL REFERENCES users(id)
);

CREATE TABLE tg_user (
    user_id BIGINT PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,

    CONSTRAINT fk_user 
        FOREIGN KEY (user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE TABLE verification(
    verified bool DEFAULT false,
    code VARCHAR(32),

    user_id BIGINT PRIMARY KEY REFERENCES users(id)
);

CREATE TABLE password_recovery(
    secret_code VARCHAR(255) UNIQUE NOT NULL,
    recovery_time TIMESTAMP NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id)
);

CREATE TABLE sessions(
    refresh_token VARCHAR(256) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    
    user_id BIGINT PRIMARY KEY REFERENCES users(id)
);

CREATE TABLE suppliers(
    total_profit NUMERIC(10, 2) NOT NULL,
    rating NUMERIC(1, 1) NOT NULL,
    total_reviews INTEGER NOT NULL,
    total_orders INTEGER NOT NULL, 
    avg_order_fulfillment_time INTEGER NOT NULL,

    id BIGINT PRIMARY KEY REFERENCES users(id)
);

CREATE TABLE buyers(
    total_expenses NUMERIC(10, 2),
    total_orders INTEGER NOT NULL,

    id BIGINT PRIMARY KEY REFERENCES users(id)
);

CREATE TABLE banned_users(
    user_id BIGINT UNIQUE REFERENCES users(id),
    status BOOLEAN NOT NULL,
    banned_at TIMESTAMP DEFAULT now()
);

CREATE TABLE bind_telegram(
    code VARCHAR(255) NOT NULL,
    user_id BIGINT PRIMARY KEY REFERENCES users(id)
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
    min_hold_time SMALLINT NOT NULL,

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

    supplier_id INTEGER REFERENCES suppliers(id),
    category_id INTEGER REFERENCES category(id),
    subcategory_id INTEGER REFERENCES subcategory(id)
);

CREATE TABLE cart(
    user_id BIGINT REFERENCES users(id) NOT NULL,
    product_id BIGINT REFERENCES products(id) NOT NULL,
    quantity INTEGER NOT NULL
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

CREATE TABLE admin_logs(
    user_id BIGINT NOT NULL REFERENCES users(id),
    message TEXT,
    created_at TIMESTAMP DEFAULT now()
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

CREATE TABLE referral_system(
    referral_code VARCHAR(255) UNIQUE NOT NULL,
    total_visitors BIGINT NOT NULL DEFAULT 0,
    description VARCHAR(127),
    created_at TIMESTAMP DEFAULT now()
);

INSERT INTO project_info (total_rating, total_reviews) VALUES (0.0, 0);


INSERT INTO users (username, password, email, balance, role, last_visit_at) VALUES ('admin1234', '6d792d73616c747b902e6ff1db9f560443f2048974fd7d386975b0', 'admin@gmail.com', 0, 'admin', now());
INSERT INTO verification (verified, user_id) VALUES (true, 1);

INSERT INTO users (username, password, email, balance, last_visit_at) VALUES ('test-user', '6d792d73616c74c56486f8b638f63e04251d0c8ab0b4fbfee8e06b', 'test-user@gmail.com', 0, now());
INSERT INTO verification (verified, user_id) VALUES (true, 2);