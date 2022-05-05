-- 2033 · Query the course name and creation time in seconds for all course schedules
-- # Description
-- Write a SQL statement to query the name and creation time in seconds of courses from the course table courses and alias created_at to created_second.

-- Table definition: courses（courses table）

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	datetime	created time
-- teacher_id	int	teacher id

-- The column names returned by the query need to be the same case as the sample output.

-- Example
-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01 09:10:12	4
-- 2	System Design	1350	2020-07-18 10:11:12	3
-- 3	Django	780	2020-02-29 12:10:12	3
-- 4	Web	340	2020-04-22 13:01:12	4
-- 5	Big Data	700	2020-09-11 16:01:12	1
-- 6	Artificial Intelligence	1660	2018-05-13 18:12:30	3
-- 7	Java P6+	780	2019-01-19 13:31:12	3
-- 8	Data Analysis	500	2019-07-12 13:01:12	1
-- 10	Object Oriented Design	300	2020-08-08 13:01:12	4
-- 12	Dynamic Programming	2000	2018-08-18 20:01:12	1
-- After running your SQL statement, the table should return.

-- name	created_second
-- Senior Algorithm	12
-- System Design	12
-- Django	12
-- Web	12
-- Big Data	12
-- Artificial Intelligence	30
-- Java P6+	12
-- Data Analysis	12
-- Object Oriented Design	12
-- Dynamic Programming	12

SELECT
	name,
	CAST(DATE_FORMAT(created_at,'%s') AS SIGNED) AS created_second	
FROM
	courses