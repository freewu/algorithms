-- 1954 · Query course information with the number of students within a specified range
-- # Description
-- Please write a SQL statement to query the information of all courses where the students population are between 800 (inclusive) and 1000 (exclusive) in the course table courses.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Tip:
-- 1.The number of students include 800，exclude 1000.
-- 2. If there is no query result, nothing will be returned

-- Example
-- Example 1:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	780	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- 5	Big Data	700	2020/9/11	1
-- 6	Artificial Intelligence	1660	2018/5/13	3
-- 7	Java P6+	780	2019/1/19	3
-- 8	Data Analysis	500	2019/7/12	1
-- 10	Object Oriented Design	300	2020/8/8	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-06-01	4
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	800	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	800	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1060	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	1000	2018-08-18	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- Because there is no data to be queried in the data, only the title is displayed here, no data
SELECT 
	*
FROM 
	courses
WHERE
	student_count >= 800 AND
	student_count < 1000