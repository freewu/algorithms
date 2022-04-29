-- 1989 · Check the age of the youngest teacher
-- # Description
-- Write an SQL statement to query the minimum age (age) of teachers in the teachers table, and named output field min_age.

-- Table definition: teachers (teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- If the query does not return any results, return None.

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

-- min_age
-- 18
-- Sample 2:

-- Table Contents : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 3	Western Venom	western.venom@163.com	33	USA
-- 5	Linghu Chong		20	CN
-- 6	Southern Emperor	southern.emperor@qq.com	21	JP
-- After running your SQL statement, the table should return.

-- min_age
-- 20

SELECT
	MIN(age) AS min_age
FROM
	teachers