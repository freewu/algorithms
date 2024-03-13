-- 1440. Evaluate Boolean Expression
-- Table Variables:
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | name          | varchar |
-- | value         | int     |
-- +---------------+---------+
-- In SQL, name is the primary key for this table.
-- This table contains the stored variables and their values.

-- Table Expressions:
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | left_operand  | varchar |
-- | operator      | enum    |
-- | right_operand | varchar |
-- +---------------+---------+
-- In SQL, (left_operand, operator, right_operand) is the primary key for this table.
-- This table contains a boolean expression that should be evaluated.
-- operator is an enum that takes one of the values ('<', '>', '=')
-- The values of left_operand and right_operand are guaranteed to be in the Variables table.

-- Evaluate the boolean expressions in Expressions table.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Variables table:
-- +------+-------+
-- | name | value |
-- +------+-------+
-- | x    | 66    |
-- | y    | 77    |
-- +------+-------+
-- Expressions table:
-- +--------------+----------+---------------+
-- | left_operand | operator | right_operand |
-- +--------------+----------+---------------+
-- | x            | >        | y             |
-- | x            | <        | y             |
-- | x            | =        | y             |
-- | y            | >        | x             |
-- | y            | <        | x             |
-- | x            | =        | x             |
-- +--------------+----------+---------------+
-- Output: 
-- +--------------+----------+---------------+-------+
-- | left_operand | operator | right_operand | value |
-- +--------------+----------+---------------+-------+
-- | x            | >        | y             | false |
-- | x            | <        | y             | true  |
-- | x            | =        | y             | false |
-- | y            | >        | x             | true  |
-- | y            | <        | x             | false |
-- | x            | =        | x             | true  |
-- +--------------+----------+---------------+-------+
-- Explanation: 
-- As shown, you need to find the value of each boolean expression in the table using the variables table.

-- Create Table If Not Exists Variables (name varchar(3), value int);
-- Create Table If Not Exists Expressions (left_operand varchar(3), operator ENUM('>', '<', '='), right_operand varchar(3));
-- Truncate table Variables;
-- insert into Variables (name, value) values ('x', '66');
-- insert into Variables (name, value) values ('y', '77');
-- Truncate table Expressions;
-- insert into Expressions (left_operand, operator, right_operand) values ('x', '>', 'y');
-- insert into Expressions (left_operand, operator, right_operand) values ('x', '<', 'y');
-- insert into Expressions (left_operand, operator, right_operand) values ('x', '=', 'y');
-- insert into Expressions (left_operand, operator, right_operand) values ('y', '>', 'x');
-- insert into Expressions (left_operand, operator, right_operand) values ('y', '<', 'x');
-- insert into Expressions (left_operand, operator, right_operand) values ('x', '=', 'x');

-- step 1 先组一张宽表 | left_operand | operator | right_operand | x  | y  |
-- SELECT
--     e.*,
--     v.*
-- FROM
--     Expressions AS e,
--     (
--         SELECT  
--             SUM(x) AS x,
--             SUM(y) AS y
--         FROM 
--         (
--             SELECT value AS x, 0 AS y FROM Variables WHERE name = 'x'
--             UNION ALL
--             SELECT 0 AS x, value AS y FROM Variables WHERE name = 'y'
--         ) AS u
--     ) AS v
-- +-------------+---------+--------------+----+----+
-- | left_operand | operator | right_operand | x  | y  |
-- +-------------+---------+--------------+----+----+
-- | x           | >       | y            | 66 | 77 |
-- | x           | <       | y            | 66 | 77 |
-- | x           | =       | y            | 66 | 77 |
-- | y           | >       | x            | 66 | 77 |
-- | y           | <       | x            | 66 | 77 |
-- | x           | =       | x            | 66 | 77 |
-- +-------------+---------+--------------+----+----+

-- SELECT
--     e.*,
--     (
--         IF(
--             (IF(left_operand = x, v.x, v.y )

--             IF(right_operand = x, v.x, v.y )),
--             "true",
--             "false"
--         )
--     ) AS value
-- FROM
--     Expressions AS e,
--     (
--         SELECT  
--             SUM(x) AS x,
--             SUM(y) AS y
--         FROM 
--         (
--             SELECT value AS x, 0 AS y FROM Variables WHERE name = 'x'
--             UNION ALL
--             SELECT 0 AS x, value AS y FROM Variables WHERE name = 'y'
--         ) AS u
--     ) AS v

SELECT 
    e.*,
    CASE
        WHEN operator = '=' AND v1.value = v2.value THEN 'true'
        WHEN operator = '>' AND v1.value > v2.value THEN 'true'
        WHEN operator = '<' AND v1.value < v2.value THEN 'true'
        ELSE 'false'
    END AS value
FROM Expressions AS e
LEFT JOIN Variables AS v1 ON e.left_operand = v1.name
LEFT JOIN Variables AS v2 ON e.right_operand = v2.name