CREATE SCHEMA field;

--
CREATE TABLE field.category (
	id uuid PRIMARY KEY,
	name text NOT NULL,
	description text,
	parent uuid REFERENCES field.category
);

--
CREATE TABLE field.field (
	id uuid PRIMARY KEY,
	name text NOT NULL,
	description text,
	category uuid NOT NULL REFERENCES field.category
);
