-- 1925 · At Least Three Players Have Scored
-- Description
-- There is an player_scores table with player and score
-- Please list all the points scored by 3 or more players

-- For example, in the table below, the score of 50 was obtained by Bryant, Carter, and Durant.
-- So it should output 50

-- Table definition: player_scores

-- columns_name	type	explaination
-- player	varchar	player's name
-- score	int	player's score
-- Example
-- Example 1:

-- Table content: player_scores

-- player	score
-- Jordan	63
-- Iverson	55
-- Bryant	50
-- Carter	50
-- McGrady	46
-- James	51
-- Durant	50
-- Wade	46
-- Anthony	42
-- Ginobili	39
-- After running your SQL statement, the table should return:

-- score
-- 50
-- Example 2:

-- Table content: player_scores

-- player	score
-- Jordan	66
-- Iverson	66
-- Bryant	66
-- Carter	66
-- McGrady	50
-- James	50
-- Durant	50
-- Wade	46
-- Anthony	46
-- Ginobili	42
-- After running your SQL statement, the table should return:

-- score
-- 50
-- 66
SELECT 
	score
FROM 
	player_scores
GROUP BY score
HAVING count(*) >= 3