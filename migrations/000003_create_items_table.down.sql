ALTER TABLE items
DROP CONSTRAINT IF EXISTS fk_items_reservation_id;

ALTER TABLE reservations
DROP CONSTRAINT IF EXISTS fk_reservations_item_id;

DROP TABLE IF EXISTS reservations;

DROP INDEX IF EXISTS idx_items_wishlist_id;

DROP TABLE IF EXISTS items;