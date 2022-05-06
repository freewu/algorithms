-- 2038 · Search for the date and time in the course creation time of the course schedule respectively
-- # Description
-- Write SQL statements to query the course name name and creation time created_at from the course table courses, 
-- extract the date and time from the course creation time created_at, and use created_date and created_time as the result set column names.

-- Table definition: courses

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course created time
-- teacher_id	int unsigned	teacher id

-- The column name returned by the query should be the same as the case of the column name output by the sample
-- If the course creation time in the input data is NULL, the date and time returned are both NULL
-- If the input data is null, NULL will be returned.
-- Example
-- Example I

-- Table content: courses (Course schedule)

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
-- After running your SQL statement, the table should return

-- name	created_at	created_date	created_time
-- Senior Algorithm	2020-06-01T09:10:12	2020-06-01	09:10:12
-- System Design	2020-07-18T10:11:12	2020-07-18	10:11:12
-- Django	2020-02-29T12:10:12	2020-02-29	12:10:12
-- Web	2020-04-22T13:01:12	2020-04-22	13:01:12
-- Big Data	2020-09-11T16:01:12	2020-09-11	16:01:12
-- Artificial Intelligence	2018-05-13T18:12:30	2018-05-13	18:12:30
-- Java P6+	2019-01-19T13:31:12	2019-01-19	13:31:12
-- Data Analysis	2019-07-12T13:01:12	2019-07-12	13:01:12
-- Object Oriented Design	2020-08-08T13:01:12	2020-08-08	13:01:12
-- Dynamic Programming	2018-08-18T20:01:12	2018-08-18	20:01:12
-- Example II

-- Table content: courses (Course schedule)

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
-- After running your SQL statement, the table should return

-- name	created_at	created_date	created_time
-- Senior Algorithm	NULL	NULL	NULL
-- System Design	NULL	NULL	NULL
-- Django	NULL	NULL	NULL
-- Web	NULL	NULL	NULL
-- Big Data	NULL	NULL	NULL
-- Artificial Intelligence	NULL	NULL	NULL
-- Java P6+	NULL	NULL	NULL
-- Data Analysis	NULL	NULL	NULL
-- Object Oriented Design	NULL	NULL	NULL
-- Dynamic Programming	NULL	NULL	NULL
-- The course creation time created_at in Example 2 is NULL, so the query result created_date and created_time are both NULL

SELECT
	name,
	created_at,
	DATE_FORMAT(created_at,'%Y-%m-%d') AS created_date,
	DATE_FORMAT(created_at,'%H:%i:%s') AS created_time
FROM
	courses