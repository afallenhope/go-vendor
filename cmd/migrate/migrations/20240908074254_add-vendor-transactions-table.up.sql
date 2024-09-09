CREATE TABLE IF NOT EXISTS vendor_transactions (
    id SERIAL PRIMARY KEY, 
    payer uuid NOT NULL, 
    receiver uuid NOT NULL, 
    merchant uuid NOT NULL, 
    vendor_id uuid NOT NULL, 
    vendor_location varchar(255) NOT NULL, 
    inventory_name varchar(255) NOT NULL, 
    "date" timestamp DEFAULT CURRENT_DATE, 
    FOREIGN KEY (payer) REFERENCES customers(id),
    FOREIGN KEY (receiver) REFERENCES customers(id),
    FOREIGN KEY (merchant) REFERENCES customers(id)
);

