-- 1977 · Sorted by age of Chinese teachers in descending order
-- # Description
-- Write an SQL statement to query the Chinese teachers in the teachers table teachers and sort them by age in descending order.

-- Table definition: teachers (teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher‘s name
-- email	varchar	teacher‘s email
-- age	int	teacher‘s age
-- country	varchar	teacher‘s nationality



-- you must first use WHERE to filter out the appropriate conditions, and then use ORDER BY to sort, otherwise the returned results may be wrong.
-- If the query does not return any results, nothing will be returned.
-- Example
-- Sample 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	22	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 5	Linghu Chong		18	CN
-- Sample 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- The result set has no data but column names because there are no teachers of Chinese nationality in the input sample.
SELECT
	*
FROM
	teachers
WHERE	
	country = 'CN'
ORDER BY
	age DESC