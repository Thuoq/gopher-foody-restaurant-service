CREATE TABLE IF NOT EXISTS foods (
    id SERIAL PRIMARY KEY,
    public_id VARCHAR(36) UNIQUE NOT NULL,
    restaurant_id INTEGER NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(12,2) NOT NULL,
    quantity INTEGER DEFAULT 0,
    status VARCHAR(50) DEFAULT 'available',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_foods_public_id ON foods(public_id);
CREATE INDEX idx_foods_restaurant_id ON foods(restaurant_id);
