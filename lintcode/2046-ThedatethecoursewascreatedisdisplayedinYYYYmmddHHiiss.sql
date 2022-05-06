-- 2046 · The date the course was created is displayed in 'year-month-day hour:minute:second'
-- # Description
-- Write an SQL statement to query the course creation time in the courses table 
-- and the result returned in the format ‘Year-Month-Day Hour:Minute:Second’ with the column named DATE_FORMAT.

-- Table definition: courses

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	Course creation time
-- teacher_id	int	lecturer id

-- The column names returned by the query need to match the case of the column names in the sample output
-- Output format as YYYY-MM-DD HH:mm:SS

-- Example
-- Sample I.

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1 09:03:12	4
-- 2	System Design	1350	2020-7-18 10:03:12	3
-- 3	Django	780	2020-2-29 12:03:12	3
-- 4	Web	340	2020-4-22 13:03:12	4
-- 5	Big Data	700	2020-9-11 16:03:12	1
-- 6	Artificial Intelligence	1660	2018-5-13 18:03:12	3
-- 7	Java P6+	780	2019-1-19 13:03:12	3
-- 8	Data Analysis	500	2019-7-12 13:03:12	1
-- 10	Object Oriented Design	300	2020-8-8 13:03:12	4
-- 12	Dynamic Programming	2000	2018-8-18 20:03:12	1
-- After running your SQL statement, the table should return.

-- DATE_FORMAT
-- 2020-06-01 09:03:12
-- 2020-07-18 10:03:12
-- 2020-02-29 12:03:12
-- 2020-04-22 13:03:12
-- 2020-09-11 16:03:12
-- 2018-05-13 18:03:12
-- 2019-01-19 13:03:12
-- 2019-07-12 13:03:12
-- 2020-08-08 13:03:12
-- 2018-08-18 20:03:12
-- Sample 2.

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	null	4
-- 2	System Design	1350	null	3
-- 3	Django	780	null	3
-- 4	Web	340	null	4
-- 5	Big Data	700	null	1
-- 6	Artificial Intelligence	1660	null	3
-- 7	Java P6+	780	null	3
-- 8	Data Analysis	500	null	1
-- 10	Object Oriented Design	300	null	4
-- 12	Dynamic Programming	2000	null	1
-- After running your SQL statement, the table should return.

-- DATE_FORMAT
-- null
-- null
-- null
-- null
-- null
-- null
-- There is no course creation time in the course list in Example 2, so the statistics are null.

SELECT
	DATE_FORMAT(created_at,'%Y-%m-%d %H:%i:%s') AS DATE_FORMAT
FROM
	courses