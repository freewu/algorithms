-- 2043 · Search for courses created in the first quarter
-- # Description
-- Write an SQL statement to query the courses created in the first quarter (January, February, March) from the course table courses.

-- Table definition: courses (Course table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course created time
-- teacher_id	int unsigned	teacher id

-- The column name returned by the query needs to be the same as the case of the column name in the sample output
-- If the course creation time in the input data is NULL, the data will be skipped
-- If the input data is null, NULL is returned
-- Example
-- Example I

-- Table content: courses (Course schedule)

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01 09:10:12	4
-- 2	System Design	1350	2020-07-18 10:11:12	3
-- 3	Django	780	2020-02-29 12:10:12	3
-- 4	Web	340	2020-04-22 13:01:12	4
-- 5	Big Data	700	2020-09-11 16:01:12	1
-- 6	Artificial Intelligence	1660	2018-05-13 18:12:30	3
-- 7	Java P6+	780	2019-01-19 13:31:12	3
-- 8	Data Analysis	500	2019-07-12 13:01:12	1
-- 10	Object Oriented Design	300	2020-08-08 13:01:12	4
-- 12	Dynamic Programming	2000	2018-08-18 20:01:12	1
-- After running your SQL statement, the table should return.

-- name	created_at
-- Django	2020-02-29
-- Java P6+	2019-01-19
-- Example II

-- Table content: courses (Course schedule)

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	NULL	4
-- 2	System Design	1350	NULL	3
-- 3	Django	780	NULL	3
-- 4	Web	340	NULL	4
-- 5	Big Data	700	NULL	1
-- 6	Artificial Intelligence	1660	NULL	3
-- 7	Java P6+	780	NULL	3
-- 8	Data Analysis	500	NULL	1
-- 10	Object Oriented Design	300	NULL	4
-- 12	Dynamic Programming	2000	NULL	1
-- After running your SQL statement, the table should return.

-- name	created_at
-- The course creation time in Example 2 is NULL, so none of the data matches the requirement and the table is returned empty.

SELECT
	name,
	created_at
FROM
	courses
WHERE
	MONTH(created_at) IN (1,2,3)