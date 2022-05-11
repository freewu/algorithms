-- 2716 · View Processing Algorithm TEMPTABLE
-- # Description
-- Please create a view v_teachers to view the information of teachers whose nationality is CN in the teachers table and use the view algorithm of TEMPTABLE

-- Table Definition : teachers (Teachers Table)
-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	Tutor's nationality

-- Example
-- Enter data:

-- teachers table.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'CN'
-- Results returned:

-- id	name	email	age	country
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'CN'

-- Write your SQL here --
CREATE OR REPLACE ALGORITHM = TEMPTABLE VIEW
	`v_teachers`
AS 
	SELECT
		*
	FROM
		`teachers`
	WHERE
		country = 'CN';