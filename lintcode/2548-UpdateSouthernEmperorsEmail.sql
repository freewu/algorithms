-- 2548 · Update Southern Emperor's email
-- # Description
-- We want to update the mailbox of Southern Emperor in the teachers table to southern.emperor@outlook.com, 
-- however, the courses table is under a write lock, 
-- so write an SQL statement to update the mailbox of Southern Emperor in the teachers table. mailboxes in theteachers` table.

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
-- Example 2:

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

-- 对 courses 表上写锁，不要删除该代码 --
LOCK TABLES courses WRITE;

-- Write your SQL Query here --
UNLOCK TABLES;

UPDATE
	teachers 
SET
	email = 'southern.emperor@outlook.com'
WHERE
	name = 'Southern Emperor'