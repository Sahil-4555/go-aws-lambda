
-- This Will Soft Delete The Order Details Associated With Orders of Customer Which will delete.

CREATE TRIGGER update_deleted_at_on_customer
AFTER UPDATE ON customer
FOR EACH ROW
BEGIN
    IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
        -- Update deleted_at in orders table
        UPDATE orders
        SET deleted_at = CURRENT_TIMESTAMP
        WHERE customer_id = NEW.id;

        -- Update deleted_at in order_details table
        UPDATE order_detail
        SET deleted_at = CURRENT_TIMESTAMP
        WHERE order_id IN (SELECT id FROM orders WHERE customer_id = NEW.id);
    END IF;
END;
