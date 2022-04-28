-- 1952 · Query teachers over 20 years old
-- # Description
-- Please write a SQL statement to query the information of all teachers who is older than 20 in the teacher table teachers.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher nationality
-- Example
-- Example 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	18	CN
-- 3	Western Venom	western.venom@163.com	19	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- Because no teacher in the input sample is older than 20, the returned result is empty.
SELECT 
	*
FROM 
	teachers
WHERE
	age > 20