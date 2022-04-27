-- 1921 · Players Who Never Recharge
-- # Description
-- A game database contains two tables, users table and recharges table.
-- Write a SQL query to find all players who never recharge.

-- Table definition 1: users

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	username
-- Table definition 2: recharges

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- user_id	int	user id
-- Tip:

-- You can solve this problem by learning SQL joint query related knowledge
-- Returned column name: player
-- Example
-- Example 1:

-- Table content 1: users

-- id	name
-- 1	XiaoXuesheng
-- 2	Mike
-- 3	John
-- 4	Maria
-- Table content 2: recharges

-- id	user_id
-- 1	3
-- 2	1
-- For example, given the above table, your query should return:

-- player
-- Mike
-- Maria
-- Example 2:

-- Table content 1: users

-- id	name
-- 1	XiaoXuesheng
-- 2	Mike
-- 3	John
-- Table content 2: recharges

-- id	user_id
-- 1	3
-- 2	1
-- For example, given the above table, your query should return:

-- player
-- Mike
SELECT 
	name AS player
FROM 
	users
WHERE 
	id NOT IN (
		SELECT	
			user_id
		FROM
			recharges
	)
