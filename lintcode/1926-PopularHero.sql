-- 1926 · Popular Hero
-- # Description
-- Write a SQL query to find all heroes whose popularity is not T3 (not popular) and whose id is odd. 
-- Please sort the results in descending order of ban rate.
-- Tx (x = 0, 1 ,2, ...) represents the popularity of heroes,
-- among which T0 is a popular hero, T1 is a sub-popular hero, T2 is a normal popular hero, T3 is a non-popular hero, and T4 and beyond are unpopular heroes .

-- Table definition: heroes

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	hero name
-- popularity	varchar	hero popularity
-- ban	float	hero banned probability
-- Note, the rate returned probability need to bring the percent sign (%)

-- Example
-- Example 1:

-- Table content: heroes

-- id	name	popularity	ban
-- 1	Lv Bu	T0	0.90
-- 2	Ju Youjing	T1	0.24
-- 3	Ma Chao	T1	10.26
-- 4	Guan Yu	T2	1.56
-- 5	Meng Tian	T3	2.10
-- For the above example, the correct output is:

-- id	name	popularity	probability
-- 3	Ma Chao	T1	10.26%
-- 1	Lv Bu	T0	0.90%
-- Example 2:

-- Table content: heroes

-- id	name	popularity	ban
-- 1	Lv Bu	T3	0.90
-- 2	Ju Youjing	T1	0.24
-- 3	Ma Chao	T1	10.26
-- 4	Guan Yu	T2	1.56
-- For the above example, the correct output is:

-- id	name	popularity	probability
-- 1	Ma Chao	T1	10.26%
SELECT 
	id,
	name,
	popularity,
	CONCAT(ban,"%") AS probability
FROM	
	heroes
WHERE
	popularity != 'T3' AND
	id % 2 = 1
ORDER BY ban DESC