-- 1944 · Bike Sharing
-- # Description
-- The shared_bicycles table stores the usage information of shared bicycles, including bicycle id (bike_id) and user id (user_id)
-- Write a SQL statement to find the shared bicycle id and user id used by the same person at least three times

-- Table definition: shared_bicycles

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- bike_id	int	bike id
-- user_id	int	user id
-- Example
-- Example 1:

-- Table content: shared_bicycles

-- id	bike_id	user_id
-- 1	1	1
-- 2	1	1
-- 3	1	1
-- 4	1	2
-- 5	1	2
-- 6	2	1
-- 7	2	1
-- After running your SQL statement, the table should return:

-- bike_id	user_id
-- 1	1
-- Example 2:

-- Table content: shared_bicycles

-- id	bike_id	user_id
-- 1	1	1
-- 2	1	1
-- 3	1	1
-- 4	2	1
-- 5	1	2
-- 6	2	1
-- 7	2	1
-- After running your SQL statement, the table should return:

-- bike_id	user_id
-- 1	1
-- 2	1
SELECT 
	bike_id,
	user_id
FROM 
	shared_bicycles
GROUP BY bike_id,user_id
HAVING COUNT(*) >= 3