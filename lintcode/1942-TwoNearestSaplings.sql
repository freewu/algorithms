-- 1942 · Two Nearest Saplings
-- # Description
-- Some newly planted saplings with varying distances are known to be in a straight line with a big watering bucket.
-- The sapling_distances table stores the distance between some saplings and the watering bucket (distance).
-- Please write a SQL statement to find the distance between the nearest two saplings (shortest_distance).

-- Table definition: sapling_distances

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- distance	int	Distance between seedlings and bucket
-- Example
-- Example 1:

-- Table content: sapling_distances

-- id	distance
-- 1	8
-- 2	1
-- 3	4
-- After running your SQL statement, the table should return:

-- shortest_distance
-- 3
-- Example 2:

-- Table content: sapling_distances

-- id	distance
-- 1	8
-- 2	1
-- 3	4
-- 4	10
-- After running your SQL statement, the table should return:

-- shortest_distance
-- 2

SELECT 
	*
FROM 
	(
		SELECT 
			MIN(c.distance) AS shortest_distance
		FROM 
			(
				SELECT
					ABS(a.distance - b.distance) AS distance
				FROM	
					sapling_distances AS a,
					sapling_distances AS b
				WHERE
					a.id != b.id
			) AS c
	) AS d
WHERE d.shortest_distance IS NOT NULL;  -- 加这个因为 如果数据为空格 min 会返回 NULL 用例跑不过