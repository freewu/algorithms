-- 1935 · Zhang San's Story II
-- # Description
-- The reporter found that Zhang San’s good grades were inseparable from the help of his classmates and teachers, and decided to interview Zhang San’s class
-- The students table records the name of the student and the class (class_id), and the classes table records the name of the class and the teacher (teacher_id)
-- Please write SQL statements to find out the class name of the class where "zhangsan" is.

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
-- Tip:

-- A student has only one class, and a class has only one head teacher
-- Data guarantee that there are no students with the same name
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
-- After running your SQL statement, the table should return:

-- name
-- Class 4
-- Example 2:

-- Table content 1: students

-- id	name	class_id
-- 1	zhangsan	2
-- 2	lisi	1
-- 3	wanger	4
-- Table content 2: classes

-- id	name	teacher_id
-- 1	Class 3	2
-- 2	Class 4	1
-- 3	Class 5	3
-- 4	Class 6	2
-- After running your SQL statement, the table should return:

-- name
-- Class 4
SELECT
	c.name
FROM 
	classes AS c,
	students AS s
WHERE
	c.id = s.class_id AND 
	s.name = 'zhangsan'