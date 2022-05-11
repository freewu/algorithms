-- 2700 · Update Linghu Chong's age by view
-- Description
-- You are provided with an updatable view v_teachers. 
-- Write SQL statements to change the age of Linghu Chong teacher to 30 years.

-- Table Definition: teachers (Teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	Instructor's name
-- email	varchar	Instructor's email
-- age	int	Tutor's age
-- country	varchar	Tutor's nationality

-- Example
-- Sample 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	30	'CN'
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		26	CN
-- Example 2:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- Because there is no data for Linghu Chong in Example 2, the original table is returned

-- Write your SQL here --
UPDATE 
	`v_teachers`
SET
	age = 30 
WHERE
	name = 'Linghu Chong'