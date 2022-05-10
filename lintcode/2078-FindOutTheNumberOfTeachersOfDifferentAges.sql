-- 2078 · Find out the number of teachers of different ages
-- # Description
-- Write an SQL statement that queries the teachers table for the number of teachers of different ages 
-- and sorts the results in descending order of age, returning the column name as age_count.

-- Table definition: teachers (Teachers table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The query needs to return the same column names as the sample output.
-- If all ages are NULL, then age is returned as NULL and number is returned as 0.
-- Example
-- Sample I

-- Table content: teachers (Teachers table)

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- age	age_count
-- 28	1
-- 21	2
-- 20	1
-- 18	1
-- Sample II

-- Table content: teachers (Teachers table)

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	NULL	UK
-- 2	Northern Beggar	northern.beggar@qq.com	NULL	CN
-- 3	Western Venom	western.venom@163.com	NULL	USA
-- 4	Southern Emperor	southern.emperor@qq.com	NULL	JP
-- 5	Linghu Chong	NULL	NULL	CN
-- After running your SQL statement, the table should return.

-- age	age_count
-- NULL	0
-- Because the age field is empty for all teachers, age returns null and age_count returns 0.

-- Write your SQL Query here --
SELECT
	age,
	IF(age IS NULL,0,COUNT(*)) AS age_count
FROM
	teachers
GROUP BY
	age
ORDER BY 
	age DESC