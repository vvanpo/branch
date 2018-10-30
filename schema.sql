
CREATE TABLE Contact (
	id serial PRIMARY KEY,
	email text NOT NULL UNIQUE,
	name text NOT NULL,
	registered timestamp(0) with time zone NOT NULL DEFAULT now()
);

CREATE TABLE ContactGroup (
	id serial PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE ContactGroupMap (
	id serial PRIMARY KEY,
	contact integer REFERENCES Contact,
	contactGroup integer REFERENCES ContactGroup
);

CREATE TABLE Role (
	id serial PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE Permission (
	id serial PRIMARY KEY,
	name text NOT NULL UNIQUE
);

CREATE TABLE RolePermissionMap (
	id serial PRIMARY KEY,
	permission integer NOT NULL REFERENCES Permission
);

CREATE TABLE ContactRoleMap (
	id serial PRIMARY KEY,
	contact integer NOT NULL REFERENCES Contact,
	role integer NOT NULL REFERENCES Role
);
