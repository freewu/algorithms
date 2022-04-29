-- 2014 · Query course name and course creation time
-- # Description
-- Write an SQL statement to query the names and creation time of all courses in the courses table courses.

-- Table definition: courses (courses table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course creation time
-- teacher_id	int	instructor id


-- If the query does not return any results, nothing is returned.

-- Example
-- Example 1：

-- Table Content：courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	3
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	2019-1-19	3
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	1
-- After running your SQL statement, the table should return:

-- name	created_at
-- Advanced Algorithms	2020-6-1
-- System Design	2020-7-18
-- Django	2020-2-29
-- Web	2020-4-22
-- Big Data	2020-9-11
-- Artificial Intelligence	2018-5-13
-- Java P6+	2019-1-19
-- Data Analysis	2019-7-12
-- Object Oriented Design	2020-8-8
-- Dynamic Programming	2018-8-18
-- Example 2:

-- Table content：courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	3
-- After running your SQL statement, the table should return:

-- name	created_at
-- Advanced Algorithms	2020-6-1
-- System Design	2020-7-18
-- Django	2020-2-29
SELECT 
	name,
	created_at
FROM
	courses