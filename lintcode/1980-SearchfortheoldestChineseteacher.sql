-- 1980 · Search for the oldest Chinese teacher
-- # Description
-- Write an SQL statement to query the information of the oldest Chinese teacher from the teachers table (teachers).

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality



-- If there is more than one teacher with the oldest age, then return the first data in the table that meets the condition
-- If the query does not return any results, nothing is returned
-- Example
-- Example 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	22	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	USA
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	22	JP
-- 5	Linghu Chong		18	USA
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- There is no eligible data in sample 2, so the output contains only table headers and no data.
SELECT
	*
FROM
	teachers
WHERE
	country = 'CN'
ORDER BY
	age DESC
LIMIT 1