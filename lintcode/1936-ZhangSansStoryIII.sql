-- 1936 · Zhang San's Story III
-- # Description
-- The reporter's investigation found that Zhang San's school also has a gold medal teacher, and his students are well-known institutions after graduation. 
-- The reporter became curious about the situation of the students taught by the teacher.
-- The students table records the name of the student and the class (class_id), the classes table records the name of the class and the teacher (teacher_id), and the teachers table records the name of the teacher
-- Please write SQL statements to find out the names of all the students taught by teacher "xujia".

-- Table definition 1: students

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	student name
-- class_id	int	student's class
-- Table definition 2: classes

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	class name
-- teacher_id	int	teacher id
-- Table definition 3: teachers

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	teacher name
-- Tip:

-- A student has only one class, and a class has only one head teacher
-- Data guarantee that there are no students with the same name
-- A head teacher may bring multiple classes
-- Example
-- Example 1:

-- Table content 1: students

-- id	name	class_id
-- 1	zhangsan	2
-- 2	lisi	1
-- 3	wanger	4
-- 4	zhaoliu	1
-- 5	niuniu	2
-- 6	zhangfei	3
-- 7	guanyu	4
-- 8	liubei	3
-- 9	linqi	2
-- Table content 2: classes

-- id	name	teacher_id
-- 1	Class 3	2
-- 2	Class 4	1
-- 3	Class 5	3
-- 4	Class 6	2
-- Table content 3: teachers

-- id	name
-- 1	lixiuqi
-- 2	xujia
-- 3	xucong
-- After running your SQL statement, the table should return:

-- name
-- lisi
-- wanger
-- zhaoliu
-- guanyu
-- Example 2:

-- Table content 1: students

-- id	name	class_id
-- 1	zhangsan	2
-- 2	lisi	1
-- 3	wanger	4
-- 4	zhaoliu	1
-- 5	niuniu	2
-- 6	zhangfei	3
-- 7	guanyu	4
-- 8	liubei	3
-- 9	linqi	2
-- Table content 2: classes

-- id	name	teacher_id
-- 1	Class 3	2
-- 2	Class 4	1
-- 3	Class 5	3
-- 4	Class 6	4
-- Table content 3: teachers

-- id	name
-- 1	lixiuqi
-- 2	xujia
-- 3	xucong
-- 4	lilei
-- After running your SQL statement, the table should return:

-- name
-- lisi
-- zhaoliu
SELECT
	s.name AS name
FROM 
	students AS s,
	classes AS c,
	teachers  AS t 
WHERE
	s.class_id = c.id AND
	c.teacher_id = t.id AND
	t.name = 'xujia'