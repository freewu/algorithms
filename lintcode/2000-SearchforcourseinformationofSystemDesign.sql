-- 2000 · Search for course information of System Design
-- # Description
-- Write an SQL statement to query the information of course named System Design in the course table courses.

-- Table Definition: courses (course table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course start time
-- teacher_id	int	teacher id
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- If the query does not return any results, nothing is returned.

-- Example
-- Sample I:

-- Table content: courses

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
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- 2	System Design	1350	2020-7-18	3
-- Sample 2:

-- Table Contents: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	Java P6+	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	3
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- There is no eligible data in Sample 2, so the output contains only table headers and no data.
SELECT
	*
FROM
	courses
WHERE
	name = 'System Design'