-- 1945 · Club Annual Competition Score Ranking I
-- # Description
-- The rankings table records the ranking and score information of a club’s annual competition, including item id (category_id), year (year), ranking (rank) and score (score)
-- The categories table records the name of the project (name). For some reasons, there may be data loss in the categories table, as in sample two.
-- Please write a SQL statement to query the name of the item (name), the year of the competition (year) and the score (score) of all items in the rankings table and categories table

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

-- name	year	score
-- volleyball	2008	90
-- volleyball	2012	98
-- basketball	2016	99
-- Example 2:

-- Table content 1: rankings

-- id	category_id	year	rank	score
-- 1	1	2008	15	90
-- 2	3	2012	11	98
-- 3	2	2016	10	99
-- 4	4	2017	40	72
-- Table content 2: categories

-- id	name
-- 1	volleyball
-- 2	basketball
-- 3	soccer
-- After running your SQL statement, the table should return:

-- name	year	score
-- volleyball	2008	90
-- soccer	2012	98
-- basketball	2016	99
SELECT 
	c.name AS name,
	r.year AS year,
	r.score AS score
FROM 
	categories AS c,
	rankings AS r 
WHERE
	r.category_id = c.id