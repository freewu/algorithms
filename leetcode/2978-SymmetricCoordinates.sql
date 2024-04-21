-- 2978. Symmetric Coordinates
-- Table: Coordinates
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | X           | int  |
-- | Y           | int  |
-- +-------------+------+
-- Each row includes X and Y, where both are integers. Table may contain duplicate values.
-- Two coordindates (X1, Y1) and (X2, Y2) are said to be symmetric coordintes if X1 == Y2 and X2 == Y1.

-- Write a solution that outputs, among all these symmetric coordintes, only those unique coordinates that satisfy the condition X1 <= Y1.

-- Return the result table ordered by X and Y (respectively) in ascending order.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Coordinates table:
-- +----+----+
-- | X  | Y  |
-- +----+----+
-- | 20 | 20 |
-- | 20 | 20 |
-- | 20 | 21 |
-- | 23 | 22 |
-- | 22 | 23 |
-- | 21 | 20 |
-- +----+----+
-- Output: 
-- +----+----+
-- | x  | y  |
-- +----+----+
-- | 20 | 20 |
-- | 20 | 21 |
-- | 22 | 23 |
-- +----+----+
-- Explanation: 
-- - (20, 20) and (20, 20) are symmetric coordinates because, X1 == Y2 and X2 == Y1. This results in displaying (20, 20) as a distinctive coordinates.
-- - (20, 21) and (21, 20) are symmetric coordinates because, X1 == Y2 and X2 == Y1. However, only (20, 21) will be displayed because X1 <= Y1.
-- - (23, 22) and (22, 23) are symmetric coordinates because, X1 == Y2 and X2 == Y1. However, only (22, 23) will be displayed because X1 <= Y1.
-- The output table is sorted by X and Y in ascending order.

-- Create table If Not Exists Coordinates (X int, Y int)
-- Truncate table Coordinates
-- insert into Coordinates (X, Y) values ('20', '20')
-- insert into Coordinates (X, Y) values ('20', '20')
-- insert into Coordinates (X, Y) values ('20', '21')
-- insert into Coordinates (X, Y) values ('23', '22')
-- insert into Coordinates (X, Y) values ('22', '23')
-- insert into Coordinates (X, Y) values ('21', '20')

# Write your MySQL query statement below
WITH t AS (
    SELECT
        *,
        ROW_NUMBER() OVER(ORDER BY X) AS rn
    FROM 
        Coordinates
)

-- SELECT 
--     a.*,
--     b.*
-- FROM
--     Coordinates AS a
-- LEFT JOIN
--     Coordinates AS b 
-- ON
--     a.X = b.Y AND a.Y = b.X
-- WHERE
--     a.X <= a.Y AND -- 满足条件 X1 <= Y1 的唯一坐标
--     b.X IS NOT NULL

SELECT 
    a.X,
    a.Y
FROM
    t AS a
LEFT JOIN
    t AS b 
ON
    a.X = b.Y AND a.Y = b.X AND a.rn != b.rn -- 需要排除自己
WHERE
    a.X <= a.Y AND -- 满足条件 X1 <= Y1 的唯一坐标
    b.X IS NOT NULL
GROUP BY
    a.X, a.Y
ORDER BY 
    a.X, a.Y -- 按照 X 和 Y 分别 升序 排列结果表。

-- DISTINCT
SELECT 
    DISTINCT a.X AS X,
    a.Y AS Y
FROM
    t AS a
JOIN
    t AS b 
ON
    a.X = b.Y AND a.Y = b.X AND a.rn != b.rn -- 需要排除自己
WHERE
    a.X <= a.Y AND -- 满足条件 X1 <= Y1 的唯一坐标
ORDER BY 
    a.X, a.Y -- 按照 X 和 Y 分别 升序 排列结果表。