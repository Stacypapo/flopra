-- =======================
-- Users
-- =======================
CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT,
    role TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =======================
-- Orders
-- =======================
CREATE TABLE orders (
    order_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    status TEXT NOT NULL,
    total_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    shipping_address TEXT
    warehouse_id BIGINT REFERENCES warehouses(warehouse_id)
);

-- =======================
-- Products
-- =======================
CREATE TABLE products (
    product_id BIGSERIAL PRIMARY KEY,
    sku TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    url TEXT
);

-- =======================
-- Order items
-- =======================
CREATE TABLE order_items (
    order_id BIGINT NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (order_id, product_id)
);

-- =======================
-- Cart items
-- =======================
CREATE TABLE cart_items (
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (user_id, product_id)
);

-- =======================
-- User queries (AI logs)
-- =======================
CREATE TABLE user_queries (
    product_id BIGINT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    query TEXT NOT NULL,
    model TEXT,
    rating INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (product_id, user_id, created_at)
);

-- =======================
-- Bouquets (composition)
-- =======================
CREATE TABLE bouquet_items (
    product_id_parent BIGINT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    product_id_child BIGINT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (product_id_parent, product_id_child)
);

-- =======================
-- Tags
-- =======================
CREATE TABLE tags (
    tag_id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT,
    color_hex CHAR(7) -- #RRGGBB
);

-- =======================
-- Product-Tag mapping
-- =======================
CREATE TABLE product_tags (
    tag_id BIGINT NOT NULL REFERENCES tags(tag_id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    PRIMARY KEY (tag_id, product_id)
);

-- =======================
-- Warehouses
-- =======================
CREATE TABLE warehouses (
    warehouse_id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT
);

-- =======================
-- Inventory
-- =======================
CREATE TABLE inventory (
    warehouse_id BIGINT NOT NULL REFERENCES warehouses(warehouse_id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity >= 0),
    PRIMARY KEY (warehouse_id, product_id)
);
