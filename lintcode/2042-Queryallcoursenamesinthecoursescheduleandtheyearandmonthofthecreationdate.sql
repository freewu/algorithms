-- 2042 · Query all course names in the course schedule and the year and month of the creation date
-- # Description
-- Question Description: Write a SQL statement to query the course name (name)
-- and the year (alias: year ) and month (alias: month) of the creation date of all courses in the course table courses.

-- Table definition: courses (course table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	datetime	Course creation time
-- teacher_id	int unsigned	instructor id

-- The column names returned by the query need to match the case of the column names output by the sample
-- If there is a null value in the SELECT value, null will be returned.
-- Example
-- Sample I

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01 09:03:12	4
-- 2	System Design	1350	2020-07-18 10:03:12	3
-- 3	Django	780	2020-02-29 12:03:12	3
-- 4	Web	340	2020-04-22 13:03:12	4
-- 5	Big Data	700	2020-09-11 16:03:12	1
-- 6	Artificial Intelligence	1660	2018-05-13 18:03:30	3
-- 7	Java P6+	780	2019-01-19 13:03:12	3
-- 8	Data Analysis	500	2019-07-12 13:03:12	1
-- 10	Object Oriented Design	300	2020-08-08 13:03:12	4
-- 12	Dynamic Programming	2000	2018-08-18 20:03:12	1
-- After running your SQL statement, the table should return.

-- name	year	month
-- Advanced Algorithms	2020	6
-- System Design	2020	7
-- Advanced Algorithms	2020	6
-- Web	2020	4
-- Big Data	2020	4
-- Artificial Intelligence	2018	5
-- Java P6+	2019	1
-- Data Analysis	2019	7
-- Data Analysis	2019	7
-- Dynamic Programming	2018	8
-- Sample II		
-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	NULL	4
-- 2	System Design	1350	NULL	3
-- 3	Django	780	NULL	3
-- 4	Web	340	NULL	4
-- 5	Big Data	700	NULL	1
-- 6	Artificial Intelligence	1660	NULL	3
-- 7	Java P6+	780	NULL	3
-- 8	Data Analysis	500	NULL	1
-- 10	Object Oriented Design	300	NULL	4
-- 12	Dynamic Programming	2000	NULL	1
-- After running your SQL statement, the table should return.

-- name	year	month
-- Advanced Algorithms	NULL	NULL
-- System Design	NULL	NULL
-- Django	NULL	NULL
-- Web	NULL	NULL
-- Big Data	NULL	NULL
-- Artificial Intelligence	NULL	NULL
-- Java P6+	NULL	NULL
-- Data Analysis	NULL	NULL
-- Object Oriented Design	NULL	NULL
-- Dynamic Programming	NULL	NULL

-- with DATE_FORMAT
SELECT
	name,
	CAST(DATE_FORMAT(created_at,'%Y') AS SIGNED) AS year,
	CAST(DATE_FORMAT(created_at,'%m') AS SIGNED) AS month
FROM
	courses

-- with YEAR() OR MONTH()
SELECT
	name,
	YEAR(created_at) AS year,
	MONTH(created_at) AS month
FROM
	courses