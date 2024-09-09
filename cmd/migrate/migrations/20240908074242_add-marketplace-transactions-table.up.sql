CREATE TABLE IF NOT EXISTS marketplace_transactions (
    id BIGSERIAL PRIMARY KEY, 
    type varchar(50), 
    payment_gross varchar(12), 
    payment_fee varchar(12), 
    payer uuid NOT NULL, 
    receiver uuid NOT NULL, 
    merchant uuid NOT NULL, 
    marketplace_id SERIAL NOT NULL, 
    marketplace_name varchar(255) NOT NULL, 
    inventory_name varchar(255) NOT NULL, 
    "date" timestamp DEFAULT CURRENT_DATE, 
    FOREIGN KEY (payer) REFERENCES customers(id),
    FOREIGN KEY (receiver) REFERENCES customers(id),
    FOREIGN KEY (merchant) REFERENCES customers(id)
);

