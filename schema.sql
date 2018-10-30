
CREATE TABLE contact (
	id serial PRIMARY KEY,
	email text NOT NULL UNIQUE,
	name text NOT NULL,
	registered timestamp(0) with time zone NOT NULL DEFAULT now()
);

CREATE TABLE contact_group (
	id serial PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE contact_group_map (
	id serial PRIMARY KEY,
	contact integer NOT NULL REFERENCES contact,
	contact_group integer NOT NULL REFERENCES contact_group,
	UNIQUE (contact, contact_group)
);

CREATE TABLE role (
	id serial PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE permission (
	id serial PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE role_permission_map (
	id serial PRIMARY KEY,
	role integer NOT NULL REFERENCES role,
	permission integer NOT NULL REFERENCES permission,
	UNIQUE (role, permission)
);

CREATE TABLE contact_role_map (
	id serial PRIMARY KEY,
	contact integer NOT NULL REFERENCES contact,
	role integer NOT NULL REFERENCES role,
	UNIQUE (contact, role)
);

CREATE TABLE contact_field (
	id serial PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE contact_field_value (
	id serial PRIMARY KEY,
	contact integer NOT NULL REFERENCES contact,
	field integer NOT NULL REFERENCES contact_field,
	data bytea,
	UNIQUE (contact, field)
);
