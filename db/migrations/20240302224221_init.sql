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

CREATE TABLE offer (
	id INT GENERATED ALWAYS AS IDENTITY,
	name VARCHAR(255) NOT NULL,
	brand_id INT NOT NULL,
	quantity INT NOT NULL,
	available INT NOT NULL,
	price INT NOT NULL,
	store_id INT NOT NULL,
	picture VARCHAR(255) NOT NULL,
	expiration_date TIMESTAMP NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (store_id) REFERENCES store(id),
	FOREIGN KEY (brand_id) REFERENCES brand(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE offer;
DROP TABLE store;
DROP TABLE brand;
-- +goose StatementEnd
