-- 2032 · Advance all course creation dates by one day
-- # Description
-- Write an SQL statement to query the course creation date of a course in the courses table, 
-- advancing the course creation date by one day and returning the course id, 
-- course name name and the revised course start date, with the revised course creation date named new_created.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- The result column name needs to be renamed
-- If the creation time is null, it can not be modified
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

-- id	name	new_created
-- 1	Advanced Algorithms	2020-5-31
-- 2	System Design	2020-7-17
-- 3	Django	2020-2-28
-- 4	Web	2020-4-21
-- 5	Big Data	2020-9-10
-- 6	Artificial Intelligence	2018-5-12
-- 7	Java P6+	2019-1-18
-- 8	Data Analysis	2019-7-11
-- 10	Object Oriented Design	2020-8-7
-- 12	Dynamic Programming	2018-8-17
-- Eample 2：

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	null	3
-- After running your SQL statement, the table should return :

-- id	name	created_at
-- 1	Advanced Algorithms	2020-5-31
-- 6	Artificial Intelligence	2018-5-12
-- 7	Java P6+	null

SELECT
	id,
	name,
	DATE_ADD(created_at,INTERVAL -1 DAY) AS new_created
FROM
	courses