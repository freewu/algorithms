-- 1966 · Use NOT BETWEEN to search for teachers with Chinese nationality
-- # Description
-- Write a SQL statement to select the information of all teachers from teachers where country = 'CN' and 'ID' not between 5 and 10.

-- Table definition: teachers (Teachers table)

-- Column Name	Type	Comments
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- The question requires the use of NOT BETWEEN AND, please do not use AND
-- Note that 5 and 10 are not included in the query result
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

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- Sample 2:

-- Table Contents : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	18	USA
-- 2	Northern Beggar	northern.beggar@qq.com	21	JP
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	UK
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- The result is no data because the teachers in the input sample are either JP or USA.

SELECT
	*
FROM
	teachers
WHERE
	country = 'CN' AND 
	id NOT BETWEEN 5 AND 10
