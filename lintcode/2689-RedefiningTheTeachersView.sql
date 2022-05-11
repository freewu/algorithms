-- 2689 · Redefining the teachers view
-- # Description
-- There is already a view v_teachers, we need to redefine this view to view the teachers whose nationality is CN in the teachers table.

-- Table definition: teachers (Teachers table)

-- column name	type	comments
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
-- Enter data:

-- teachers table.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- Returned results:

-- id	name	email	age	country
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'

-- Write your SQL here --
CREATE OR REPLACE VIEW
	`v_teachers`
AS
	SELECT
		*
	FROM
		`teachers`
	WHERE
		country = 'CN';