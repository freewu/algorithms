-- 2499 · Height of players appearing consecutively
-- # Description
-- Write an SQL statement to find the height of a player in the players table that has at least three players 
-- with consecutive ids reaching this height, and return the results in any order.

-- Table definition: players (players table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- height	int	player height

-- The returned results can be arranged in any order

-- Example
-- Sample 1:

-- Table content: players

-- id	height
-- 1	198
-- 2	198
-- 3	198
-- 4	226
-- As in the players table above, the players with ids 1, 2 and 3 all have a height value of 198, so 198 is returned.

-- height
-- 198
-- Sample 2:

-- Table Contents: players

-- id	height
-- 1	198
-- 2	196
-- 3	188
-- If the players table above does not have a height that meets the criteria, nothing will be returned

-- height
-- null

SELECT
	height
FROM
	players
GROUP BY
	height
HAVING 
	COUNT(*) >= 3