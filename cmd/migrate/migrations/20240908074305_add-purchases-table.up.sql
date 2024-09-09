CREATE TABLE IF NOT EXISTS purchases (
    id SERIAL PRIMARY KEY, 
    inventory_id SERIAL, 
    customer uuid NOT NULL, 
    "date" timestamp DEFAULT CURRENT_DATE, 
    FOREIGN KEY (inventory_id) REFERENCES inventory(id),
    FOREIGN KEY (customer) REFERENCES customers(id)
);

