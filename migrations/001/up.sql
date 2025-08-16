CREATE TABLE author (
	id SERIAL PRIMARY KEY,
	name VARCHAR(64) NOT NULL,
	description BPCHAR NOT NULL,
	photo_url VARCHAR(255)
);

CREATE TABLE category (
	id SERIAL PRIMARY KEY,
	name VARCHAR(32) NOT NULL
);

CREATE TABLE tag (
	id SERIAL PRIMARY KEY,
	name VARCHAR(32) NOT NULL
);

CREATE TABLE language (
	code VARCHAR(8) PRIMARY KEY,
	name VARCHAR(32) NOT NULL
);

CREATE TABLE book (
	id CHAR(16) PRIMARY KEY,
	file_path VARCHAR(128) NOT NULL UNIQUE,
	title VARCHAR(255) NOT NULL,
	description BPCHAR,
	format VARCHAR(8) NOT NULL,
	category_id INTEGER NOT NULL REFERENCES category(id),
	language_code VARCHAR(8) NOT NULL REFERENCES language(code),
	published TIMESTAMP DEFAULT NOW()
);

CREATE TABLE tag_to_book (
	tag_id INTEGER,
	book_id VARCHAR(255),
	PRIMARY KEY (tag_id, book_id)
);

CREATE TABLE author_to_book (
	author_id INTEGER,
	book_id VARCHAR(255),
	PRIMARY KEY (author_id, book_id)
);
