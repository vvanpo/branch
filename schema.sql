
-- E-mail is the canonical identifier for a contact.
CREATE TABLE contact (
	id uuid PRIMARY KEY,
	email text NOT NULL UNIQUE
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

CREATE TABLE contact_field_category (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	description text,
	parent uuid REFERENCES contact_field_category
);

CREATE TABLE contact_field (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	datatype text NOT NULL,
	description text,
	required boolean NOT NULL DEFAULT FALSE,
	category uuid REFERENCES contact_field_category,
	weight integer,
	UNIQUE (category, weight)
);

-- Relates fields to groups they exist in. Fields not present in this table are
-- common to all groups.
CREATE TABLE contact_field_groups (
	id uuid PRIMARY KEY,
	field uuid NOT NULL REFERENCES contact_field,
	contact_group uuid NOT NULL REFERENCES contact_group,
	UNIQUE (field, contact_group)
);

-- Relates field values to contacts.
CREATE TABLE contact_field_value (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	field uuid NOT NULL REFERENCES contact_field,
	value bytea,
	UNIQUE (contact, field)
);

CREATE TABLE role (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	description text
);

CREATE TABLE role_permissions (
	id uuid PRIMARY KEY,
	role uuid NOT NULL REFERENCES role,
	permission text NOT NULL,
	UNIQUE (role, permission)
);

-- Relates contacts to their roles.
CREATE TABLE contact_roles (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	role uuid NOT NULL REFERENCES role,
	UNIQUE (contact, role)
);
