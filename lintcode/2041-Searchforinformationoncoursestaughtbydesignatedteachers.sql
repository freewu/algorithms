-- 2041 · Search for information on courses taught by designated teachers
-- # Description
-- Please write a SQL statement to query the course name, teacher id, 
-- and course creation time with teacher id (teacher_id) of 1, 2 or 3 in the course table courses, 
-- and arrange the result set in ascending order by teacher id. If the teacher id is the same, then Sorted in descending order by the time the course was created.

-- Table definition: Courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of student
-- created_at	datetime	course creation time
-- teacher_id	int	teacher id

-- Note the keywords in ORDER BY ascending and descending order

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

-- name	teacher_id	created_at
-- Big Data	1	2020-09-11
-- Data Analysis	1	2019-07-12
-- Dynamic Programming	1	2018-08-18
-- System Design	3	2020-07-18
-- Django	3	2020-02-29
-- Java P6+	3	2019-01-19
-- Artificial Intelligence	3	2018-05-13
-- Example 2

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	NULL	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	NULL	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, the table should return:

-- name	teacher_id	created_at
-- Big Data	1	2020-09-11
-- Dynamic Programming	1	2018-08-18
-- Data Analysis	1	NULL
-- System Design	3	2020-07-18
-- Java P6+	3	2019-01-19
-- Artificial Intelligence	3	2018-05-13
-- Django	3	NULL
-- The created_at field in input sample 2 is empty, and with the same teacher_id, we will put the record with NULL created_at last in descending order
SELECT
	name,
	teacher_id,
	created_at
FROM
	courses
WHERE
	teacher_id IN (1,2,3)