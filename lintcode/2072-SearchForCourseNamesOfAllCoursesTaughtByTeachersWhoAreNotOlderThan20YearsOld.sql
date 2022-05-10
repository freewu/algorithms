-- 2072 · Search for course names of all courses taught by teachers who are not older than 20 years old
-- Description
-- Write an SQL statement to query the name of all courses taught by teachers
-- who are not older than 20 years old in the courses table by combining the teachers table and the courses table.
-- Table definition: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	instructor's name
-- email	varchar	Instructor's email
-- age	int	instructor's age
-- country	varchar	instructor's nationality
-- Table definition: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	datetime	course creation time
-- teacher_id	int unsigned	instructor id

-- If there is an age of null in teachers, the data will be skipped.
-- If the age of all teachers are null, or if the teachers table is empty, then return null.
-- Example
-- Sample I:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents: Courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1 09:03:12	4
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
-- Big Data
-- Data Analysis
-- Dynamic Programming
-- Example 2:

-- Table Contents : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	21	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	21	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	21	CN
-- Table Contents: Courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1 09:03:12	4
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
-- Because there is no teacher in the input sample that is not older than 20 years old, 
-- the queried course name is empty, so only the title is shown here and there is no data.

-- Write your SQL Query here --
SELECT
	name
FROM
	courses
WHERE
	teacher_id IN (
		SELECT 
			id
		FROM
			`teachers`
		WHERE
			age <= 20
	)