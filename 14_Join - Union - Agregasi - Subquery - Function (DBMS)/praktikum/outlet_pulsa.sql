CREATE TABLE product_types(
	id int(11) NOT NULL PRIMARY KEY,
	name varchar(255),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE operators(
	id int(11) NOT NULL PRIMARY KEY,
	name VARCHAR(255),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE products(
	id int(11) NOT NULL PRIMARY KEY,
	product_type_id INT (11),
	operator_id int(11),
	code VARCHAR(50),
	name VARCHAR(50),
	status int,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	FOREIGN KEY (product_type_id) REFERENCES product_types(id),
	FOREIGN KEY (operator_id) REFERENCES operators(id)
);

CREATE TABLE product_descriptions(
	id int(11) NOT NULL PRIMARY KEY,
	description TEXT,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE users(
	id int(11) NOT NULL PRIMARY KEY,
	status int,
	dob DATE,
	gender CHAR(1),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE payment_methods(
	id int(11) NOT NULL PRIMARY KEY,
	name VARCHAR(255),
	status int,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE transactions(
	id int(11) NOT NULL PRIMARY KEY,
	user_id INT(11),
	payment_methods_id int(11),
	status VARCHAR(10),
	total_qty int(11),
	total_price NUMERIC(25,2),
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (payment_methods_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details(
	transaction_id INT (11),
	product_id INT(11),
	status VARCHAR(10),
	qty int(11),
	price NUMERIC(25,2),
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	PRIMARY KEY (transaction_id, product_id)
);

INSERT into operators 