CREATE TABLE author (
	id CHAR(8) PRIMARY KEY DEFAULT substr(md5(random()::text), 1, 8),
	name VARCHAR(64) NOT NULL,
	description BPCHAR NOT NULL,
	photo_url VARCHAR(255)
);

CREATE TABLE book (
	id CHAR(8) PRIMARY KEY DEFAULT substr(md5(random()::text), 1, 8),
	filename VARCHAR(144) NOT NULL UNIQUE,
	title VARCHAR(128) NOT NULL,
	description BPCHAR,
	format VARCHAR(8) NOT NULL,
	category VARCHAR(64) NOT NULL,
	language_code VARCHAR(8) NOT NULL,
	published TIMESTAMP DEFAULT NOW()
);

CREATE TABLE request (
	id CHAR(8) PRIMARY KEY DEFAULT substr(md5(random()::text), 1, 8),
	filename VARCHAR(144) NOT NULL UNIQUE,
	title VARCHAR(128) NOT NULL,
	author_name VARCHAR(64) NOT NULL,
	description BPCHAR,
	category VARCHAR(64) NOT NULL,
	language_code VARCHAR(8) NOT NULL,
	added TIMESTAMP DEFAULT NOW()
);

CREATE TABLE tag (
	name VARCHAR(32) PRIMARY KEY
);

CREATE TABLE tag_to_request (
	tag_name VARCHAR(32) REFERENCES tag(name) ON DELETE CASCADE,
	request_id CHAR(8) REFERENCES request(id) ON DELETE CASCADE,
	PRIMARY KEY (tag_name, request_id)
);

CREATE TABLE tag_to_book (
	tag_name VARCHAR(32) REFERENCES tag(name) ON DELETE CASCADE,
	book_id CHAR(8) REFERENCES book(id) ON DELETE CASCADE,
	PRIMARY KEY (tag_name, book_id)
);

CREATE TABLE author_to_book (
	author_id CHAR(8) REFERENCES author(id) ON DELETE CASCADE,
	book_id CHAR(8) REFERENCES book(id) ON DELETE CASCADE,
	PRIMARY KEY (author_id, book_id)
);
