-- 1946 · Club Annual Competition Score Ranking II
-- # Description
-- The rankings table records the ranking and score information of a club's annual competition, including item id (category_id), year (year), ranking (rank) and score (score)
-- The categories table records the name of the item (name)
-- Please write SQL statements to query the project name (name) and average score (average_score) of all items in the rankings table and categories table
-- The average score needs to keep two decimal places

-- Table definition 1: rankings

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- category_id	int	category id
-- year	int	year
-- rank	int	rank
-- score	int	score
-- Table definition 2: categories

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	subject name
-- Tip:

-- average_score needs to keep two decimal places
-- Example
-- Example 1:

-- Table content 1: rankings

-- id	category_id	year	rank	score
-- 1	1	2008	15	90
-- 2	1	2012	11	98
-- 3	2	2016	10	99
-- Table content 2: categories

-- id	name
-- 1	volleyball
-- 2	basketball
-- 3	soccer
-- After running your SQL statement, the table should return:

-- name	average_score
-- volleyball	94.00
-- basketball	99.00
-- Example 2:

-- Table content 1: rankings

-- id	category_id	year	rank	score
-- 1	1	2008	15	90
-- 2	1	2012	11	98
-- 3	2	2016	10	99
-- Table content 2: categories

-- id	name
-- 1	volleyball
-- 2	basketball
-- 3	soccer
-- After running your SQL statement, the table should return:

-- name	average_score
-- volleyball	94.00
-- basketball	99.00
-- soccer	91.00

-- use TRUNCATE
SELECT 
	c.name AS name,
	TRUNCATE(SUM(r.score) / COUNT(*),2) AS average_score -- keep two decimal places
FROM 
	categories AS c,
	rankings AS r 
WHERE
	r.category_id = c.id
GROUP BY c.name

-- use CONVERT
SELECT 
	c.name AS name,
	CONVERT((SUM(r.score) / COUNT(*)), decimal(10,2)) AS average_score -- keep two decimal places
FROM 
	categories AS c,
	rankings AS r 
WHERE
	r.category_id = c.id
GROUP BY c.name