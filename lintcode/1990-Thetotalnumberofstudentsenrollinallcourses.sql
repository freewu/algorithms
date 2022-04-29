-- 1990 · The total number of students enroll in all courses
-- # Description
-- Please write SQL statements to count the total number of students of all courses in the course table courses, and use all_student_count as the result set column name.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	course start time
-- teacher_id	int	teacher id



-- If the query does not return any results, nothing is returned.

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

-- all_student_count
-- 277
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	english	0	2021-1-27	8
-- 2	mathematics	0	2021-1-24	8
-- After running your SQL statement, the table should return:

-- all_student_count
-- 0

SELECT	
	SUM(student_count) AS all_student_count
FROM
	courses