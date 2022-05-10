-- 2565 · Create a Trigger "before_teachers_update"
-- # Description
-- We want to create a trigger for the teachers table and name it before_teachers_update, 
-- which should be executed before adding a new data, the trigger execution should be SET new.country = 'CN', 
-- please write SQL statement to achieve it

-- Table definition : teachers (Teachers table)
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
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'CN'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'CN'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'CN'
-- 5	'Linghu Chong'	None	18	'CN'
-- 31	'CN'	

-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_teachers_update`;
CREATE TRIGGER `before_teachers_update`
BEFORE UPDATE ON `teachers`
FOR EACH ROW
BEGIN
	SET new.country = 'CN';
END
