-- 2027 · Query course creation date by 'year month'
-- Description
-- Write a SQL statement that queries the course creation time in the course tablecourses, and returns the result in the format ‘Year Month' with the column named DATE_FORMAT.

-- Table Definition: courses

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student count	int	student count
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- The column names returned by the query need to match the case of the column names in the sample output

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

-- DATE_FORMAT
-- 2020 09
-- 2020 08
-- 2020 07
-- 2020 06
-- 2018 08
-- 2018 05
-- Example 2.

-- Table Contents: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	null	4
-- 2	System Design	1350	null	3
-- After running your SQL statement, the table should return.

-- DATE_FORMAT
-- null
-- null
-- The course table in Example 2 does not have a course creation time, so the statistics are empty.

SELECT
	DATE_FORMAT(created_at,'%Y %m') AS DATE_FORMAT
FROM
	courses