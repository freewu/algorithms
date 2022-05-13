-- 2460 · Player Height Ranking1
-- # Description
-- Write an SQL statement to get the height rank of the players table.
-- If two players are the same height, then both players have the same height rank. There should be no interval between the ranks.

-- Table definition: players

-- columns_name	type	explanation
-- id	int unsigned	primary key
-- height	int	player height

-- There should be no spacing between names
-- In MySQL, if you want to escape reserved words used as column names, you can use backquotes before and after the keyword
-- Example
-- Sample 1:

-- Table content: players

-- id	height
-- 1	198
-- 2	226
-- 3	200
-- 4	226
-- As with the players table above, the SQL query should return.

-- height	Rank
-- 226	1
-- 226	1
-- 200	2
-- 198	3
-- Sample 2:

-- Table Contents: players

-- id	height
-- 1	198
-- 2	198
-- 3	198
-- As with the players table above, the SQL query should return.

-- height	Rank
-- 198	1
-- 198	1
-- 198	1

-- Write your SQL Query here --
--@row:=@row + 1 
SELECT
    p.height,
    r.`rank` AS `Rank`
FROM
    players AS p,
    (
        SELECT 
            CAST(@row:=@row + 1 AS DECIMAL(10)) AS `rank`,
            ps.height 
        FROM 
            (
                SELECT
                    DISTINCT(height) AS height
                FROM
                    players
                ORDER BY
                    height desc
            ) AS ps,
            (
                SELECT 
                    @row:= 0
            ) AS n 
    ) AS r 
WHERE
    p.height = r.height
ORDER BY
    p.height DESC


-- dense_rank
SELECT 
    height,
    DENSE_RANK() OVER(ORDER BY height DESC) AS `Rank`
FROM 
    `players`
ORDER BY 
    `Rank`;