CREATE TABLE IF NOT EXISTS inventory (
  id SERIAL,
  name VARCHAR(200) NOT NULL,
  description TEXT,
  image VARCHAR(255),
  marketplace_id INT,
  product_id INT,
  inventory_name VARCHAR(63),

  PRIMARY KEY (id)
);
