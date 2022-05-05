-- 2023 · Calculate the number of days from 01/13/2018 to the time of course creation
-- # Description
-- Write a SQL statement to calculate the number of days difference from 01/13/2018 to the course creation time (created_at) from the courses table, with the result column named date_diff.

-- Table definition: courses (course table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student count	int	number of students
-- created at	date	course creation time
-- teacher id	int	teacher id


-- If the course creation time is earlier than 13/01/2018, the number of days returned by the calculation is negative
-- If the creation time is empty, NULL is returned
-- Example
-- Sample I.

-- Table content: courses (Course List)

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2018-01-13	4
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

-- date_diff
-- 0
-- 917
-- 777
-- 830
-- 972
-- 120
-- 371
-- 545
-- 938
-- 217
-- Example 2.

-- Table content: courses (Course List)

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	null	4
-- 2	System Design	1350	null	3
-- After running your SQL statement, the table should return.

-- date_diff
-- Because there is empty date data in Sample 2, only the title is shown here, with no data.


-- Sql Server 第二个参数减第一个参数 有三个参数
-- SELECT DATEDIFF(day,‘2008-12-29‘,‘2008-12-30‘) AS DiffDate
-- 1
-- SELECT DATEDIFF(day,‘2008-12-30‘,‘2008-12-29‘) AS DiffDate
-- -1

-- DATEDIFF(datepart,startdate,enddate)
-- startdate 和 enddate 参数是合法的日期表达式。

-- datepart 参数可以是下列的值：
-- datepart	缩写
-- 年	yy, yyyy
-- 季度	qq, q
-- 月	mm, m
-- 年中的日	dy, y
-- 日	dd, d
-- 周	wk, ww
-- 星期	dw, w
-- 小时	hh
-- 分钟	mi, n
-- 秒	ss, s
-- 毫秒	ms
-- 微妙	mcs
-- 纳秒	ns

SELECT 
	DATEDIFF(created_at,'2018-01-13') AS date_diff
FROM
	courses