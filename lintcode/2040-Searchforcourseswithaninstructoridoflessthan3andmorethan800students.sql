-- 2040 · Search for courses with an instructor id of less than 3 and more than 800 students
-- # Description
-- Query all courses whose instructor id is not 3 and the number of students is more than 800 in the course table courses, 
-- and finally return all information of the courses that satisfy the conditions.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id


-- Please use the logical operator NOT for this question
-- Number of students over 800 but not including 800
-- If the query fails to find the value that meets the conditions, the return result is null
-- Example
-- Example 1：

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
-- After running your SQL statement, the table should return :

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- Example 2：

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	300	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	780	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- After running your SQL statement, the table should return :

-- id	name	student_count	created_at	teacher_id
-- Because there is no data to be queried in the data, only the title is shown here, no data.
SELECT
	*
FROM
	courses
WHERE
	student_count > 800 AND
	teacher_id != 3