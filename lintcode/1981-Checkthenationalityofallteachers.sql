-- 1981 · Check the nationality of all teachers
-- # Description
-- Please write an SQL statement to query all the unique teacher's nationality (country) in the teacher table teacher.

-- Table definition: teachers (teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	iteacher's age
-- country	varchar	teacher's nationality



-- Please remember the returned nationality should not be duplicated
-- If the query does not return results, it should return null
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

-- country
-- UK
-- CN
-- USA
-- JP
-- Sample 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	
-- 2	Northern Beggar	northern.beggar@qq.com	21	
-- 3	Western Venom	western.venom@163.com	28	
-- 4	Southern Emperor	southern.emperor@qq.com	21	
-- 5	Linghu Chong		18	
-- After running your SQL statement, the table should return.

-- country
-- The result is null because the nationality information is null.

-- use group by 
SELECT
	country
FROM
	teachers
GROUP BY
	country

-- use distinct
SELECT
	distinct(country) AS country
FROM
	teachers