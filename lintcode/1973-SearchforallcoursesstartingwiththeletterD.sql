-- 1973 · Search for all courses starting with the letter 'D'
-- # Description
-- Write a SQL statement to query the information of all courses whose name starts with the letter 'D'.
-- Table Definition: courses (Course List)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	start time
-- teacher_id	int	teacher id



-- If there is no query results, nothing will be returned
-- Example
-- Sample I:

-- Table Contents : courses

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
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- 3	Django	780	2020/2/29	3
-- 8	Data Analysis	500	2019/7/12	1
-- 12	Dynamic Programming	2000	2018/8/18	1
-- Sample 2:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	web	800	2020-6-1	4
-- 2	java	1440	2020-7-18	3
-- 3	Cloud computing	850	2020-2-29	3
-- 4	C++	970	2020-4-22	4
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- Because there is no course information starting with 'D', the result is empty
SELECT 
	*
FROM	
	courses
WHERE
	SUBSTR(name,1,1) = 'D'