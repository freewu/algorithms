-- 2062 · Query the id and name of all courses taught by the specified teacher
-- # Description
-- Write an SQL statement to inner join the course table courses 
-- and the teacher table teachers to query the course names 
-- and course ids of all the courses taught by the teacher "Eastern Heretic" 
-- and display the result columns with the course number id, course name and teacher name respectively. 
-- The result columns are displayed as id, course_name and teacher_name.

-- Table Definition 1: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course_name
-- student_count	int	total_students
-- created_at	date	time of course creation
-- teacher_id	int	lecturer_id
-- Table Definition 2 : teachers

-- column_name	type	comment
-- id	int	primary key
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
-- 3	Web	340	2020-4-22	4
-- 4	Big Data	700	2020-9-11	1
-- 5	Artificial Intelligence	1660	2018-5-13	3
-- 6	Java P6+	780	2019-1-19	3
-- 7	Data Analysis	500	2019-7-12	1
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- id	course_name	teacher_name
-- 4	Big Data	Eastern Heretic
-- 7	Data Analysis	Eastern Heretic
-- Example 2:
-- Table Contents 1 : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	1
-- 3	Django	780	2020-2-29	3
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 1	NULL	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- id	course_name	teacher_name
-- Because there is no "Eastern Heretic" teacher in the Teachers table in Example 2, no matches were found and the query returns null.

SELECT
	c.id AS id,
	c.name AS	course_name,
	t.name AS teacher_name
FROM
	teachers AS t,
	courses AS c 
WHERE
	t.id = c.teacher_id AND
	t.name = 'Eastern Heretic'