-- 2567 · Delete the Trigger "before_teachers_insert"
-- # Description
-- We want to remove the trigger before_teachers_insert from the teachers table, 
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
-- 1	'Linghu Chong'	None	18	'CN'
-- Return results :

-- id	name	email	age	country
-- 1	'Linghu Chong'	None	18	'CN'
-- 2	'Eastern heretic'	'eastern.heretic@gmail.com'	33	'UK'
-- 3	'Northern Beggar'	'northern.beggar@qq.com'	21	'JP'
-- 4	'Western Venom'	'western.venom@163.com'	28	'CN'
-- 5	'Southern Emperor'	'southern.emperor@qq.com'	31	'USA'

-- MySQL删除触发器：
-- 1、可以使用drop trigger删除触发器：
-- drop trigger trigger_name;
 
-- 2、删除前先判断触发器是否存在：
-- drop trigger if exists trigger_name

-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_teachers_insert`;