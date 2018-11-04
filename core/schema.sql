
-- The application shouldn't be the schema owner, and shouldn't be able to alter
-- the schema.
REVOKE CREATE ON SCHEMA public FROM PUBLIC;

-- A contact can represent an individual or organization.
CREATE TABLE contact (
	id uuid PRIMARY KEY,
	email text NOT NULL UNIQUE
);

CREATE TABLE contact_group (
	id uuid PRIMARY KEY,
	name text NOT NULL UNIQUE,
	description text
);

-- Relates groups to sets of permissions.
-- The capabilities of permissions are determined by the application.
CREATE TABLE contact_group_permissions (
	id uuid PRIMARY KEY,
	contact_group uuid NOT NULL REFERENCES contact_group,
	permission text NOT NULL,
	UNIQUE (contact_group, permission)
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
	description text,
	category uuid REFERENCES contact_field_category,
	datatype text NOT NULL
);

-- Relates fields to groups they exist in. Fields not present in this table are
-- common to all contacts, regardless of group.
CREATE TABLE contact_field_groups (
	id uuid PRIMARY KEY,
	contact_field uuid NOT NULL REFERENCES contact_field,
	contact_group uuid NOT NULL REFERENCES contact_group,
	required boolean NOT NULL DEFAULT FALSE,
	UNIQUE (contact_field, contact_group)
);

-- Relates fields to the groups that can access them. Fields not in this table
-- are public.
CREATE TABLE contact_field_group_access (
	id uuid PRIMARY KEY,
	contact_field uuid NOT NULL REFERENCES contact_field,
	contact_group uuid NOT NULL REFERENCES contact_group,
	readonly boolean NOT NULL DEFAULT TRUE,
	UNIQUE (contact_field, contact_group)
);

-- Relates field values to contacts.
CREATE TABLE contact_field_value (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL REFERENCES contact,
	contact_field uuid NOT NULL REFERENCES contact_field,
	value bytea,
	UNIQUE (contact, contact_field)
);

-- Users are contacts that are able to access the application.
CREATE TABLE users (
	id uuid PRIMARY KEY,
	contact uuid NOT NULL UNIQUE REFERENCES contact
);

CREATE TABLE user_session (

);
