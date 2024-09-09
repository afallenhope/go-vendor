CREATE TABLE IF NOT EXISTS customers (
  id UUID DEFAULT gen_random_uuid() NOT NULL,
  firstname VARCHAR(31) NOT NULL,
  lastname VARCHAR(31) NOT NULL DEFAULT 'resident',
  flags INT NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
