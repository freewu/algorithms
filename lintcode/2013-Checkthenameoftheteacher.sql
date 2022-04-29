-- 2013 · Check the name of the teacher
-- # Description
-- Write an SQL statement to query the names of all teachers in the teachers table teachers.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality


-- If the query does not return any results, nothing is returned.

-- Example
-- Example 1:

-- Table content: teachers

-- | id | name | email | age | country |
-- | ---- | ---- | ---- | ---- | ---- | ---- | ----
-- | 1 | Eastern Heretic | eastern.heretic@gmail.com | 20 | UK |
-- | 2 | Northern Beggar | northern.beggar@qq.com | 21 | CN |
-- | 3 | Western Venom | western.venom@163.com | 28 | USA |
-- | 4 | Southern Emperor | southern.emperor@qq.com | 21 | JP |
-- | 5 | Linghu Chong | | 18 | CN |

-- After running your SQL statement, the table should return:

-- name
-- Eastern Heretic
-- Northern Beggar
-- Western Venom
-- Southern Emperor
-- Linghu Chong
-- Example 2:

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	USA
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	USA
-- 6	Zhang Fe	zhang.fe@gmail.com	20	UK
-- 7	Guan Yu	guan.yu@qq.com	21	USA
-- 8	Liu Bei	liu.bei@163.com	28	USA
-- After running your SQL statement, the table should return:

-- name
-- Eastern Heretic
-- Northern Beggar
-- Western Venom
-- Southern Emperor
-- Linghu Chong
-- Zhang Fe
-- Guan Yu
-- Liu Bei

SELECT
	name
FROM
	teachers