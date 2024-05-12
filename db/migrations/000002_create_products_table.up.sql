CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30),
    sku VARCHAR(30),
    category VARCHAR(30),
    image_url VARCHAR(255),
    notes VARCHAR(200),
    price FLOAT,
    stock INT,
    location VARCHAR(200),
    is_available BOOLEAN,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_products_id ON products (id);
CREATE INDEX IF NOT EXISTS idx_products_name ON products (name);
CREATE INDEX IF NOT EXISTS idx_products_sku ON products (sku);
CREATE INDEX IF NOT EXISTS idx_products_category ON products (category);
CREATE INDEX IF NOT EXISTS idx_products_is_available ON products (is_available);