\connect lms_db;

ALTER TABLE lms.teachers
  ADD PRIMARY KEY (id),
  ADD CHECK (academic_degree in ('master', 'doctor', 'profesional', 'bachelor'));

ALTER TABLE lms.students
  ADD PRIMARY KEY (id),
  ADD CHECK (birthdate BETWEEN '1960-01-01' AND '2007-01-01');

ALTER TABLE lms.institutions
  ADD PRIMARY KEY (id),
  ADD CHECK (phone LIKE '[0-9]{5, 20}');

ALTER TABLE lms.courses
  ADD PRIMARY KEY (id),
  ADD CONSTRAINT fk_teacher_id FOREIGN KEY (teacher_id) REFERENCES lms.teachers (id),
  ADD CONSTRAINT fk_institution_id FOREIGN KEY (institution_id) REFERENCES lms.institutions (id);

ALTER TABLE lms.studying_processes
  ADD CHECK (score BETWEEN 1 AND 5),
  ADD CHECK (status in ('in_progress', 'done')),
  ADD CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES lms.courses (id),
  ADD CONSTRAINT fk_student_id FOREIGN KEY (student_id) REFERENCES lms.students (id);

