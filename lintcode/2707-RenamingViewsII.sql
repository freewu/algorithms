-- 2707 · Renaming views (II)
-- # Description
-- Now we have a view v_teachers_test that needs to be renamed to v_teachers, write the SQL statement to do so

-- Example
-- Enter data:

-- teachers table.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- The result should be returned after executing your SQL statement:

-- name	country
-- 'Eastern heretic'	'UK'
-- 'Northern Beggar'	'CN'
-- 'Western Venom'	'USA'
-- 'Southern Emperor'	'JP'
-- 'Linghu Chong'	'CN'

-- Write your SQL here --
DROP VIEW `v_teachers_test`;

CREATE OR REPLACE VIEW
	`v_teachers`
AS 
	SELECT 
		name,
		country
	FROM
		`teachers`;