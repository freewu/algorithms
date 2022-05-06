-- 2057 · Modify course information created by instructor Eastern Heretic
-- # Description
-- Write a SQL statement to query the information of Eastern Heretic from the teachers table,
-- and change the course name to PHP and student count to 300 of courses taught by Eastern Heretic according to teacher id.
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

-- The column name returned by the query needs to be the same case as the column name outputed by the sample.
-- If there is a teacher id that is null, then the teacher id will also return null.
-- Example
-- Example I:

-- Table Contents 1: courses

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
-- Table of Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, we will execute the following statement.

-- SELECT *
-- FROM `courses`;
-- Returns.

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01T00:00:00	4
-- 2	System Design	1350	2020-07-18T00:00:00	3
-- 3	Django	780	2020-02-29T00:00:00	3
-- 4	Web	340	2020-04-22T00:00:00	4
-- 5	PHP	300	2020-09-11T00:00:00	1
-- 6	Artificial Intelligence	1660	2018-05-13T00:00:00	3
-- 7	Java P6+	780	2019-01-19T00:00:00	3
-- 8	PHP	300	2019-07-12T00:00:00	1
-- 10	Object Oriented Design	300	2020-08-08T00:00:00	4
-- 12	PHP	300	2018-08-18T00:00:00	1
-- Example 2:

-- Table Contents 1: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	Java	1350	2020-07-18	3
-- 3	Java	780	2020-02-29	3
-- Table of Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, we will execute the following statement.

-- SELECT *
-- FROM `courses`;
-- Returns.

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01T00:00:00	4
-- 2	Java	1350	2020-07-18T00:00:00	3
-- 3	Java	780	2020-02-29T00:00:00	3

UPDATE
	courses
SET
	student_count = 300,
	name = 'PHP'
WHERE
	teacher_id IN (
		SELECT
			id
		FROM
			teachers
		WHERE
			name = 'Eastern Heretic'
	)