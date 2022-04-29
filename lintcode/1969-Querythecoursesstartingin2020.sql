-- 1969 · Query the courses starting in 2020
-- # Description
-- Please write SQL statements to query the information of all courses which were opened in 2020 in the course tablecourses.

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- Within 2020 refers to the period from '2020-01-01' to '2020-12-31', and both days include
-- If there is no query result, nothing will be returned
-- Example
-- Example 1:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Big Data	700	2020-09-11	1
-- 5	Artificial Intelligence	1660	2018-05-13	3
-- 6	Java P6+	780	2019-01-19	3
-- 7	Data Analysis	500	2019-07-12	1
-- 8	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Big Data	700	2020-09-11	1
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Web	340	2021-04-22	4
-- 2	Artificial Intelligence	1060	2018-05-13	3
-- 3	Java P6+	780	2019-01-19	3
-- 4	Data Analysis	500	2019-07-12	1
-- 5	Object Oriented Design	300	2019-08-08	4
-- 6	Dynamic Programming	1000	2018-08-18	1
-- After running your SQL statement, the table should return:

-- id	name	student_count	created_at	teacher_id
-- Because there is no data to be queried in the data, only the title is displayed here, no data.
SELECT
	*
FROM
	courses
WHERE
	created_at >= '2020-01-01' AND
	created_at <= '2020-12-31'