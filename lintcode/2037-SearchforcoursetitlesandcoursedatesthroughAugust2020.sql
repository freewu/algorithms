-- 2037 · Search for course titles and course dates through August 2020
-- # Description
-- Write a SQL statement to query the name and creation date of courses which created before August 2020 
-- from the course table courses and alias the column name for the creation date as created_date 
-- (the date refers to the part of created_at that does not include a specific time).

-- Table Definition: courses (course table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	datetime	Creation time
-- teacher_id	int	instructor id

-- -The column name returned by the query should be consistent with the case of the column name output by the sample

-- Before August 2020 is not including August 2020.
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

-- name	created_date
-- Senior Algorithm	2020-06-01
-- System Design	2020-07-18
-- Django	2020-02-29
-- Web	2020-04-22
-- Artificial Intelligence	2018-05-13
-- Java P6+	2019-01-19
-- Data Analysis	2019-07-12
-- Data Analysis	2019-07-12
-- Dynamic Programming	2018-08-18
-- Example 2

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	null	4
-- 2	System Design	1350	null	3
-- 3	Django	780	null	3
-- 4	Web	340	null	4
-- 5	Big Data	700	null	1
-- 6	Artificial Intelligence	1660	null	3
-- 7	Java P6+	780	null	3
-- 8	Data Analysis	500	null	1
-- 10	Object Oriented Design	300	null	4
-- 12	Dynamic Programming	2000	null	1
-- After running your SQL statement, the table should return.

-- name	created_date
-- Because there is no eligible data in the input sample, only the title is shown here, no data.
SELECT
	name,
	DATE_FORMAT(created_at,'%Y-%m-%d') AS created_date
FROM
	courses
WHERE
	created_at < '2020-08-01'