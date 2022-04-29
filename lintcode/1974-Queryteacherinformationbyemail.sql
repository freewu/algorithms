-- 1974 · Query teacher information by email
-- Description
-- Write a SQL statement to query the name and email of all teachers who use qq email(mailboxes ending with "@qq.com" ) in the teacher tableteachers.

-- Table definition: teachers (teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality



-- Note that the information returned is the teacher's name and email address
-- If there is no query results, nothing will be returned
-- Example
-- Sample I:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return.

-- name	email
-- Northern Beggar	northern.beggar@qq.com
-- Southern Emperor	southern.emperor@qq.com
-- Sample 1:

-- Table Contents : Teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 3	Western Venom	western.venom@163.com	28	USA
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return.

-- name	email
-- No teachers use qq email, so the result is empty.
SELECT
	name,
	email
FROM
	teachers
WHERE
	email LIKE "%@qq.com"