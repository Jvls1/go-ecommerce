\c product_service;

CREATE TABLE products (
  id UUID DEFAULT gen_random_uuid() NOT NULL,
  name VARCHAR NOT NULL,
  description TEXT,
  image_url VARCHAR,
  price NUMERIC(10, 2) NOT NULL,
  quantity INT NOT NULL,
  department_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

INSERT INTO products (id, name, description, image_url, price, quantity, department_id, created_at, updated_at, deleted_at)
VALUES ('6ba7b814-9dad-11d1-80b4-00c04fd430c8', 'Produto de Exemplo', 'Esta é uma descrição de exemplo', 'https://example.com/image.jpg', 19.99, 100, '880da008-32e4-4083-8b8a-43f9f3c9bbf2', current_timestamp, current_timestamp, null);
