-- 1913 · Query Student Enrollment Information
-- Write an SQL statement that satisfies the condition that, regardless of whether students have enrolments, students need to provide the following information based on the following two tables.

-- student_name,
-- phone,
-- hometown,
-- address
-- Table Definition 1: students (Student table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- student_name	varchar	student_name
-- phone	varchar	student_phone
-- Table Definition 2: enrollments (Academic Records Table)

-- column name	type	comments
-- id	int unsigned	primary key
-- student_id	int unsigned	student id
-- hometown	varchar	hometown
-- address	varchar	address
-- Example
-- Example 1:

-- Table content 1: students

-- id	student_name	phone
-- 1	Li Lei	13888888888
-- 2	Han Meimei	13999999999
-- 3	Amy	13788889999
-- Table content 2: enrollments

-- id	student_id	hometown	address
-- 1	1	Shi Jiazhuang	Hang Zhou
-- 2	2	Heng Shui	Tang Shan
-- 3	3	Cang Zhou	Shi Jiazhuang
-- Result：

-- student_name	phone	hometown	address
-- Li Lei	13888888888	Shi Jiazhuang	Hang Zhou
-- Han Meimei	13999999999	Heng Shui	Tang Shan
-- Amy	13788889999	Cang Zhou	Shi Jiazhuang
-- Example 2:

-- Table content 1: students

-- id	student_name	phone
-- 1	Li Lei	13888888888
-- 2	Han Meimei	13999999999
-- 3	Amy	13788889999
-- 4	Jason	13788789999
-- Table content 2: enrollments

-- id	student_id	hometown	address
-- 1	1	Shi Jiazhuang	Hang Zhou
-- 2	2	Heng Shui	Tang Shan
-- 3	3	Cang Zhou	Shi Jiazhuang
-- 4	1	Guang Zhou	Shi Hezi
-- Should return:

-- student_name	phone	hometown	address
-- Li Lei	13888888888	Shi Jiazhuang	Hang Zhou
-- Han Meimei	13999999999	Heng Shui	Tang Shan
-- Amy	13788889999	Cang Zhou	Shi Jiazhuang
-- Li Lei	13888888888	Guang Zhou	Shi Hezi
-- Jason	13788789999	null	null

SELECT
	s.student_name AS student_name,
	s.phone AS phone,
	e.hometown AS hometown,
	e.address AS address
FROM 
	students AS s
LEFT JOIN  
	enrollments AS e
ON s.id = e.student_id