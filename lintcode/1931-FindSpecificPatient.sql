-- 1931 · Find Specific Patient
-- # Description
-- The patients table saves all patient information and the person who infected them (infected_by_id)
-- Write an SQL statement to return a list. The names appearing in the list need to meet the condition: the id of the person who infected them is not 2.

-- Table definition: patients

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	patient name
-- infected_by_id	int	infected id
-- Example
-- Example 1:

-- Table content: patients

-- id	name	infected_by_id
-- 1	Amy	null
-- 2	Bob	null
-- 3	Catalina	2
-- 4	Deng	null
-- 5	Eason	1
-- 6	Frank	2
-- After running your SQL statement, the table should return:

-- name
-- Amy
-- Bob
-- Deng
-- Eason
-- Example 2:

-- Table content: patients

-- id	name	infected_by_id
-- 1	Amy	null
-- 2	Bob	null
-- 3	Catalina	2
-- After running your SQL statement, the table should return:

-- name
-- Amy
-- Bob
SELECT
	name
FROM 
	patients
WHERE
	infected_by_id IS NULL OR
	infected_by_id <> 2

-- infected_by_id IS NULL 必须有