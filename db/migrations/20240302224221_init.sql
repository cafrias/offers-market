-- +goose Up
-- +goose StatementBegin
CREATE TABLE brand (
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	name VARCHAR(255) NOT NULL
);

CREATE TABLE store (
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	name VARCHAR(255) NOT NULL,
	address VARCHAR(255),
	phone VARCHAR(255),
	website VARCHAR(255),
	email VARCHAR(255),
	location POINT
);

CREATE TABLE product (
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	name VARCHAR(255) NOT NULL,
	brand_id INT NOT NULL,
	picture VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	FOREIGN KEY (brand_id) REFERENCES brand(id)
);

CREATE TABLE offer (
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	available INT NOT NULL,
	quantity INT NOT NULL,
	price INT NOT NULL,
	store_id INT NOT NULL,
	product_id INT NOT NULL,
	expiration_date TIMESTAMP NOT NULL,
	description TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (store_id) REFERENCES store(id),
	FOREIGN KEY (product_id) REFERENCES product(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE offer;
DROP TABLE store;
DROP TABLE product;
DROP TABLE brand;
-- +goose StatementEnd
