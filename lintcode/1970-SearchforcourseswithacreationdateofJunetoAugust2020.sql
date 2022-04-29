-- 1970 · Search for courses with a creation date of June to August 2020
-- # Description
-- Write a SQL statement that use BETWEEN AND to query the information of all courses with creation time from June 2020 to August 2020 in the course table courses.

-- Table Definition: courses (courses table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	datetime	start date
-- teacher_id	int	teacher id
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- When using the BETWEEN AND statement, please pay attention to the formatting specification, otherwise no results will be returned.
-- When using the BETWEEN AND statement to take a range, the boundary values are also included.
-- If there is no query results, nothing will be returned.
-- Example
-- Sample 1:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	mathematics	10	2020-02-05	1
-- 2	biology	20	2020-03-05	2
-- 3	chinese	30	2020-08-03	3
-- 4	chemistry	40	2020-06-01	4
-- 5	chinese	50	2020-09-02	5
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- 3	chinese	30	2020-08-03	3
-- 4	chemistry	40	2020-06-01	4
-- Sample 2:

-- Table Contents : schedules

-- id	name	student_count	created_at	teacher_id
-- 1	mathematics	10	2020-02-05	1
-- 2	biology	20	2020-03-05	2
-- 5	chinese	50	2020-09-02	5
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- Because the input sample does not have a course created between 2020-06-01 and 2020-08-31, the result is null.
SELECT 
	*
FROM
	courses
WHERE
	created_at BETWEEN '2020-06-01' AND '2020-08-31'