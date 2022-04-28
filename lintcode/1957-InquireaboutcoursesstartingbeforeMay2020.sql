-- 1957 · Inquire about courses starting before May 2020
-- # Description
-- Please write a SQL statement to query the name and creation time created_at of courses which were created at '2020-01-01'’（included) to '2020-05-01'（excluded）in the course table courses .

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Query for the course name and creation time
-- Course creation time includes '2020-01-01' but not '2020-05-01'
-- If there is no query result, nothing will be returned
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

-- name	created_at
-- Django	2020-02-29
-- Web	2020-04-22
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2021-2-24	3
-- 4	Web	340	2020-7-22	4
-- 5	Big Data	700	2020-9-11	1
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	2019-1-19	3
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	1
-- After running your SQL statement, the table should return:

-- name	created_at
-- Because there is no data to be queried in the data, only the title is displayed here, no data
SELECT 
	name,
	created_at
FROM
	courses
WHERE
	created_at >= '2020-01-01' AND
	created_at <  '2020-05-01'