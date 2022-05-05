-- 2019 · Delete all rows in the table
-- Description
-- Write an SQL statement to delete the data of all rows in the courses table without deleting the table.

-- Table Definition: courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course start time
-- teacher_id	int	teacher id

-- Be extra careful when deleting records! Because you can't start over!
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
-- After the DELETE code is executed, we will execute SELECT * FROM courses and the table should return.

-- id	name	student_count	created_at	teacher_id

DELETE FROM 
    `courses`
WHERE
    1 = 1
