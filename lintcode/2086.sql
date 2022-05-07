-- 2086 · Search for the nationality of the teacher starting with 'U' and the total number of students between 2000 and 5000 and the total number of students of that nationality
-- # Description
-- Write an SQL statement to count the total number of students of courses offered by teachers from different country by linking the courses and teachers tables, 
-- and count the total number of students as 0 for teachers teach nothing.

-- Finally, query the nationality of teachers whose nationality begins with 'U' 
-- and the total number of students (alias student_count) of all teachers of this nationality is between 2000 and 5000 (Include 2000 and 5000), 
-- then sort the results in descending order by the total number of students, or in ascending order by the teacher's nationality if the total number of students is the same.

-- Table Definition 1: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	creation time
-- teacher_id	int	instructor id
-- Table Definition 2: teachers

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The column name returned by the query needs to be the same as the case of the column name output by the sample.
-- If there is a NULL value for the total number of students in the input data, you need to set the total number of students to 0.
-- If the query does not return any results, nothing will be returned.
-- Example
-- Example I:

-- Table Contents 1: courses

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
-- Table of Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- country	student_count
-- USA	4570
-- UK	3200
-- Example 2:

-- Table Contents 1: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	NULL	2020-06-01	1
-- 2	System Design	NULL	2020-07-18	2
-- 3	Django	NULL	2020-02-29	3
-- 4	Web	NULL	2020-04-22	4
-- Table Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- country	student_count
-- Because the data of student_count in the input sample are all NULL, 
-- the total number of students is counted as 0, but the filtering condition is not satisfied, so only the title is shown here, no data.

SELECT 
	*
FROM 
	(
		SELECT
			t.country AS country,
			SUM(c.student_count) AS student_count
		FROM
			teachers AS t,
			courses AS c
		WHERE
			t.id = c.teacher_id AND
			c.student_count IS NOT NULL AND
			t.country LIKE 'U%'
		GROUP BY
			t.country
	) AS p
WHERE
	student_count BETWEEN 2000 AND 5000
ORDER BY 
	student_count DESC, country ASC