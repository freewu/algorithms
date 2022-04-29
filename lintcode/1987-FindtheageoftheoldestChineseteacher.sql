-- 1987 · Find the age of the oldest Chinese teacher
-- # Description
-- Write an SQL statement that uses the aggregate function MAX() to query the oldest Chinese teacher from the teachers table and return the age of that teacher, with the column name max_age.
-- Table definition: teachers (Teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality



-- The column names returned by the query need to match the case of the column names in the sample output
-- If the number of students is NULL in the input data, then the data is skipped
-- Returns NULL if all the number of student in the input data are NULL, or if the input data is empty.
-- Example
-- Sample 1:

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- max_age
-- 21
-- Sample 2:

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Western Venom	western.venom@163.com	28	USA
-- 3	Southern Emperor	southern.emperor@qq.com	21	JP
-- After running your SQL statement, the table should return.

-- max_age
-- NULL
-- There is no Chinese teacher in sample 2, so null will be returned.
SELECT
	MAX(age) AS max_age
FROM
	teachers
WHERE
	country = 'CN'