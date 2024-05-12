CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    phone_number VARCHAR(16),
    role SMALLINT DEFAULT 2,
    password VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_phone_number_role UNIQUE (phone_number, role)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_users_phone_number_role ON users (phone_number, role);