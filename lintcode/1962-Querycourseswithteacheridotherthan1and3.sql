-- 1962 · Query courses with teacher id other than 1 and 3
-- # Description
-- Write an SQL statement to query all courses whose teacher ID teacher_id is not 1 or 3 in the course table courses and return the courses' name that meet the query criteria.

-- The courses table is defined as follows：

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- The question requires the use of the NOT IN operator
-- If there is no query result, nothing will be returned
-- Example
-- Eample 1 ：

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	780	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- 5	Big Data	700	2020/9/11	1
-- 6	Artificial Intelligence	1660	2018/5/13	3
-- 7	Java P6+	780	2019/1/19	3
-- 8	Data Analysis	500	2019/7/12	1
-- 10	Object Oriented Design	300	2020/8/8	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- After running your SQL statement, the table should return :

-- name
-- Advanced Algorithms
-- Web
-- Object Oriented Design
-- Example 2 ：

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Web	800	2019-08-09	3
-- 2	Database	1440	2018-10-08	1
-- 3	Cloud computing	850	2020-04-05	3
-- 4	C++	970	2017-05-28	3
-- After running your SQL statement, the table should return :

-- name
-- Because there is no teacher whose id is not 1 or 3 in the input sample, the returned result is null

SELECT
	name
FROM
	courses
WHERE
	teacher_id NOT IN (1,3)