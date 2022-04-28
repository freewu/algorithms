-- 1959 · Query information about Chinese and British teachers
-- # Description
-- Please write a SQL statement that use IN to query the information of all teachers who are Chinese or British in the teacher table teachers.
-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Tip:

-- The question requires the use of IN, please do not use AND
-- If there is no query result, nothing will be returned
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
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 5	Linghu Chong		18	CN
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	18	USA
-- 2	Northern Beggar	northern.beggar@qq.com	21	JP
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	USA
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- Because no teacher’s nationality is China (CN) or United Kingdom (UK) in the input sample, the returned result is empty.
SELECT 
	*
FROM
	teachers
WHERE
	country IN ('CN','UK')