CREATE TABLE categorie (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
);

CREATE TABLE customer (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(50),
    contact_name VARCHAR(30),
    address TEXT,   
    city VARCHAR(255),
    postal_code VARCHAR(16),
    country VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
);

CREATE TABLE employee (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    last_name VARCHAR(30),
    first_name VARCHAR(30),
    birth_date VARCHAR(30),
    photo VARCHAR(50),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
);

CREATE TABLE shipper (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    shipper_name VARCHAR(50),
    phone VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
);

CREATE TABLE supplier (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    supplier_name VARCHAR(50),
    contact_name VARCHAR(50),
    address TEXT,
    city VARCHAR(30),
    postal_code VARCHAR(16),
    country VARCHAR(30),
    phone VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
);

CREATE TABLE orders (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    customer_id INT(11),
    employee_id INT(11),
    shipper_id INT(11),
    order_date VARCHAR(30),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (customer_id) REFERENCES customer(id),
    FOREIGN KEY (employee_id) REFERENCES employee(id),
    FOREIGN KEY (shipper_id) REFERENCES shipper(id)
);

CREATE TABLE product (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    product_name VARCHAR(50),
    supplier_id INT(11),
    category_id INT(11),
    unit VARCHAR(30),
    price FLOAT(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (supplier_id) REFERENCES supplier(id),
    FOREIGN KEY (category_id) REFERENCES categorie(id)
);

CREATE TABLE order_detail (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    order_id INT(11),
    product_id INT(11),
    quantity INT(11),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);