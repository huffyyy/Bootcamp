CREATE TABLE oe.categories (
    category_id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_name VARCHAR(15),
    description TEXT,
    picture BYTEA
);

CREATE TABLE oe.suppliers (
    supplier_id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_name VARCHAR(40),
    contact_name VARCHAR(30),
    contact_title VARCHAR(30),
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    phone VARCHAR(24),
    fax VARCHAR(24),
    homepage TEXT
);

CREATE TABLE oe.products (
    product_id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_name VARCHAR(40),
    quantity_per_unit VARCHAR(20),
    unit_price REAL,
    units_in_stock SMALLINT,
    units_in_order SMALLINT,
    reorder_level SMALLINT,
    discontinued INTEGER,
    supplier_id SMALLINT,
    category_id SMALLINT,
    CONSTRAINT fk_products_supplier
        FOREIGN KEY (supplier_id)
        REFERENCES oe.suppliers (supplier_id),
    CONSTRAINT fk_products_category
        FOREIGN KEY (category_id)
        REFERENCES oe.categories (category_id)
);

CREATE TABLE oe.customers (
    customer_id CHAR(5) PRIMARY KEY,
    company_name VARCHAR(40),
    contact_name VARCHAR(30),
    contact_title VARCHAR(30),
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    phone VARCHAR(24),
    fax VARCHAR(24)
);

CREATE TABLE oe.employees (
    employee_id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    last_name VARCHAR(20),
    first_name VARCHAR(10),
    title VARCHAR(30),
    title_of_courtesy VARCHAR(25),
    birth_date DATE,
    hire_date DATE,
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    home_phone VARCHAR(24),
    extension VARCHAR(4),
    photo BYTEA,
    notes TEXT,
    reports_to SMALLINT,
    photo_path VARCHAR(255),
    CONSTRAINT fk_employees_manager
        FOREIGN KEY (reports_to)
        REFERENCES oe.employees (employee_id)
);

CREATE TABLE oe.shippers (
    shipper_id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_name VARCHAR(40),
    phone VARCHAR(24)
);

CREATE TABLE oe.orders (
    order_id SMALLINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_date DATE,
    required_date DATE,
    shipped_date DATE,
    freight REAL,
    ship_name VARCHAR(40),
    ship_address VARCHAR(60),
    ship_city VARCHAR(15),
    ship_region VARCHAR(15),
    ship_postal_code VARCHAR(10),
    ship_country VARCHAR(15),
    customer_id CHAR(5),
    employee_id SMALLINT,
    ship_via SMALLINT,
    CONSTRAINT fk_orders_customer
        FOREIGN KEY (customer_id)
        REFERENCES oe.customers (customer_id),
    CONSTRAINT fk_orders_employee
        FOREIGN KEY (employee_id)
        REFERENCES oe.employees (employee_id),
    CONSTRAINT fk_orders_shipper
        FOREIGN KEY (ship_via)
        REFERENCES oe.shippers (shipper_id)
);

CREATE TABLE oe.order_detail (
    order_id SMALLINT,
    product_id SMALLINT,
    unit_price REAL,
    quantity SMALLINT,
    discount REAL,
    CONSTRAINT pk_order_detail
        PRIMARY KEY (order_id, product_id),
    CONSTRAINT fk_od_order
        FOREIGN KEY (order_id)
        REFERENCES oe.orders (order_id),
    CONSTRAINT fk_od_product
        FOREIGN KEY (product_id)
        REFERENCES oe.products (product_id)
);