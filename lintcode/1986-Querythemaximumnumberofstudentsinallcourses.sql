-- 1986 · Query the maximum number of students in all courses
-- # Description
-- Write an SQL statement that queries the course table courses for the maximum number of students (student_count) in all courses and names the output field max_student_count.

-- Table definition: courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	datetime	course start time
-- teacher_id	int	teacher id



-- The column names returned by the query need to match the case of the column names in the sample output
-- If the number of students is NULL in the input data, then the data is skipped
-- Returns NULL if all the number of student in the input data are NULL, or if the input data is empty
-- Example
-- Example 1

-- Table content: courses

-- id	name	student_count	create_time	teacher_id
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

-- max_student_count
-- 2000
-- Example 2

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	NULL	2020-06-01	4
-- 2	System Design	NULL	2020-07-18	3
-- 3	Django	NULL	2020-02-29	3
-- 4	Web	NULL	2020-04-22	4
-- 5	Big Data	NULL	2020-09-11	1
-- 6	Artificial Intelligence	NULL	2018-05-13	3
-- 7	Java P6+	NULL	2019-01-19	3
-- 8	Data Analysis	NULL	2019-07-12	1
-- 10	Object Oriented Design	NULL	2020-08-08	4
-- 12	Dynamic Programming	NULL	2018-08-18	1
-- After running your SQL statement, the table should return:

-- max_student_count
-- NULL

SELECT
	MAX(student_count) AS max_student_count
FROM
	courses