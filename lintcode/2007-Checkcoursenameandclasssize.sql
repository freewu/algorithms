-- 2007 · Check course name and class size
-- # Description
-- Write an SQL statement to obtain the columns of the course name name and the number of students student_count from the course table courses.

-- Table definition: courses（courses table）

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course start time
-- teacher_id	int	teacher id


-- The column names returned by the query need to be the same as the case of the column names in the sample output.
-- If the column name of the SELECT does not exist, an error will occur.
-- If there is a null value in the SELECT value, null will be returned.
-- Example
-- Example 1:

-- Table Contents : courses

-- id	name	student_count	create_time	teacher_id
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
-- After running your SQL statement, the table should return.

-- name	student_count
-- Senior Algorithm	880
-- System Design	1350
-- Django	780
-- Web	340
-- Big Data	700
-- Artificial Intelligence	1660
-- Java P6+	780
-- Data Analysis	500
-- Object Oriented Design	300
-- Dynamic Programming	2000
-- Example 2:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 2	Zhang San	NULL	2021-03-05	2
-- 3	NULL	20	2021-08-03	2
-- 4	NULL	NULL	2021-06-01	4
-- After running your SQL statement, the table should return.

-- name	student_count
-- Zhang San	NULL
-- NULL	20
-- NULL	NULL
-- The value in Example 2 will return a null value if it has a null value.
SELECT
	name,
	student_count
FROM
	courses