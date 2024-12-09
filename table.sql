CREATE TABLE brands (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE cars (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  license_plate VARCHAR(50) NOT NULL,
  color VARCHAR(50) NOT NULL,
  brand_id INT NOT NULL, -- Relasi ke Brand
  published_date DATE,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT fk_brand_id FOREIGN KEY (brand_id) REFERENCES brands (id)
);