DROP TRIGGER IF EXISTS set_wishlist_updated_at ON wishlists;

DROP FUNCTION IF EXISTS update_wishlist_updated_at_column;

ALTER TABLE wishlists
DROP CONSTRAINT IF EXISTS fk_user;

DROP INDEX IF EXISTS idx_wishlists_user_id;

DROP TABLE IF EXISTS wishlists;
