-- 1961 · Query information about non-Japanese and non-American teachers
-- # Description
-- Please write a SQL statement and use NOT IN to query the information of all teachers whose nationality (country) is not Japan (JP) or the United States (USA) in the teachers table (teachers).

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Tip:

-- The question requires the use of NOT IN
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
-- Because the nationality of the teachers in the input sample is JP or USA, the returned result has no data.
SELECT
	*
FROM
	teachers
WHERE
	country NOT IN ('JP','USA')