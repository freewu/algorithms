-- 2080 · Search for the name of the instructor and the total number of students in all the instructor's courses with less than 3000 students
-- # Description
-- Write an SQL statement that join the courses and teachers tables to count the total number of students of courses offered by same teacher, 
-- and count the total number of students as 0 for teachers teach nothing.
-- Finally, query the names of teachers and the total number of students (alias student_count) with total students fewer than 3000. 
-- Sort the results in ascending order by the total number of students or, if the total number of students were the same, by the teacher's name.

-- Table definition 1: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	creation time
-- teacher_id	int	teacher id
-- Table Definition 2: teachers

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The column name returned by the query needs to be the same as the case of the column name output by the sample.
-- If the teacher taught nothing, or if there is a NULL value in the input data, you need to set the total number of students to 0.
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
-- After running your SQL statement, table should return.

-- name	student_count
-- Linghu Chong	0
-- Northern Beggar	0
-- Southern Emperor	1520
-- Example 2:

-- Table Contents 1: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	NULL	2020-06-01	4
-- 6	Artificial Intelligence	NULL	2018-05-13	3
-- 8	Data Analysis	NULL	2019-07-12	1
-- 10	Object Oriented Design	NULL	2020-08-08	4
-- Table Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, table should return.

-- name	student_count
-- Eastern Heretic	0
-- Linghu Chong	0
-- Northern Beggar	0
-- Southern Emperor	0
-- Western Venom	0
-- Since the student_count data in the input exmple are all NULL, the total number of students here is 0.
SELECT
	p.teacher_name AS name,
	p.student_count AS student_count
FROM
	(
		SELECT
			t.name AS teacher_name,
			IFNULL(SUM(c.student_count),0) AS student_count
		FROM
			teachers AS t
		LEFT JOIN
			courses AS c
		ON 
			t.id = c.teacher_id
		GROUP BY
			t.id
	) AS p
WHERE
	p.student_count < 3000

-- use having 
SELECT
    t.name AS name,
    IFNULL(SUM(c.student_count),0) AS student_count
FROM
    teachers AS t
LEFT JOIN
    courses AS c
ON 
    t.id = c.teacher_id
GROUP BY
    t.id
HAVING 
    student_count < 3000