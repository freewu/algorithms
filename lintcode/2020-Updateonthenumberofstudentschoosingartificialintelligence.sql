-- 2020 · Update on the number of students choosing artificial intelligence
-- # Description
-- Write an SQL statement to update the number of students to 500 of the Artificial Intelligence course in the course table courses.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- Please change the value directly, do not change the value by adding or subtracting
-- If the number of students is null, update the data directly to the set value
-- If there is no artificial intelligence course, then do not update any data
-- Example
-- Eample 1 ：

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
-- After running your SQL statement, we will execute the following statement :

-- SELECT *
-- FROM `courses`
-- Return ：

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	780	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- 5	Big Data	700	2020/9/11	1
-- 6	Artificial Intelligence	500	2018/5/13	3
-- 7	Java P6+	780	2019/1/19	3
-- 8	Data Analysis	500	2019/7/12	1
-- 10	Object Oriented Design	300	2020/8/8	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- Eample 2 ：

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 4	Web	340	2020/4/22	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- After running your SQL statement, we will execute the following statement :

-- SELECT *
-- FROM `courses`
-- Return ：

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 4	Web	340	2020/4/22	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- There is no Artificial Intelligence course in Sample 2, so the data are not changed.

UPDATE
	`courses`
SET
	student_count = 500
WHERE
	name = 'Artificial Intelligence'