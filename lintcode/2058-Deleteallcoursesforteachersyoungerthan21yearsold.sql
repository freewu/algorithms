-- 2058 · Delete all courses for teachers younger than 21 years old
-- # Description
-- Write an SQL statement to delete courses in the course table courses where the teacher's age age is less than 21 years old.

-- Table Definition 1: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table Definition 2: courses (Course Table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int unsigned	teacher id

-- The column name returned by the query needs to be the same as the case of the column name output by the sample.
-- If the input data is null, NULL is returned.
-- Example
-- Example I

-- Table Contents 1: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	53	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	20	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents 2: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Java P6+	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, we will execute SELECT * FROM courses and the table should return.

-- id	name	student_count	created_at	teacher_id
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 6	Java P6+	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- Example II

-- Table Contents 1: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	18	CN
-- 3	Western Venom	western.venom@163.com	19	USA
-- 4	Southern Emperor	southern.emperor@qq.com	17	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents 2: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Java P6+	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, we will execute SELECT * FROM courses and the table should return.

-- id	name	student_count	created_at	teacher_id
-- There is no eligible data in example 2, so the output contains only table headers and no data.

DELETE FROM
	courses
WHERE
	teacher_id IN (
		SELECT
			id 
		FROM
			teachers
		WHERE
			age < 21
	)