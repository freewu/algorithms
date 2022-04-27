-- 1937 · Consecutive Empty Boxes
-- # Description
-- There are some boxes marked with ID in one place, some of which are filled with things and some of which are free.
-- Please write SQL statements, find empty and consecutive boxes, and return them in ascending order of id.

-- Table definition: boxes

-- columns_name	type	explaination
-- id	int	primary key
-- is_empty	int	Box status (0 means the box is occupied, 1 means the box is free and available)
-- Tip:

-- Continuous empty boxes refer to: empty and continuous boxes ≥ 2
-- Data guarantee that the value of is_empty is 0 or 1
-- The id of the boxes are all continuous
-- Example
-- Example 1:

-- Table content: boxes

-- id	is_empty
-- 1	1
-- 2	0
-- 3	1
-- 4	1
-- 5	1
-- After running your SQL statement, the table should return:

-- id
-- 3
-- 4
-- 5
-- Example 2:

-- Table content: boxes

-- id	is_empty
-- 1	1
-- 2	0
-- 3	1
-- 4	1
-- 5	0
-- After running your SQL statement, the table should return:

-- id
-- 3
-- 4
SELECT
	DISTINCT(a.id) AS id
FROM 
	boxes AS a
LEFT JOIN 
	boxes AS b
ON 
	ABS(a.id - b.id) = 1 -- 这个是关键点,相邻两个记录
WHERE
	a.is_empty = 1 AND
	b.is_empty = 1
ORDER BY id