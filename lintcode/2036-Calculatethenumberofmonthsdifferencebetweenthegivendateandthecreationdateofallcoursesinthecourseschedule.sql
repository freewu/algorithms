-- 2036 · Calculate the number of months difference between the given date and the creation date of all courses in the course schedule
-- # Description
-- Write a SQL statement to calculate the difference in months between '2020-04-22' and the course creation date in the courses table, with the returned column named MonthDiff.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- The query needs to return the same column names as the sample output.

-- Example
-- Sample I.

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
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

-- MonthDiff
-- -1
-- -2
-- 1
-- 0
-- -4
-- 23
-- 15
-- 9
-- -3
-- 20
-- The course start time of the course in Sample 1 is equal to the specified time and the result is 0. The result is negative if it is less than the specified time.

-- Example 2.

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2022-06-01	4
-- 2	System Design	1350	2023-07-18	3
-- After running your SQL statement, the table should return.

-- MonthDiff
-- null
-- null
-- The course start time in Sample 2 is less than the initial time, so the statistics are negative

SELECT 
	TIMESTAMPDIFF(MONTH,created_at,'2020-04-22') AS MonthDiff
FROM courses

-- TIMESTAMPDIFF(unit,begin,end); 根据单位返回时间差,对于传入的begin和end不需要相同的数据结构,可以存在一个为Date一个DateTime
-- Unit 支持的单位有
-- MICROSECOND 毫秒
-- SECOND 秒
-- MINUTE 分
-- HOUR 小时
-- DAY 天 
-- WEEK 周
-- MONTH 月
-- QUARTER 季
-- YEAR 年