-- 2692 · Count the total number of students taught by each faculty member
-- # Description
-- We have provided a view v_courses with the teacher id and the total number of students taught by that teacher. 
-- Now we want to get more information about the teachers, create a new view v_teachers based on the view v_courses to show more information.

-- Table Definition 1: teachers (Teachers table)
-- column name	type	comments
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality

-- Table Definition 2: courses (Course List)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total_students
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor id

-- View Definition: v_courses(view)
-- column_name	type	comment
-- teacher_id	int unsigned	instructor id
-- student_count	int	total_students

-- Result view definition: v_teachers(view)
-- column_name	type	comment
-- teacher_id	int unsigned	instructor id
-- name	varchar	instructor's name
-- email	varchar	instructor's email
-- age	int	instructor's age
-- country	varchar	Tutor's nationality
-- student_count	int	total number of students

-- Example
-- Enter data:

-- courses table.

-- id	name	student_count	created_at	teacher_id
-- 1	'Advanced Algorithms'	880	'2020-6-1 09:03:12'	4
-- 2	'System Design'	1350	'2020-7-18 10:03:12'	2
-- 3	'Django'	780	'2020-2-29 12:03:12'	2
-- 4	'Web'	340	'2020-4-22 13:03:12'	4
-- 5	'Big Data'	700	'2020-9-11 16:03:12'	3
-- 6	'Artificial Intelligence'	1660	'2018-5-13 18:03:12'	3
-- 7	'Java P6+'	780	'2019-1-19 13:03:12'	1
-- 8	'Data Analysis'	500	'2019-7-12 13:03:12'	1
-- 10	'Object Oriented Design'	300	'2020-8-8 13:03:12'	4
-- 12	'Dynamic Programming'	2000	'2018-8-18 20:03:12'	1
-- teachers table.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	'linghu@163.com'	18	'CN'
-- Returned results:

-- teacher_id	name	email	age	country	student_count
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'	3280
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'	3770
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'	2360
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'	1520

-- Write your SQL here --
-- v_courses
CREATE OR REPLACE VIEW
	`v_courses`
AS 
	SELECT 
		teacher_id,
		SUM(student_count) AS student_count
	FROM
		courses
	GROUP BY
		teacher_id;

-- v_teachers
CREATE OR REPLACE VIEW
	`v_teachers`
AS 
	SELECT 
		vc.teacher_id AS teacher_id,
		t.name AS name,
		t.email AS email,
		t.age AS age,
		t.country AS country,
		vc.student_count
	FROM
		teachers AS t
	RIGHT JOIN 
		v_courses AS vc 
	ON
		vc.teacher_id = t.id;
