-- 2001 · Query the course information of 'Web' or 'Big Data'
-- # Description
-- Please write a SQL statement, select the course information of the course named'Web' or'Big Data' from the courses table. 
-- If these two courses exist, please return all the information of the two courses.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course start time
-- teacher_id	int	teacher id



-- If there is no query result, nothing will be returned
-- Example
-- Example 1:

-- Table content : courses

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
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	3
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	2019-1-19	3
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- Because there is no data to be queried in the data, only the title is displayed here, no data
SELECT
	* 
FROM
	courses
WHERE
	name IN ('Web','Big Data')