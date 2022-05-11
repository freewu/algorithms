-- 2674 · Creating a simple SQL View (I)
-- # Description
-- Now, your supervisor has asked you to write a simple view named v_courses_teachers. 
-- You need to provide all the information in the courses table and show the names and emails of the teachers in the associated table teachers.

-- Table Definition 1: teachers (Teachers table)

-- column name	type	comment
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

-- Example
-- Enter data:

-- courses table.
-- id	name	student_count	created_at	teacher_id
-- 1	'Advanced Algorithms'	880	'2020-6-1 09:03:12'	4
-- 2	'System Design'	1350	'2020-7-18 10:03:12'	8
-- 3	'Django'	780	'2020-2-29 12:03:12'	2
-- 4	'Web'	340	'2020-4-22 13:03:12'	4
-- 5	'Big Data'	700	'2020-9-11 16:03:12'	7
-- 6	'Artificial Intelligence'	1660	'2018-5-13 18:03:12'	3
-- 7	'Java P6+'	780	'2019-1-19 13:03:12'	3
-- 8	'Data Analysis'	500	'2019-7-12 13:03:12'	6
-- 10	'Object Oriented Design'	300	'2020-8-8 13:03:12'	4
-- 12	'Dynamic Programming'	2000	'2018-8-18 20:03:12'	1

-- teachers table.
-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'

-- Returned results:
-- id	name	student_count	created_at	teacher_id	teacher_name	teacher_email
-- 1	'Advanced Algorithms'	880	'2020-06-01'	4	'Southern Emperor'	'southern.emperor@qq.com'
-- 2	'System Design'	1350	'2020-07-18'	8	None	None
-- 3	'Django'	780	'2020-02-29'	2	'Northern Beggar'	'northern.beggar@qq.com'
-- 4	'Web'	340	'2020-04-22'	4	'Southern Emperor'	'southern.emperor@qq.com'
-- 5	'Big Data'	700	'2020-09-11'	7	None	None
-- 6	'Artificial Intelligence'	1660	'2018-05-13'	3	'Western Venom'	'western.venom@163.com'
-- 7	'Java P6+'	780	'2019-01-19'	3	'Western Venom'	'western.venom@163.com'
-- 8	'Data Analysis'	500	'2019-07-12'	6	None	None
-- 10	'Object Oriented Design'	300	'2020-08-08'	4	'Southern Emperor'	'southern.emperor@qq.com'
-- 12	'Dynamic Programming'	2000	'2018-08-18'	1	'Eastern heretic'	'eastern.heretic@gmail.com'

CREATE VIEW 
	`v_courses_teachers`
AS 
	SELECT 
		c.id AS id,
		c.name AS name,
		c.student_count AS student_count,
		c.created_at AS created_at,
		c.teacher_id AS teacher_id,
		t.name AS teacher_name,
		t.email AS teacher_email
	FROM 
		courses  AS c
	LEFT JOIN
		teachers AS t
	ON
		c.teacher_id = t.id;