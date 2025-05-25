CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    wishlist_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    url TEXT,
    image_url VARCHAR(255),
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    is_reserved BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_items_wishlist_id ON items(wishlist_id);

CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL UNIQUE,
    user_id INTEGER NOT NULL,
    reserved_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE reservations
ADD CONSTRAINT fk_reservations_item_id
FOREIGN KEY (item_id) REFERENCES items(id)
ON DELETE CASCADE;

ALTER TABLE items
ADD CONSTRAINT fk_items_reservation_id
FOREIGN KEY (id) REFERENCES reservations(item_id)
ON DELETE SET NULL;