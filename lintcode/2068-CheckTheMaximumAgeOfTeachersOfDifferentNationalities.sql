-- 2068 · Check the maximum age of teachers of different nationalities
-- # Description
-- Write an SQL statement to query the name name, age age, and nationality country of the teacher 
-- with the oldest age age of every country country in the teacher table teachers.

-- Table Definition: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The column name returned by the query needs to be the same as the case of the column name output by the sample.
-- If the input data is null, NULL is returned.
-- If a country has more than one teacher with the oldest age, then return all of these teachers.
-- Example
-- Example 1

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- name	age	country
-- Eastern heretic	20	UK
-- Northern Beggar	21	CN
-- Western Venom	28	USA
-- Southern Emperor	21	JP
-- Example 2

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	linghu.chong@linghuchong.com	18	CN
-- 6	Jia He	jia.he@163.com	26	CN
-- 7	Men Tu	men.tu@guju.com	26	CN
-- After running your SQL statement, the table should return.

-- name	age	country
-- Eastern heretic	20	UK
-- Western Venom	28	USA
-- Southern Emperor	21	JP
-- Jia He	26	CN
-- Men Tu	26	CN
-- Example 2 Chinese nationality is CN Two of the teachers are equal and oldest, so the information about both of them is returned

-- Write your SQL Query here --
SELECT
	t.name AS name,
	t.age AS age,
	t.country AS country
FROM
	teachers AS t,
	(
		SELECT
			MAX(age) AS age,
			country
		FROM
			teachers
		GROUP BY
			country
	) AS p
WHERE
	t.country = p.country AND
	t.age = p.age