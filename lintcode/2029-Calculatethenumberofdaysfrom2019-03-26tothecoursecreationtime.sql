-- 2029 · Calculate the number of days from 03/26/2019 to the course creation time
-- # Description
-- Write a SQL statement to calculate the number of days difference from 26/3/2019 to the course creation time (created at) in the courses table, 
-- with the result column named date_diff.
-- Table definition: courses(courses table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int	instructor id

-- If the course creation time is earlier than March 26, 2019, the number of days returned by the calculation is negative
-- If the creation time is empty , NULL is returned
-- Example
-- Sample I.

-- Table content: courses (Course List)

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2019-03-26	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1

-- After running your SQL statement, the table should return.	

-- date_diff
-- 0
-- 480
-- 340
-- 393
-- 535
-- Example 2.

-- Table content: courses (Course List)

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	null	4
-- 2	System Design	1350	null	3

-- After running your SQL statement, the table should return.

-- date_diff
-- Because there is empty date data in Sample 2, only the title is shown here, with no data.

SELECT 
	DATEDIFF(created_at,'2019-03-26') AS date_diff
FROM
	courses 