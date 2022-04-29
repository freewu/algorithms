-- 1984 · Count the number of different teacher_id
-- # Description
-- Write an SQL statement to count the number of different teacher ids teacher_id in the courses table courses and return the statistics with the column named teacher_count.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- If teacher_id is empty, by default the teacher_id does not exist and is not counted in the statistics
-- The name of the column that returns the statistics needs to be changed
-- Returns 0 if no data is counted
-- Example
-- Example 1:

-- Table content: courses

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

-- teacher_count
-- 3
-- Example 2:

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	
-- 2	System Design	1350	2020/7/18	
-- 10	Object Oriented Design	300	2020/8/8	
-- 12	Dynamic Programming	2000	2018/8/18	
-- After running your SQL statement, the table should return :

-- teacher_count
-- 0
-- Because the teacher_id in sample 2 is empty, the result is 0
SELECT
	COUNT(DISTINCT(teacher_id)) AS teacher_count
FROM
	courses
WHERE
	teacher_id IS NOT NULL