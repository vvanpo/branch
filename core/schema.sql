
-- The application shouldn't be the schema owner, and shouldn't be able to alter
-- the schema.
REVOKE CREATE ON SCHEMA public FROM PUBLIC;

CREATE SCHEMA contact;

-- A contact can represent an individual or organization.
CREATE TABLE contact.contact (
	id uuid PRIMARY KEY
);

CREATE TABLE contact.group (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	description text
);

-- Relates contacts and the groups they belong to.
CREATE TABLE contact.groups (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact.contact,
	contact_group uuid NOT NULL REFERENCES contact.group,
	UNIQUE (contact.contact, contact.group)
);
