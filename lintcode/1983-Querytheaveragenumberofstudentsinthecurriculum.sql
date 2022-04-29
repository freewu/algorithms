-- 1983 · Query the average number of students in the curriculum
-- # Description
-- Write an SQL statement to query the average number of students student_count in the courses table courses and return the result column named as average_student_count.

-- Table Definition: courses (Course Schedule)

-- Column Name	Type	Comments
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	class start time
-- teacher_id	int	teacher id



-- If student_count is empty, nothing is returned
-- The name of the column that returns the statistics needs to be changed
-- If the query does not return any results, nothing will be returned
-- Example
-- Sample I:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-28	3
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- After running your SQL statement, the table should return.

-- average_student_count
-- 810
-- Sample 2:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	0	2020-6-1	4
-- 2	System Design	0	2020-7-18	3
-- 3	Django	0	2020-2-28	3
-- After running your SQL statement, the table should return.

-- average_student_count
-- 0

SELECT
	SUM(student_count) / COUNT(*) AS average_student_count
FROM
	courses 