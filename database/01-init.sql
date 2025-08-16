DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS user_preferences;


CREATE TABLE categories (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE products (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(100) NOT NULL,
    category_id VARCHAR(32) REFERENCES categories(id),
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    description VARCHAR(255),
    specs JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE user_preferences (
    id VARCHAR(32) PRIMARY KEY,
    user_session VARCHAR(255),
    preferred_categories VARCHAR(32)[],
    price_range_min DECIMAL(10,2),
    price_range_max DECIMAL(10,2),
    preferred_brands VARCHAR(255)[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX idx_products_category ON products(category_id);
CREATE INDEX idx_products_brand ON products(brand);
CREATE INDEX idx_products_stock ON products(stock_quantity);
CREATE INDEX idx_categories_name ON categories(name);