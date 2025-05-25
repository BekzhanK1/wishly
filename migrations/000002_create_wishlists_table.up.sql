CREATE TABLE wishlists (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    slug VARCHAR(255) UNIQUE NOT NULL,
    is_public BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_wishlists_user_id ON wishlists(user_id);

ALTER TABLE wishlists
ADD CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
ON DELETE CASCADE ON UPDATE CASCADE;

CREATE OR REPLACE FUNCTION update_wishlist_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_wishlist_updated_at
BEFORE UPDATE ON wishlists
FOR EACH ROW
EXECUTE PROCEDURE update_wishlist_updated_at_column();
