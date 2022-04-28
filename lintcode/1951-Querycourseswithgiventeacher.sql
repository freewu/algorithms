-- 1951 · Query courses with given teacher
-- # Description
-- Please write a SQL statement to query the names of all courses whose teacher id(teacher_id) is 3 in the course table.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- If there is no query result, nothing will be returned.

-- Example
-- Example 1:

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
-- After running your SQL statement, the table should return:

-- name
-- System Design
-- Django
-- Artificial Intelligence
-- Java P6+
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	2
-- 3	Django	780	2020-2-29	2
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- 6	Artificial Intelligence	1660	2018-5-13	2
-- 7	Java P6+	780	2019-1-19	1
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	1
-- After running your SQL statement, the table should return:

-- name
-- Because there is no data with the teacher id of 3 in the input sample, the returned result is empty.
SELECT 
	name
FROM 
	courses
WHERE
	teacher_id = 3