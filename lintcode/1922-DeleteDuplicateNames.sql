-- 1922 · Delete Duplicate Names
-- Description
-- Write a SQL statement to delete all duplicate names in the contacts table, and keep only the one with the smallest id among the duplicate names.

-- Table definition: contacts

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	name
-- Use the delete statement
-- After executing your SQL statement, we will execute SELECT * FROM contacts to verify your statement
-- We will individually verify whether the data in the database has been modified to the following information.

-- Example
-- Example 1:

-- Table content: contacts

-- id	name
-- 1	Song Jiang
-- 2	Lu Junyi
-- 3	Wu Yong
-- 4	Song Jiang
-- 5	Wu Yong
-- 6	Lin Chong
-- After running your SQL statement, the Chart contacts above should return the following lines:

-- id	name
-- 1	Song Jiang
-- 2	Lu Junyi
-- 3	Wu Yong
-- 6	Lin Chong
-- Example 2:

-- Table content: contacts

-- id	name
-- 1	Song Jiang
-- 2	Lu Junyi
-- 3	Wu Yong
-- 4	Song Jiang
-- After running your SQL statement, the Chart contacts above should return the following lines:

-- id	name
-- 1	Song Jiang
-- 2	Lu Junyi
-- 3	Wu Yong

DELETE 
FROM contacts 
WHERE id NOT IN (
	SELECT
		id
	FROM (
		SELECT 
			MIN(id) AS id
		FROM contacts
		GROUP BY name
	) AS p
);

-- 不能直接  DELETE FROM contacts WHERE id NOT IN (SELECT MIN(id) AS id FROM contacts GROUP BY name)
-- 会抛出 You can't specify target table 'contacts' for update in FROM clause 异常