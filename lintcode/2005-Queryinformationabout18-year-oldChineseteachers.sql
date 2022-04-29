-- 2005 · Query information about 18-year-old Chinese teachers
-- # Description
-- Please write a SQL statement to find out the information of Chinese teachers whose age is 18 from the teachers table.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality


-- If there is no query result, nothing will be returned.

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
-- 5	Linghu Chong		18	CN
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	18	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	USA
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- Because there is no data that meets the conditions in the input sample, only the title is shown here, and there is no data.
SELECT
	*
FROM
	teachers
WHERE
	age = 18 AND
	country = 'CN'