-- 2035 · Calculate the number of years difference between the start date and the current date of all courses in the course schedule
-- Description
-- Write a SQL statement to query the courses table to calculate the year difference between the time a course was created (created_at) and April 1, 2021,
-- with the course name column displayed as courses_name, the time a course was created as courses_created_at, and the year difference column name displayed as year_diff.

-- Table Definition: courses (course table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total_students
-- created_at	date	Course creation time
-- teacher_id	int	instructor id

-- This problem is calculated by rounding down the number of years difference, for example: 2000-03-06 - 2001-03-05 is less than a year, according to 0 years processing
-- If there is a course record creation time data is empty, return NULL
-- Example
-- Sample I:

-- Table content: courses (course list)

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, the table should return.

-- courses_name	courses_created_at	year_diff
-- Advanced Algorithms	2020-06-01	0
-- System Design	2020-07-18	0
-- Django	2020-02-29	1
-- Web	2020-04-22	0
-- Big Data	2020-09-11	0
-- Artificial Intelligence	2018-05-13	2
-- Java P6+	2019-01-19	2
-- Data Analysis	2019-07-12	1
-- Object Oriented Design	2020-08-08	0
-- Dynamic Programming	2018-08-18	2
-- Sample 2:

-- Table content: courses (course schedule)

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	null	4
-- 2	System Design	1350	null	3
-- After running your SQL statement, the table should return.

-- courses_name	courses_created_at	year_diff
-- Because sample 2 has empty date data, only the title is shown here, no data.

SELECT
	name AS courses_name,
	created_at AS courses_created_at,
	FLOOR(ABS(DATEDIFF('2021-04-01',created_at)) / 365) AS year_diff
FROM
	courses 

-- ROUND(X) -- 表示将值 X 四舍五入为整数，无小数位
-- ROUND(X,D) -- 表示将值 X 四舍五入为小数点后 D 位的数值，D为小数点后小数位数。若要保留 X 值小数点左边的 D 位，可将 D 设为负值。
-- FLOOR(X)表示向下取整，只返回值X的整数部分，小数部分舍弃。
-- CEILING(X) 表示向上取整，只返回值X的整数部分，小数部分舍弃。