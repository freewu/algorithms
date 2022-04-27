-- 1920 · Find a Classmate with the Same Name
-- Description
-- Write a SQL query to find all students with the same name in the students table.

-- Table definition: students

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	student's name
-- Example
-- Example 1:

-- Table content: students

-- id	name
-- 1	DaMing
-- 2	Amy
-- 3	HanMeimei
-- 4	Amy
-- Your SQL query should return the following results:

-- name
-- Amy
-- Example 2:

-- Table content: students

-- id	name
-- 1	DaMing
-- 2	Amy
-- 3	DaMing
-- 4	Amy
-- Your SQL query should return the following results:

-- name
-- Amy
-- DaMing

SELECT 
	name
FROM
	students
GROUP BY name
HAVING count(*) > 1