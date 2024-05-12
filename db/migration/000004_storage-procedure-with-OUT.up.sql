CREATE PROCEDURE GetOrderDetailsForCustomer(
    IN customer_id INT,
    OUT total_amount FLOAT,
    OUT order_count INT
)
BEGIN
    SELECT
        SUM(CAST(od.quantity * p.price AS FLOAT)) AS total_amount,
        COUNT(o.id) AS order_count
    INTO
        total_amount,
        order_count
    FROM
        customer c
    LEFT JOIN
        orders o ON c.id = o.customer_id
    LEFT JOIN
        order_detail od ON o.id = od.order_id
    LEFT JOIN
        product p ON od.product_id = p.id
    WHERE
        c.id = customer_id;
END;
