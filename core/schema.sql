
-- The application shouldn't be the schema owner, and shouldn't be able to alter
-- the schema.
REVOKE CREATE ON SCHEMA public FROM PUBLIC;

-- A contact can represent an individual or organization.
CREATE TABLE contact (
	id uuid PRIMARY KEY
);

-- E-mail is the unique identifier for a contact. Every verified e-mail (i.e.
-- known to identify a particular contact) is associated with a contact, and
-- vice versa.
CREATE TABLE contact_email (
	id uuid PRIMARY KEY,
	email text NOT NULL UNIQUE,
	contact uuid NOT NULL REFERENCES contact,
	primary_email boolean NOT NULL DEFAULT TRUE,
	UNIQUE (contact, primary_email)
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

CREATE TYPE contact_field_type AS ENUM (
	'label', 'text', 'number', 'flag', 'checklist', 'selection', 'date', 'money', 'image', 'upload'
);

CREATE TABLE contact_field (
	id uuid PRIMARY KEY,
	name text NOT NULL,
	description text,
	category uuid NOT NULL REFERENCES contact_field_category,
	field_type contact_field_type NOT NULL,
	typedata bytea,
	user_read_access bool NOT NULL DEFAULT TRUE,
	user_write_access bool NOT NULL DEFAULT TRUE
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

-- Defines which users can access their own fields via a relation to the groups
-- they belong to. Fields not present in this table are fully accessible to
-- their owners.
CREATE TABLE contact_field_user_access (
	id uuid PRIMARY KEY,
	contact_field uuid NOT NULL REFERENCES contact_field,
	contact_group uuid NOT NULL REFERENCES contact_group,
	readonly boolean NOT NULL DEFAULT TRUE,
	UNIQUE (contact_field, contact_group)
);

-- Defines which groups can access which fields of any contact. Fields not
-- present in this table are public.
CREATE TABLE contact_field_group_visibility (
	id uuid PRIMARY KEY,
	contact_field uuid NOT NULL REFERENCES contact_field,
	contact_group uuid NOT NULL REFERENCES contact_group,
	admin boolean NOT NULL DEFAULT FALSE,
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
