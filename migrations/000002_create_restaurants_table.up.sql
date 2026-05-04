CREATE TABLE IF NOT EXISTS restaurants (
    id SERIAL PRIMARY KEY,
    public_id VARCHAR(36) UNIQUE NOT NULL,
    owner_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    description TEXT,
    logo_url TEXT,
    banner_url TEXT,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_restaurants_public_id ON restaurants(public_id);
CREATE INDEX idx_restaurants_owner_id ON restaurants(owner_id);
