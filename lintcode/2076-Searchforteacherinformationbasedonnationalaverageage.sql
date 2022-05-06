-- 2076 · Search for teacher information based on national average age
-- # Description
-- Please write a SQL statement to query the teachers table to filter for countries 
-- where the average age of teachers in that country is greater than the average age of teachers in all countries, 
-- and query for information about teachers in those countries.

-- Table Definition: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The column name returned by the query needs to be the same as the case of the column name output by the sample.
-- If the input data is null, NULL is returned.
-- Example
-- Example 1

-- Table Contents : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 3	Western Venom	western.venom@163.com	28	USA
-- Example 2

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	20	CN
-- 3	Western Venom	western.venom@163.com	20	USA
-- 4	Southern Emperor	southern.emperor@qq.com	20	JP
-- 5	Linghu Chong	linghu.chong@linghuchong.com	20	CN
-- 6	Jia He	jia.he@163.com	20	CN
-- 7	Men Tu	men.tu@guju.com	20	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- There is no eligible data in Example 2, so the output contains only table headers and no data.
SELECT
	*
FROM
	teachers
WHERE
	country IN (
		SELECT
			p.country
		FROM
			( -- 国家,平均年龄
				SELECT
					t.country,
					SUM(t.age) / COUNT(*) AS avg_age
				FROM
					teachers AS t 
				GROUP BY
					t.country
			) AS p
		WHERE
			p.avg_age > ( -- 平均年龄
				SELECT
					SUM(t1.age) / COUNT(*)
				FROM
					teachers AS t1
			)
	)