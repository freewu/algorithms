-- 1964 · Query for course information about the number of students within the specified range
-- # Description
-- Please write a SQL statement to query the information of all courses where the number of students is between 50 and 55 in the course table Courses.

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
-- 1	english	50	2021-1-3	5
-- 2	biology	53	2021-1-5	1
-- 3	chinese	58	2021-1-16	3
-- 4	english	50	2021-1-8	2
-- 5	chinese	58	2021-1-10	8
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 1	english	50	2021-1-3	5
-- 2	biology	53	2021-1-5	1
-- 4	english	50	2021-1-8	2
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	biology	59	2021-1-13	10
-- 2	chemistry	58	2021-1-17	7
-- 3	biology	56	2021-1-28	9
-- 4	mathematics	57	2021-1-1	5
-- 5	chemistry	59	2021-1-19	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- Because there are no enrolled courses between 50 and 55 students in the input sample, the results are not returned.

-- use BETWEEN
SELECT
	*
FROM
	courses
WHERE
	student_count BETWEEN 50 AND 55

-- use >= & <=
SELECT
	*
FROM
	courses
WHERE
	student_count >= 50 AND 
	student_count <= 55