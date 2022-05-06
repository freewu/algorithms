-- 2065 · Check the course names of all courses taught by all teachers who are older than 20 years old
-- # Description
-- Write an SQL statement that combine the teacher table 
-- and the course table to query the names of all courses taught by teachers older than 20 years old in the course table courses.
-- Table definition: teachers (teachers table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table definition: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course creation time
-- teacher_id	int unsigned	teacher id

-- If the age is null in teachers, the data will be skipped.
-- If all of age is null in teachers, or if the teachers table is empty, then return null.
-- Example
-- Sample I:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1 09:03:12	4
-- 2	System Design	1350	2020-7-18 10:03:12	3
-- 3	Django	780	2020-2-29 12:03:12	3
-- 4	Web	340	2020-4-22 13:03:12	4
-- 5	Big Data	700	2020-9-11 16:03:12	1
-- 6	Artificial Intelligence	1660	2018-5-13 18:03:12	3
-- 7	Java P6+	780	2019-1-19 13:03:12	3
-- 8	Data Analysis	500	2019-7-12 13:03:12	1
-- 10	Object Oriented Design	300	2020-8-8 13:03:12	4
-- 12	Dynamic Programming	2000	2018-8-18 20:03:12	1
-- After running your SQL statement, the table should return.

-- name
-- Advanced Algorithms
-- System Design
-- Django
-- Web
-- Artificial Intelligence
-- Java P6+
-- Object Oriented Design
-- Sample 2:

-- Table Contents : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	18	CN
-- 3	Western Venom	western.venom@163.com	19	USA
-- 4	Southern Emperor	southern.emperor@qq.com	19	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1 09:03:12	4
-- 4	Web	340	2020-4-22 13:03:12	4
-- 5	Big Data	700	2020-9-11 16:03:12	1
-- 8	Data Analysis	500	2019-7-12 13:03:12	1
-- 10	Object Oriented Design	300	2020-8-8 13:03:12	4
-- 12	Dynamic Programming	2000	2018-8-18 20:03:12	1
-- After running your SQL statement, the table should return.

-- name
-- Because there is no teacher over 20 years old in the input sample, and the course name queried is empty, only the title is shown here, and there is no data.
SELECT
	name
FROM
	courses
WHERE
	teacher_id IN (
		SELECT
			id 
		FROM
			teachers 
		WHERE
			age > 20
	)