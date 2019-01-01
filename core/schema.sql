
-- The application shouldn't be the schema owner, and shouldn't be able to alter
-- the schema.
REVOKE CREATE ON SCHEMA public FROM PUBLIC;

--
CREATE TABLE field_category (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	description text,
	parent uuid REFERENCES field_category
);

CREATE TABLE field (
	id uuid PRIMARY KEY,
	name text NOT NULL,
	description text,
	category uuid NOT NULL REFERENCES field_category
);

-- A contact can represent an individual or organization.
CREATE TABLE contact (
	id uuid PRIMARY KEY
);

-- E-mail is the unique identifier for a contact. Every verified e-mail is
-- associated with a contact, and vice versa.
CREATE TABLE emailaddress (
	id uuid PRIMARY KEY,
	address text NOT NULL UNIQUE,
	contact uuid NOT NULL REFERENCES contact
);

CREATE TABLE contact_group (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	description text
);

-- Relates contacts and the groups they belong to.
CREATE TABLE contact_groups (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	contact_group uuid NOT NULL REFERENCES contact_group,
	UNIQUE (contact, contact_group)
);

-- Relates field values to contacts.
CREATE TABLE contact_field (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	field uuid NOT NULL REFERENCES field,
	UNIQUE (contact, field)
);

