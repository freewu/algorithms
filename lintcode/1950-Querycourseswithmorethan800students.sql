-- 1950 · Query courses with more than 800 students
-- # Description
-- Please write a SQL statement to query all courses where the number of students student_count exceeds 800 in the course table courses , 
-- and return all the course information that meets the query conditions.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- The number of students exceeds 800 but does not include 800.

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
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 12	Dynamic Programming	2000	2018-08-18	1
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Web	800	2019-08-09	3
-- 2	Database	100	2018-10-08	1
-- 3	cloud computing	590	2020-04-05	2
-- 4	C++	70	2017-05-28	4
-- 5	virtual reality	350	2020-11-21	4
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- Because the number of students without courses in the input sample exceeds 800, the returned result is empty.

SELECT 
	* 
FROM
	courses
WHERE 
	student_count > 800