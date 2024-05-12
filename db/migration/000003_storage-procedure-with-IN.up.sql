CREATE PROCEDURE GetAllCustomer(
    IN limit_value INT
)
BEGIN
    SELECT 
        * 
    FROM customer 
    ORDER BY customer_name
    LIMIT limit_value;
END
