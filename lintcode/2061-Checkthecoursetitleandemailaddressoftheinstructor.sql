-- 2061 · Check the course title and email address of the instructor
-- # Description
-- Please write an SQL statement to query the course id, 
-- course name and teachers' email address by combining the course table id , 
-- courses and the teacher table teachers, with the column names displayed as course_name and teacher_email respectively.

-- Table Definition 1: courses

-- column name	type	comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	time of course creation
-- teacher_id	int	lecture id
-- Table Definition 2 : teachers

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Tutor's email
-- age	int	Tutor's age
-- country	varchar	Tutor's nationality

-- This question is an inner join of two tables
-- If both tables are queried as empty, nothing is returned
-- Example
-- Example 1:

-- Table Contents 1 : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	5
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- id	course_name	teacher_email
-- 1	Advanced Algorithms	southern.emperor@qq.com
-- 2	System Design	western.venom@163.com
-- 3	Django	NULL
-- 4	Web	southern.emperor@qq.com
-- 5	Big Data	eastern.heretic@gmail.com
-- 6	Artificial Intelligence	western.venom@163.com
-- Example 2:
-- Table Contents 1 : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	3
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 11	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 12	Northern Beggar	northern.beggar@qq.com	21	CN
-- 13	Western Venom	western.venom@163.com	28	USA
-- 14	Southern Emperor	southern.emperor@qq.com	21	JP
-- 15	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- id	course_name	teacher_email
-- Because the teachers in the Teacher table in Example 2 do not have a corresponding course, 
-- i.e. the teacher numbers in the two tables do not have the same item, the query does not find anything that matches the condition and returns null.
SELECT
	c.id AS id,
	c.name AS	course_name,
	t.email AS teacher_email
FROM
	teachers AS t,
	courses AS c 
WHERE
	t.id = c.teacher_id