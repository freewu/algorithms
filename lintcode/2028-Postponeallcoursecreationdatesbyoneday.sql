-- 2028 · Postpone all course creation dates by one day
-- # Description
-- Write an SQL statement to query the course creation date of a course in the courses table and defer the course creation date by one day, 
-- returning the course name name and the modified course creation time named new_created.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- -The name of result column needs to be renamed

-- Example
-- Eample 1：

-- Table content : courses

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
-- After running your SQL statement, the table should return :

-- name	new_created
-- Advanced Algorithms	2020-06-02
-- System Design	2020-07-19
-- Django	2020-03-01
-- Web	2020-04-23
-- Big Data	2020-09-12
-- Artificial Intelligence	2018-05-14
-- Java P6+	2019-01-20
-- Data Analysis	2029-07-13
-- Object Oriented Design	2020-08-09
-- Dynamic Programming	2018-08-19
-- Eample 2：

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-2	4
-- 6	Artificial Intelligence	1660	2018-5-14	3
-- 7	Java P6+	780	2019-1-19	3
-- After running your SQL statement, the table should return :

-- name	new_created
-- Advanced Algorithms	2020-06-3
-- Artificial Intelligence	2018-05-15
-- Java P6+	2019-01-20

SELECT
	name,
	DATE_ADD(created_at,INTERVAL 1 DAY) AS new_created
FROM
	courses
