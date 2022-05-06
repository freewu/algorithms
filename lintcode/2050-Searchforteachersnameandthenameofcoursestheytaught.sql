-- 2050 · Search for teacher‘s name and the name of courses they taught
-- # Description
-- Write a SQL statement to make a left join between the course table courses and the teacher table teachers,
-- then query the name of the teacher and the name of the course they taught, 
-- with the result columns named course name course_name and teacher name teacher_name respectively.

-- Table Definition 1: courses (course table)

-- column name	type	comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Table Definition 2 : teachers (teachers table)

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The question is a left join query, with teachers as the left table and courses as the right table.
-- If the query of both tables are empty, then nothing is returned.
-- Example
-- Sample I:

-- Table Contents 1 : Courses (Courses table)

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- Table Contents 2 : teachers (Teachers table)

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- course_name	teacher_name
-- Big Data	Eastern Heretic
-- Data Analysis	Eastern Heretic
-- Dynamic Programming	Eastern Heretic
-- NULL	Northern Beggar
-- System Design	Western Venom
-- Django	Western Venom
-- Artificial Intelligence	Western Venom
-- Java P6+	Western Venom
-- Senior Algorithm	Southern Emperor
-- Web	Southern Emperor
-- Web	Southern Emperor
-- NULL	Linghu Chong
-- Sample Example 2:
-- Table Contents 1 : Courses (Courses table)

-- id	name	student_count	created_at	teacher_id
-- 1	NULL	880	2018-01-13	4
-- 2	NULL	1350	2020-07-18	3
-- Table Contents 2 : teachers (Teachers table)

-- id	name	email	age	country
-- 1	NULL	eastern.heretic@gmail.com	20	UK
-- 2	NULL	northern.beggar@qq.com	21	CN
-- 3	NULL	western.venom@163.com	28	USA
-- After running your SQL statement, the table should return.

-- course_name	teacher_name
-- Because the course name and the teacher name in the sample are empty data, only the title is shown here, no data.

SELECT
	c.name AS course_name,
	t.name AS teacher_name
FROM
	teachers AS t
LEFT JOIN 
	courses  AS c
ON
	t.id = c.teacher_id