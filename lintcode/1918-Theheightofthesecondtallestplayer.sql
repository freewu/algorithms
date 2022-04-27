-- 1918 · The height of the second tallest player
-- Description
-- Write a SQL query to get the second tallest height in the players (height)

-- Table definition: players

-- columns_name	type	explanation
-- id	int unsigned	primary key
-- height	int	player height
-- Note that the column name of the output result: second_height

-- Example
-- Example 1:

-- Table content: players

-- id	height
-- 1	198
-- 2	226
-- 3	200
-- 4	226
-- For example, in the Table players above, the SQL query should return 200 as the second tallest height.
-- If there is no second tallest height, then the query should return null

-- second_height
-- 200
-- Example 2:

-- Table content: players

-- id	height
-- 1	198
-- 2	198
-- 3	198
-- As in the players table above, the SQL query should return null

-- second_height
-- null

SELECT 
	MAX(height) AS second_height
FROM
	players
WHERE
	height < (
		SELECT MAX(height) FROM players
	)