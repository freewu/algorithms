-- 2459 · Nth highest player height
-- Description
-- Write an SQL statement to get the height of the Nth tallest player in the players table - nth_height

-- Table definition: players

-- columns_name	type	explanation
-- id	int unsigned	primary key
-- height	int	player height

-- Note that the column name of the output is: nth_height
-- You can only output a table with one row and one column

-- Example
-- Example 1:

-- Table content: players

-- id	height
-- 1	198
-- 2	226
-- 3	200
-- 4	226
-- As with the players table above, when N = 2, the SQL query should return 200 as the second tallest height. 
-- If the second tallest height does not exist, then the query should return null

-- nth_height
-- 200
-- Example 2:

-- Table content: players

-- id	height
-- 1	198
-- 2	198
-- 3	198
-- As with the players table above, the SQL query should return null when N = 4

-- nth_height
-- null

CREATE FUNCTION get_nth_height(N INT) RETURNS INT
BEGIN
    DECLARE M INT;
    SET M = N-1;
    RETURN (
        -- Write your SQL Query here --
        SELECT 
            p.height AS nth_height
        FROM 
            (
                SELECT
                    height
                FROM
                    players
                GROUP BY
                    height
            ) AS p
        ORDER BY 
            p.height DESC 
        LIMIT
            1
        OFFSET
            M
    );
END


CREATE FUNCTION get_nth_height(N INT) RETURNS INT
BEGIN
    DECLARE M INT;
    SET M = N-1;
    RETURN (
        -- Write your SQL Query here --
        SELECT 
            DISTINCT(height) AS nth_height
        FROM 
            players
        ORDER BY 
            height DESC 
        LIMIT
            1
        OFFSET
            M
    );
END