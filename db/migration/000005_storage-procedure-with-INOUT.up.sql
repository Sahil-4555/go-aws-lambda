CREATE PROCEDURE IncProductPriceByAmount(
    INOUT amount FLOAT,
    IN product_id INT
)
BEGIN
    UPDATE product
    SET price = price + amount 
    WHERE id = product_id;

    -- Retrieve the updated price
    SELECT price INTO amount FROM product WHERE id = product_id;
END;

