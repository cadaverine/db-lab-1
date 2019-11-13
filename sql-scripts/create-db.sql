DROP SCHEMA IF EXISTS lms CASCADE;
CREATE SCHEMA lms;

CREATE TABLE lms.users (
	id bigserial primary key,
	first_name varchar(30),
	last_name varchar(30),
	phone varchar(30),
	email varchar(30)
);

CREATE TABLE lms.teachers (
	academic_degree varchar(30)
) INHERITS (lms.users);

CREATE TABLE lms.students (
	birthdate date
) INHERITS (lms.users);

CREATE TABLE lms.institutions (
	id bigserial primary key,
	name varchar(30),
	address varchar(30),
	phone varchar(30),
	email varchar(30)
);

CREATE TABLE lms.courses (
	id bigserial primary key,
	teacher_id bigserial,
	institution_id bigserial,
	name varchar(80),
	description text,
	duration varchar(30),
	registration_date date
);

CREATE TABLE lms.studying_processes (
	course_id bigserial,
	student_id bigserial,
	receipt_date date,
	deadline date,
	score int,
	status varchar(30)
);
