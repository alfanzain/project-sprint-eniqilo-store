CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    customer_id VARCHAR(26),
    product_details JSONB,
    paid FLOAT,
    change FLOAT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transactions_customer_id ON transactions (customer_id);