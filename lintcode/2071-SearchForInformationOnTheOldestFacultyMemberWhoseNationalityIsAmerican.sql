-- 2071 · Search for information on the oldest faculty member whose nationality is American
-- # Description
-- Write an SQL statement to query the information of the oldest teacher who come from USA by using an inline view.

-- Table Definition: teachers

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The query returns column names that need to match the case of the column names in the sample output.
-- There may be more than one teacher who is the oldest and come from USA.
-- If the input data of teacher's age or nationality is NULL, the data will be skipped.
-- If the query returns no results, nothing is returned.
-- Example
-- Example I:

-- Table content: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 3	Western Venom	western.venom@163.com	28	USA
-- Example 2:

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- Because there is no data that matches the criteria in the input sample, only the title is shown here, no data.

-- Write your SQL Query here --
SELECT
	t.*
FROM
	teachers AS t,
	(
		SELECT
			MAX(age) AS age,
			country
		FROM
			teachers
		WHERE
			country = 'USA'
		GROUP BY
			country
	) AS p
WHERE
	t.country = p.country AND
	t.age = p.age