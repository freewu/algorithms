-- 2022 · Delayed all course creation dates by one year
-- # Description
-- Write a SQL statement, select the course creation date and delay one year later in the course table courses, and finally return the course name name and the modified creation date, which is named new created.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	date	course creation date
-- teacher_id	int	teacher id

-- -Change only the year when the course was created
-- -If the creation date is the 29th of February in leap year,after postponing,the date will be the last day of February in the following year,that is 28th.

-- Example
-- Eample 1 ：

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
-- After running your SQL statement, the table should return :

-- name	new_created
-- Advanced Algorithms	2021-6-1
-- System Design	2021-7-18
-- Django	2021-2-29
-- Web	2021-4-22
-- Big Data	2021-9-11
-- Artificial Intelligence	2019-5-13
-- Java P6+	2020-1-19
-- Data Analysis	2020-7-12
-- Object Oriented Design	2020-8-8
-- Dynamic Programming	2019-8-18
-- Eample 1 ：

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	2019-1-19	3
-- 3	Django	780	2020-2-29	3
-- After running your SQL statement, the table should return :

-- name	new_created
-- Advanced Algorithms	2021-6-1
-- Artificial Intelligence	2019-5-13
-- Java P6+	2020-1-19
-- Django	2021-2-28

-- DATE_ADD('2018-03-22 10:23:00', INTERVAL 30 MINUTE)
-- DATE_ADD(<date>, INTERVAL <num> <TYPE>)
-- date 只要是合法的日期表达式即可
-- num 是希望添加的时间间隔值
-- TYPE 是时间间隔的单位
-- MICROSECOND
-- SECOND
-- MINUTE
-- HOUR
-- DAY
-- WEEK
-- MONTH
-- QUARTER
-- YEAR
-- SECOND_MICROSECOND
-- MINUTE_MICROSECOND
-- MINUTE_SECOND
-- HOUR_MICROSECOND
-- HOUR_SECOND
-- HOUR_MINUTE
-- DAY_MICROSECOND
-- DAY_SECOND
-- DAY_MINUTE
-- DAY_HOUR
-- YEAR_MONTH

SELECT
	name,
	DATE_ADD(created_at,INTERVAL 1 YEAR) AS new_created
FROM
	courses