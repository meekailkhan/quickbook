-- +goose Up

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    photo VARCHAR(255),
    gender VARCHAR(255),
    phone VARCHAR(10),
    civil REAL,
    "password" VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS lessors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(1000) NOT NULL,
    city VARCHAR(255) NOT NULL,
    "state" VARCHAR(255) NOT NULL,
    longitude VARCHAR(100),
    latitude VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS inventories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lessor_id UUID NOT NULL REFERENCES lessors(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    rent_per_day INT NOT NULL,
    rent_per_month INT,
    available_units INT,
    health REAL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS inventory_photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    inventory_id UUID NOT NULL REFERENCES inventories(id) ON DELETE CASCADE,
    photos TEXT[]
);

CREATE TABLE IF NOT EXISTS pre_owned (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    inventory_id UUID NOT NULL REFERENCES inventories(id) ON DELETE CASCADE
);
-- +goose Down

DROP TABLE IF EXISTS pre_owned;
DROP TABLE IF EXISTS inventory_photos;
DROP TABLE IF EXISTS inventories;
DROP TABLE IF EXISTS lessors; 
DROP TABLE IF EXISTS users;