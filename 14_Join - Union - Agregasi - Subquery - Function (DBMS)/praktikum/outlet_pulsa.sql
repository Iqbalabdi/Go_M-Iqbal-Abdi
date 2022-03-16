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
	names VARCHAR(255),
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
	FOREIGN KEY (transaction_id) REFERENCES transactions(id),
	FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT into operators (id, name) VALUES (1, 'telkomsel');
INSERT into operators (id, name) VALUES (2, 'indosat');
INSERT into operators (id, name) VALUES (3, 'axiata');
INSERT into operators (id, name) VALUES (4, 'byu');
INSERT into operators (id, name) VALUES (5, 'tri');

INSERT into operators (id, product_type_id) VALUES (1, 'tri');

INSERT into `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
	(1, 'gopay', NULL, NULL, NULL),
	(2, 'e-banking', NULL, NULL, NULL),
	(3, 'qris', NULL, NULL, NULL);

INSERT into `users` (`id`, `names`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
	(1, 'dedi', NULL, '2000-03-16', '1', NULL, NULL),
	(2, 'corbuzer', NULL, '2000-05-31', '1', NULL, NULL),
	(3, 'emma', NULL, '2000-10-09', '2', NULL, NULL),
	(4, 'watson', NULL, '2000-08-08', '2', NULL, NULL),
	(5, 'ilham', NULL, '2000-12-12', '1', NULL, NULL);

INSERT into `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
	(1, 1, 3, NULL, 'product dummy', NULL, NULL, NULL),
	(2, 1, 3, NULL, 'pulsa 20k', NULL, NULL, NULL),
	(3, 2, 1, NULL, 'kuota 50gb', NULL, NULL, NULL),
	(4, 2, 1, NULL, 'kuota 20gb', NULL, NULL, NULL),
	(5, 2, 1, NULL, 'kuota 100gb', NULL, NULL, NULL),
	(6, 3, 4, NULL, 'instagram 5gb', NULL, NULL, NULL),
	(7, 3, 4, NULL, 'whatsapp 2gb', NULL, NULL, NULL),
	(8, 3, 4, NULL, 'line 10gb', NULL, NULL, NULL);

INSERT into `transaction_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
	(1, 1, NULL, 1, NULL, NULL, NULL),
	(1, 1, NULL, 2, NULL, NULL, NULL),
	(1, 1, NULL, 3, NULL, NULL, NULL),
	(2, 2, NULL, 1, NULL, NULL, NULL),
	(2, 2, NULL, 2, NULL, NULL, NULL),
	(2, 2, NULL, 3, NULL, NULL, NULL),
	(3, 3, NULL, 1, NULL, NULL, NULL),
	(3, 3, NULL, 2, NULL, NULL, NULL),
	(3, 3, NULL, 3, NULL, NULL, NULL);

Insert into `transactions` (`id`, `user_id`, `payment_methods_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
	(1, 1, 1, NULL, 2, 12000.00, NULL, NULL),
	(2, 1, 1, NULL, 3, 13000.00, NULL, NULL),
	(3, 1, 1, NULL, 1, 10000.00, NULL, NULL),
	(4, 2, 2, NULL, 2, 15000.00, NULL, NULL),
	(5, 2, 2, NULL, 3, 30000.00, NULL, NULL),
	(6, 2, 2, NULL, 2, 20000.00, NULL, NULL),
	(7, 3, 3, NULL, 1, 10000.00, NULL, NULL),
	(8, 3, 3, NULL, 2, 20000.00, NULL, NULL),
	(9, 4, 3, NULL, 3, 30000.00, NULL, NULL),
	(10, 4, 3, NULL, 1, 10000.00, NULL, NULL),
	(11, 5, 2, NULL, 2, 20000.00, NULL, NULL),
	(12, 5, 2, NULL, 3, 30000.00, NULL, NULL);

ALTER TABLE users ADD names varchar(255) NOT NULL AFTER id;
