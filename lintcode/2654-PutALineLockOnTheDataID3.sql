-- 2654 · Put a line lock on the data 'id = 3'
-- # Please put a read lock on the row id = 3 in the teachers' table.

-- Example
-- Enter data:

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- Return results:

-- id	name	email	age	country
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'

-- Write your SQL here --
SELECT
	*
FROM
	`teachers`
WHERE
	id = 3
FOR UPDATE;