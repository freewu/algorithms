-- 1979 · Check the three courses with the lowest number of students
-- # Description
-- Please write SQL statements to query the information of three courses with the least number of students (student_count) in the course table courses, and sort them in ascending order.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id



-- The result will return at most three pieces of data
-- If the query does not return any results, nothing will be returned
-- Example
-- Example 1:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	781	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- 5	Big Data	700	2020/9/11	1
-- 6	Artificial Intelligence	1660	2018/5/13	3
-- 7	Java P6+	780	2019/1/19	3
-- 8	Data Analysis	500	2019/7/12	1
-- 10	Object Oriented Design	300	2020/8/8	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 10	Object Oriented Design	300	2020/8/8	4
-- 4	Web	340	2020/4/22	4
-- 8	Data Analysis	500	2019/7/12	1
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	881	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- 5	Big Data	780	2020/9/11	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 4	Web	340	2020/4/22	4
-- 5	Big Data	780	2020/9/11	1
-- 1	Advanced Algorithms	880	2020/6/1	4
SELECT
	*
FROM
	courses
ORDER BY
	student_count ASC
LIMIT 3