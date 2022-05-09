-- 2568 · Delete the Trigger "before_teachers_update"
-- # Description
-- We want to remove the trigger before_teachers_update from the teachers table, 
-- please write the SQL statement to do so

-- table definition : teachers (teachers table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	Tutor's nationality

-- Example
-- Input data :

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- Return results :

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'RS'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'UK'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'CN'
-- 5	'Linghu Chong'	None	18	'CN'


-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_teachers_update`;