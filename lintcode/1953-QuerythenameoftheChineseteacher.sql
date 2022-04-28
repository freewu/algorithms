-- 1953 · Query the name of the Chinese teacher
-- # Description
-- Please write a SQL statement to query the teacher table teachers for all teachers whose nationality country is China (CN), 
-- and return the names of all teachers who meet the query conditions

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Tip:

-- Nationality is represented by capital letters in the teachers table
-- If there is no query result, nothing will be returned

-- Example 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return:

-- name
-- Northern Beggar
-- Linghu Chong
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	JP
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	USA
-- After running your SQL statement, the table should return:

-- name
-- Because there is no Chinese teacher in the data, only the title is shown here, no data
SELECT
	name
FROM
	teachers
WHERE
	country = 'CN'