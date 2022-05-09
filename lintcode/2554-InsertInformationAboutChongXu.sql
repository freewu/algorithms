-- 2554 · Insert information about Chong Xu
-- # Description
-- We want to insert the information of Chong Xu with email chong.xu@wudang.com, age 88 and nationality CN into the teachers table, 
-- however, the courses table is write-locked, write an SQL statement to implement the insertion of Chong Xu.

-- Table Definition 1: teachers (Teachers table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Tutor's email
-- age	int	lecturer's age
-- country	varchar	Tutor's nationality
-- Table Definition 2: courses (Course List)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course_name
-- student_count	int	total_students
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor id
-- Translated with www.DeepL.com/Translator (free version)

-- Please note that the teachers table is write-locked

-- Example
-- Sample 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@outlook.com	21	JP
-- 5	Linghu Chong		18	CN
-- 6	Chong Xu	chong.xu@wudang.com	88	CN
-- Sample 2:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Linghu Chong		18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Linghu Chong		18	CN
-- 5	Chong Xu	chong.xu@wudang.com	88	CN

-- 对 courses 表上写锁，不要删除该代码 --
LOCK TABLES courses WRITE;

-- Write your SQL Query here --
UNLOCK TABLES;

INSERT INTO 
	`teachers` (
		`name`,
		`email`,
		`age`,
		`country`
	)
VALUES (
	'Chong Xu',
	'chong.xu@wudang.com',
	88,
	'CN'
)