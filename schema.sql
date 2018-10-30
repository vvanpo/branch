
CREATE TABLE permission (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE role (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE role_permissions (
	id uuid PRIMARY KEY,
	role uuid NOT NULL REFERENCES role,
	permission uuid NOT NULL REFERENCES permission,
	UNIQUE (role, permission)
);

-- E-mail is the sole identifier for a contact.
CREATE TABLE contact (
	id uuid PRIMARY KEY,
	email text NOT NULL UNIQUE
);

-- Relates contacts to their roles.
CREATE TABLE contact_roles (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	role uuid NOT NULL REFERENCES role,
	UNIQUE (contact, role)
);

CREATE TABLE contact_group (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE
);

-- Relates contacts belonging to groups.
CREATE TABLE contact_groups (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	contact_group uuid NOT NULL REFERENCES contact_group,
	UNIQUE (contact, contact_group)
);

CREATE TABLE contact_field (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	datatype text NOT NULL
);

-- Relates fields and their values to contacts.
CREATE TABLE contact_field_value (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	field uuid NOT NULL REFERENCES contact_field,
	value bytea,
	UNIQUE (contact, field)
);

-- Relates groups with their required fields.
CREATE TABLE contact_group_field_required (
	id uuid PRIMARY KEY,
	contact_group uuid NOT NULL REFERENCES contact_group,
	field uuid NOT NULL REFERENCES contact_field,
	UNIQUE (contact_group, field)
);
