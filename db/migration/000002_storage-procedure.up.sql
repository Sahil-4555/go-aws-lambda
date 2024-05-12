    CREATE PROCEDURE GetCustomer()
    BEGIN
        SELECT 
            customer_name, 
            city 
        FROM customer
        ORDER BY city;
    END
