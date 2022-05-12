-- 2721 · Creating a view that ensures consistency (II)
-- # Description
-- Now you need to create an updatable view v_CN_teachers 
-- that only allows viewing and inserting information about teachers whose nationality is China 'CN', 
-- write SQL statement to achieve it.

-- Table Definition : teachers (Teachers table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality

-- View Definition: v_teachers(View)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	Tutor's nationality

-- Example
-- Input：

-- teachers：

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- Return：

-- id	name	email	age	country
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'

-- Write your SQL here --
CREATE OR REPLACE VIEW 
	`v_CN_teachers`
AS 
	SELECT 
		* 
	FROM
		`teachers`
	WHERE
		country = 'CN'
WITH CHECK OPTION;
