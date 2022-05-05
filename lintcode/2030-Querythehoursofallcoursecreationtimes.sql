-- 2030 · Query the hours of all course creation times
-- # Description
-- Write a SQL statement that queries the course name (name) and the hours of the course creation time (created_at) from the course table courses, 
-- and aliases the column name from which the hours are extracted to created_hour.

-- Table definition: courses（courses table）

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	datetime	creation time
-- teacher_id	int	teacher id

-- The column names returned by the query need to be the same case as the sample output.
-- If there is a NULL value in the input data, NULL is returned.
-- Example
-- Example 1

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

-- name	created_hour
-- Senior Algorithm	9
-- System Design	10
-- Django	12
-- Web	13
-- Big Data	16
-- Artificial Intelligence	18
-- Java P6+	13
-- Data Analysis	13
-- Object Oriented Design	13
-- Dynamic Programming	20
-- Example 2

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	NULL	880	NULL	4
-- 2	IDE	300	NULL	4
-- 3	NULL	2000	2018-08-18 20:01:12	1
-- After running your SQL statement, the table should return.

-- name	created_hour
-- NULL	NULL
-- IDE	NULL
-- NULL	12
-- If the value in Example 2 has a null value, the null value will be returned.

SELECT
	name,
	CAST(DATE_FORMAT(created_at,'%H') AS SIGNED) AS created_hour
FROM
	courses