CREATE TABLE IF NOT EXISTS products (
  id SERIAL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL, 
  image UUID NOT NULL DEFAULT '32dfd1c8-7ff6-5909-d983-6d4adfb4255d',
  price INT NOT NULL,
  permissions INT DEFAULT 0,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id)
);
