-- 2093 · Check the portfolio of courses and teachers
-- Description
-- Write a SQL statement that cross join the course table courses 
-- and the teacher table teachers to query the combination of course name and teacher, 
-- with the column named course_name and teacher_name. respectively.

-- Table Definition 1: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Table Definition 2 : teachers

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- This question is a cross join of two tables.
-- If the query of both tables are empty, nothing is returned.
-- Example
-- Example 1:

-- Table Contents 1 : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	5
-- 4	Web	340	2020-04-22	4
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- course_name	teacher_name
-- Web	Eastern Heretic
-- Django	Eastern Heretic
-- System Design	Eastern Heretic
-- Senior Algorithm	Eastern Heretic
-- Web	Northern Beggar
-- Django	Northern Beggar
-- System Design	Northern Beggar
-- Senior Algorithm	Northern Beggar
-- Web	Western Venom
-- Django	Western Venom
-- System Design	Western Venom
-- Senior Algorithm	Western Venom
-- Web	Southern Emperor
-- Django	Southern Emperor
-- System Design	Southern Emperor
-- Senior Algorithm	Southern Emperor
-- Web	Linghu Chong
-- Django	Linghu Chong
-- System Design	Linghu Chong
-- Senior Algorithm	Linghu Chong
-- Example 2:

-- Table Contents 1 : courses

-- id	name	student_count	created_at	teacher_id
-- 1	NULL	880	2020-06-01	4
-- 2	NULL	1350	2020-07-18	3
-- 3	NULL	780	2020-02-29	5
-- 4	NULL	340	2020-04-22	4
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 1	NULL	eastern.heretic@gmail.com	20	UK
-- 2	NULL	northern.beggar@qq.com	21	CN
-- 3	NULL	western.venom@163.com	28	USA
-- 4	NULL	southern.emperor@qq.com	21	JP
-- After running your SQL statement, the table should return.

-- course_name	teacher_name
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- NULL	NULL
-- Because the course name in the course list and the teacher name in the teacher list in Sample 2 are both empty, no matches are returned for the query and null is returned.

SELECT
	c.name AS course_name,
	t.name AS teacher_name
FROM
	teachers t 
CROSS JOIN 
	courses c

-- 交叉连接返回的结果是被连接的两个表中所有数据行的笛卡尔积。
-- 需要注意的是，交叉连接产生的结果是笛卡尔积，并没有实际应用的意义。