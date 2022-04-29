-- 2011 · Search for information on courses with more than 1000 participants
-- # Description
-- Write an SQL statement to query the information of all courses with more than 1000 students in the course table courses.

-- Table definition: courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course start date
-- teacher_id	int	teacher id

-- Number of students greater than 1000 but not including 1000
-- If the query returns no results, nothing is returned
-- Example
-- Example 1

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
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
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 2	System Design	1350	2020-07-18	3
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 12	Dynamic Programming	2000	2018-08-18	1
-- Example 2

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	200	2018-08-18	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- Sample 2 has no course records in the input data that meet the requirements of the topic and returns null, 
-- so the output contains only the table header with no data

SELECT
	*
FROM
	courses
WHERE
	student_count > 1000