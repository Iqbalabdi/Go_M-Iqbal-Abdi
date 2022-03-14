CREATE TABLE user(
	id VARCHAR(50),
	nama VARCHAR(100),
	alamat VARCHAR(255),
	tanggal_lahir VARCHAR(30),
	status_user VARCHAR(20),
	gender VARCHAR(20),
	created_at VARCHAR(30),
	updated_at VARCHAR(30)
);

CREATE TABLE item (
	id VARCHAR(50),
	product VARCHAR(50),
	product_type VARCHAR(250),
	product_description VARCHAR(255),
	operator VARCHAR(20),
	payment_methods VARCHAR(20)
);

CREATE TABLE transaction(
	no_nota VARCHAR(50),
	id_product VARCHAR(50),
	id_user VARCHAR(50),
	harga VARCHAR(50),
	nominal VARCHAR(50)
);

CREATE TABLE kurir(
	id VARCHAR(20),
	name VARCHAR(50),
	created_at VARCHAR(30),
	updated_at VARCHAR(30)
);

ALTER TABLE kurir ADD COLUMN ongkos_dasar VARCHAR(30);
ALTER TABLE kurir RENAME TO shipping;
DROP TABLE shipping;

ALTER TABLE transaction ADD FOREIGN KEY (id_user) REFERENCES user(id);
