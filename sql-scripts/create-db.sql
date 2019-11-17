DROP DATABASE IF EXISTS lms_db;
CREATE DATABASE lms_db;

\connect lms_db;

DROP SCHEMA IF EXISTS lms CASCADE;
CREATE SCHEMA lms;

CREATE TABLE lms.teachers (
	id bigserial,
	first_name varchar(30),
	last_name varchar(30),
	phone varchar(30),
	email varchar(30),
	academic_degree varchar(30)
);

CREATE TABLE lms.students (
	id bigserial,
	first_name varchar(30),
	last_name varchar(30),
	phone varchar(30),
	email varchar(30),
	birthdate date
);

CREATE TABLE lms.institutions (
	id bigserial,
	name varchar(30),
	address varchar(30),
	phone varchar(30),
	email varchar(30)
);

CREATE TABLE lms.courses (
	id bigserial,
	teacher_id bigserial,
	institution_id bigserial,
	name varchar(80),
	description text,
	duration numeric,
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
