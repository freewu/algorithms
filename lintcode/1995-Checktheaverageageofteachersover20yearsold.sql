-- 1995 · Check the average age of teachers over 20 years old
-- # Description
-- Write an SQL statement that queries the average age of teachers over 20 years old in the teachers table teachers,
-- returning the field avg_teacher_age, with the result rounded to the nearest integer.

-- Table definition: teachers

-- column name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality



-- The column names returned by the query need to be the same as the case of the column names in the sample output
-- If the teacher's age is NULL in the input data, the data will be skipped.
-- If all the teachers' ages in the input data are NULL, or the input data is empty, then return NULL
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

-- avg_teacher_age
-- 23
-- Sample 2:

-- Table Contents : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	NULL	UK
-- 2	Northern Beggar	northern.beggar@qq.com	NULL	CN
-- 3	Western Venom	western.venom@163.com	NULL	USA
-- 4	Southern Emperor	southern.emperor@qq.com	NULL	JP
-- 5	Linghu Chong		NULL	CN
-- After running your SQL statement, the table should return.

-- avg_teacher_age
-- NULL
-- Example 2 teacher age data is empty, so the result will also be NULL
SELECT
	ROUND(SUM(age) / COUNT(*)) AS avg_teacher_age
FROM
	teachers
WHERE
	age > 20 AND
	age IS NOT NULL