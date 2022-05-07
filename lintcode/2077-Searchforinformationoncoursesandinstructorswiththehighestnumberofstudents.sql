-- 2077 · Search for information on courses and instructors with the highest number of students
-- # Description
-- Write an SQL statement to join the teachers table and the courses table by using an inline view(INNER JOIN), 
-- and use the name of the course (alias course_name), the total number of students, 
-- and the name of the instructor (alias teacher_name) as the query table, 
-- and select the information with the highest number of students from this query table.

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

-- The query return column name needs to be the same case as the sample output column name.
-- There may be more than one course with the largest number of students .
-- If the information about the total number of students present in the input data is NULL, the data will be skipped.
-- If the query does not return any results, nothing is returned.
-- Example
-- 样例一：

-- 表内容 1：courses

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
-- 表内容 2：teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- 在运行你的 SQL 语句之后，表应返回：

-- course_name	student_count	teacher_name
-- Dynamic Programming	2000	Eastern Heretic
-- 样例二：

-- 表内容 1：courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	NULL	2020-06-01	4
-- 6	Artificial Intelligence	NULL	2018-05-13	3
-- 8	Data Analysis	NULL	2019-07-12	1
-- 10	Object Oriented Design	NULL	2020-08-08	4
-- 表内容 2：teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- 在运行你的 SQL 语句之后，表应返回：

-- course_name	student_count	teacher_name
-- 因为输入样例中 student_count 的数据都是 NULL，所以这里只展示了标题，没有数据。

SELECT
	c.name AS course_name,
	c.student_count AS student_count,
	t.name AS teacher_name
FROM
	courses AS c
INNER JOIN
	teachers AS t
ON 
	c.teacher_id = t.id
WHERE
	c.student_count = (
		SELECT
			MAX(student_count)
		FROM
			courses
	)