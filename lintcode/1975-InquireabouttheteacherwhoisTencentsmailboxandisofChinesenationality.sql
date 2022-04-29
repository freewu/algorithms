-- 1975 · Inquire about the teacher who is Tencent's mailbox and is of Chinese nationality
-- # Description
-- Write a SQL statement that using LIKE to query all teachers who are Chinese and use Tencent email (mailboxes ending with "@qq.com") in the teachers tableteachers.

-- The teachers table is defined as follows：

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality



-- You need to use fuzzy query to find Tencent mailboxes (mailboxes ending with "@qq.com")
-- When the mailbox is empty, it means the mailbox does not exist and does not belong to Tencent mailbox
-- If there is no query results, nothing will be returned
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
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- Eample 2 ：

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	UK
-- 3	Linghu Chong		18	CN
-- After running your SQL statement, the table should return :

-- id	name	email	age	country
-- Since there is no data that meets the query criteria in Sample 2, an empty table is returned and only the headers are shown here.
SELECT
	*
FROM
	teachers
WHERE
	country = 'CN' AND
	email LIKE '%@qq.com'