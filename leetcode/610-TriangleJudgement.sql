-- 610. Triangle Judgement
-- Table: Triangle
--
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | x           | int  |
-- | y           | int  |
-- | z           | int  |
-- +-------------+------+
-- (x, y, z) is the primary key column for this table.
-- Each row of this table contains the lengths of three line segments.
--  
-- Write an SQL query to report for every three line segments whether they can form a triangle.
-- Return the result table in any order.
-- The query result format is in the following example.
--

-- Example 1:
--
-- Input:
-- Triangle table:
-- +----+----+----+
-- | x  | y  | z  |
-- +----+----+----+
-- | 13 | 15 | 30 |
-- | 10 | 20 | 15 |
-- +----+----+----+
-- Output:
-- +----+----+----+----------+
-- | x  | y  | z  | triangle |
-- +----+----+----+----------+
-- | 13 | 15 | 30 | No       |
-- | 10 | 20 | 15 | Yes      |
-- +----+----+----+----------+
--
-- 三角形三边关系是三角形三条边关系的定则，具体内容是在一个三角形中，任意两边之和大于第三边，任意两边之差小于第三边。”

-- Write your MySQL query statement below
SELECT
    *,
    CASE
        WHEN x + y > z AND x + z > y AND y + z > x THEN 'Yes'
        ELSE 'No'
    END AS `triangle`
FROM
    `triangle`
