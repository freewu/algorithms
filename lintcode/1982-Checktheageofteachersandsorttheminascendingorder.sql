-- 1982 · Check the age of teachers and sort them in ascending order
-- # Description
-- Write an SQL statement that queries the teachers table teachers for the unique value of teacher age age and sorts the results in ascending order by age age.
-- Table definition: teachers (teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality



-- Note that the returned age cannot be duplicated
-- If the query does not return results, return null
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

-- age
-- 18
-- 20
-- 21
-- 28
-- Sample 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com		UK
-- 2	Northern Beggar	northern.beggar@qq.com		CN
-- 3	Western Venom	western.venom@163.com		USA
-- 4	Southern Emperor	southern.emperor@qq.com		JP
-- 5	Linghu Chong			CN
-- After running your SQL statement, the table should return.

-- age
-- Because the age information is empty, the return result is also empty.
SELECT
	distinct(age) AS age
FROM
	teachers 
ORDER BY
	age ASC