\connect lms_db;

\copy lms.teachers(first_name,last_name,phone,email,academic_degree) FROM '../test-data/teachers.csv' CSV HEADER;

\copy lms.students(first_name,last_name,phone,email,birthdate) FROM '../test-data/students.csv' CSV HEADER;

\copy lms.institutions(name,address,phone,email) FROM '../test-data/institutions.csv' CSV HEADER;

\copy lms.courses(teacher_id,institution_id,name,description,duration,registration_date) FROM '../test-data/courses.csv' CSV HEADER;

\copy lms.studying_processes(course_id,student_id,receipt_date,deadline,score,status) FROM '../test-data/stydying-processes.csv' CSV HEADER;
