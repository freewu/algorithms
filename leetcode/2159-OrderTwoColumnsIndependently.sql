-- 2159. Order Two Columns Independently
-- Table: Data
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | first_col   | int  |
-- | second_col  | int  |
-- +-------------+------+
-- This table may contain duplicate rows.
 
-- Write a solution to independently:
--     order first_col in ascending order.
--     order second_col in descending order.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Data table:
-- +-----------+------------+
-- | first_col | second_col |
-- +-----------+------------+
-- | 4         | 2          |
-- | 2         | 3          |
-- | 3         | 1          |
-- | 1         | 4          |
-- +-----------+------------+
-- Output: 
-- +-----------+------------+
-- | first_col | second_col |
-- +-----------+------------+
-- | 1         | 4          |
-- | 2         | 3          |
-- | 3         | 2          |
-- | 4         | 1          |
-- +-----------+------------+

-- Create table If Not Exists Data (first_col int, second_col int)
-- Truncate table Data
-- insert into Data (first_col, second_col) values ('4', '2')
-- insert into Data (first_col, second_col) values ('2', '3')
-- insert into Data (first_col, second_col) values ('3', '1')
-- insert into Data (first_col, second_col) values ('1', '4')

-- Write your MySQL query statement below
SELECT 
    a.first_col,
    b.second_col 
FROM
    (
        SELECT 
            first_col,
            ROW_NUMBER() OVER (ORDER BY first_col) AS rn 
        FROM
            Data 
    ) AS a,
    (
        SELECT 
            second_col,
            ROW_NUMBER() OVER (ORDER BY second_col DESC) AS rn 
        FROM
            Data 
    ) AS b
WHERE
    a.rn = b.rn

-- solution 2
WITH t AS
(
    SELECT 
        first_col, 
        ROW_NUMBER() OVER (ORDER BY first_col DESC) AS rn1,
        second_col, 
        ROW_NUMBER() OVER (ORDER BY second_col DESC) AS rn2
    FROM 
        Data
) 

SELECT
    a.first_col,
    b.second_col
FROM
    t AS a,
    t AS b 
WHERE
    a.rn1 = b.rn2 
ORDER BY
    a.rn1