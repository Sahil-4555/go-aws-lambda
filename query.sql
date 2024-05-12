-- name: GetCustomer :one
SELECT * FROM customer
WHERE id = ? LIMIT 1;

-- name: ListCustomer :many
SELECT * FROM customer
ORDER BY customer_name;

-- name: CreateCustomer :execresult
INSERT INTO customer (
  customer_name, contact_name, address, city, postal_code, country
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: UpdateCustomer :execresult
UPDATE customer
SET
  customer_name = ?,
  contact_name = ?,
  address = ?,
  city = ?,
  postal_code = ?,
  country = ?
WHERE
  id = ?;

-- name: CreateCategorie :execresult
INSERT INTO categorie (
  category_name, description
) VALUES (
  ?, ?
);

-- name: CreateEmployee :execresult
INSERT INTO employee (
  last_name, first_name, birth_date, photo, notes
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: CreateShipper :execresult
INSERT INTO shipper (
  shipper_name, phone
) VALUES (
  ?, ?
);

-- name: CreateSupplier :execresult
INSERT INTO supplier (
  supplier_name, contact_name, address, city, postal_code, country, phone
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateOrder :execresult
INSERT INTO orders (
  id, customer_id, employee_id, shipper_id, order_date
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: CreateProduct :execresult
INSERT INTO product (
  product_name, supplier_id, category_id, unit, price
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: CreateOrderDetail :execresult
INSERT INTO order_detail (
  order_id, product_id, quantity
) VALUES (
  ?, ?, ?
);

-- name: DeleteCustomer :exec
DELETE FROM customer
WHERE id = ?;

-- name: OrdersAssocaitedWithCustomer :many
SELECT
    c.id AS customer_id,
    c.customer_name,
    c.contact_name,
    c.address,
    c.city,
    c.postal_code,
    c.country,
    o.id AS order_id,
    o.order_date,
    od.id AS order_detail_id,
    p.product_name,
    od.quantity,
    p.price,
    CAST((od.quantity * p.price) AS FLOAT) AS total_amount
FROM
    customer c
LEFT JOIN
    orders o ON c.id = o.customer_id
LEFT JOIN
    order_detail od ON o.id = od.order_id
LEFT JOIN
    product p ON od.product_id = p.id
WHERE
    c.id = ?;

-- name: SoftDeleteCustomer :exec
UPDATE customer
SET
  deleted_at = ?
WHERE
  id = ?;






