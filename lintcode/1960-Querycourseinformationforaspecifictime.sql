-- 1960 · Query course information for a specific time
-- # Description
-- Please write a SQL statement to query the information of all courses which created at 2021-01-01 or 2021-01-03 in the course table courses.

-- Table definition: courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- If there is no query result, nothing will be returned

-- Example
-- Example 1:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	english	61	2021-1-3	3
-- 2	mathematics	62	2021-1-10	8
-- 3	english	58	2021-1-29	10
-- 4	physics	53	2021-1-10	8
-- 5	biology	43	2021-1-19	5
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 1	english	61	2021-1-3	3
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	english	52	2021-1-27	8
-- 2	mathematics	58	2021-1-24	8
-- 3	english	58	2021-1-29	4
-- 4	chemistry	62	2021-1-18	10
-- 5	chinese	65	2021-1-5	7
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- Because there are no courses starting on 2021-01-01 or 2021-01-03 in the input sample, the results are not returned.
SELECT
	*
FROM
	courses
WHERE
	created_at = '2021-01-01'  OR 
	created_at = '2021-01-03'