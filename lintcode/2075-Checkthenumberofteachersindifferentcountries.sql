-- 2075 · Check the number of teachers in different countries
-- # Description
-- Write a SQL statement that counts the number of teachers of different nationalities from the teachers table teachers,
-- and sorts the results in ascending order by the number of teachers from different country, 
-- or in ascending order by nationality if the number of teachers were same, with the column named teacher_count.

-- Table definition: teachers

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The query needs to return the same column names as the sample output.

-- Example
-- Sample 1

-- Table content: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- country	teacher_count
-- JP	1
-- UK	1
-- USA	1
-- CN	2
-- Sample 2

-- Table content: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	NULL
-- 2	Northern Beggar	northern.beggar@qq.com	21	NULL
-- 3	Western Venom	western.venom@163.com	28	NULL
-- 4	Southern Emperor	southern.emperor@qq.com	21	NULL
-- 5	Linghu Chong	NULL	18	NULL
-- After running your SQL statement, the table should return.

-- country	teacher_count
-- NULL	0
-- Since the counrtry field is empty for all teachers, country returns null and teacher_count returns 0.
SELECT
	country,
	IF(country IS NULL,0,COUNT(*)) AS teacher_count
FROM
	teachers
GROUP BY
	country
ORDER BY 
	teacher_count ASC, country ASC