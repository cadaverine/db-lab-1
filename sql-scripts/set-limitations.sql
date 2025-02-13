\connect lms_db;

ALTER TABLE lms.teachers
  ADD PRIMARY KEY (id),
  ADD CHECK (academic_degree in ('master', 'doctor', 'profesional', 'bachelor'));

ALTER TABLE lms.students
  ADD PRIMARY KEY (id),
  ADD CHECK (birthdate BETWEEN '1960-01-01' AND '2007-01-01');

ALTER TABLE lms.institutions
  ADD PRIMARY KEY (id);

ALTER TABLE lms.courses
  ADD PRIMARY KEY (id),
  ADD CHECK (duration > 0),
  ADD CONSTRAINT fk_teacher_id FOREIGN KEY (teacher_id) REFERENCES lms.teachers (id),
  ADD CONSTRAINT fk_institution_id FOREIGN KEY (institution_id) REFERENCES lms.institutions (id);

ALTER TABLE lms.studying_processes
  ADD CHECK (score BETWEEN 0 AND 100),
  ADD CHECK (status in ('registered', 'in progress', 'done', 'dropped out')),
  ADD CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES lms.courses (id),
  ADD CONSTRAINT fk_student_id FOREIGN KEY (student_id) REFERENCES lms.students (id);

