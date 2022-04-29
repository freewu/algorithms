-- 1998 · Statistics for courses in January to May 2020
-- # Description
-- Please write an SQL statement to count the number of courses started from January to May in 2020 in the course table courses , 
-- and finally return the statistical value with the result column named course_count.

-- Table definition: courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course start time
-- teacher_id	int	teacher id
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- The restricted dates of the query include January and May, and the courses could be duplicated
-- If there is a null value in the data, it will not be counted in the statistics
-- If no data is counted, 0 is returned
-- Example
-- Example 1

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
-- After running your SQL statement, the table should return:

-- course_count
-- 2
-- Example 2

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2019-02-29	3
-- 4	Web	340	2021-04-22	4
-- After running your SQL statement, the table should return:

-- course_count
-- 0
-- There are no eligible data in sample 2, so the statistic result is 0
SELECT
	COUNT(*) AS course_count
FROM
	courses
WHERE
	created_at BETWEEN '2020-01-01' AND '2020-05-31'