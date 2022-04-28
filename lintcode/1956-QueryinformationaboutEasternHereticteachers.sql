-- 1956 · Query information about Eastern Heretic teachers
-- # Description
-- Please write a SQL statement to query the information of teachers whose name is Eastern Heretic and age is greater than 18 from the teachers table

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Over 18 years old, excluding 18
-- There may be teachers with the same name in the data
-- If there is no query result, nothing will be returned
-- Example
-- Example 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	CN
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	18	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- 6	Jia He	jia.he@163.com	22	CN
-- 7	Men Tu	men.tu@guju.com	26	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- Because there is no data that meets the conditions in the input sample, only the title is shown here, and there is no data.
SELECT
	*
FROM
	teachers
WHERE
	name = 'Eastern Heretic' AND
	age > 18