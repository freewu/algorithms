-- 1988 · Check the number of students in the course with the lowest number of students
-- # Description
-- Write an SQL statement to query the minimum number of students (student_count) of all courses in the course table courses, and named after min_student_count

-- Table definition: courses(课程表)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	total number of students
-- created_at	date	class start time
-- teacher_id	int	teacher ID
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- If the query does not return any results, nothing is returned
-- Example
-- Sample I:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	780	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- 10	Object Oriented Design	300	2020/8/8	4
-- After running your SQL statement, the table should return :

-- min_student_count
-- 300
-- Sample 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	web	800	2020-6-1	4
-- 2	java	800	2020-7-18	3
-- 3	Django	850	2020-2-29	3
-- 4	C++	970	2020-4-22	4
-- After running your SQL statement, the table should return :

-- min_student_count
-- 800
-- Although there are two identical min values, it still only returns this one result
SELECT
	MIN(student_count) AS min_student_count
FROM
	courses