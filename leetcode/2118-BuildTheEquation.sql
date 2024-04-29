-- 2118. Build the Equation
-- Table: Terms
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | power       | int  |
-- | factor      | int  |
-- +-------------+------+
-- power is the column with unique values for this table.
-- Each row of this table contains information about one term of the equation.
-- power is an integer in the range [0, 100].
-- factor is an integer in the range [-100, 100] and cannot be zero.
 
-- You have a very powerful program that can solve any equation of one variable in the world. 
--     The equation passed to the program must be formatted as follows:
--     The left-hand side (LHS) should contain all the terms.
--     The right-hand side (RHS) should be zero.
--     Each term of the LHS should follow the format "<sign><fact>X^<pow>" where:
--         <sign> is either "+" or "-".
--         <fact> is the absolute value of the factor.
--         <pow> is the value of the power.
--     If the power is 1, do not add "^<pow>".
--         For example, if power = 1 and factor = 3, the term will be "+3X".
--     If the power is 0, add neither "X" nor "^<pow>".
--         For example, if power = 0 and factor = -3, the term will be "-3".
--     The powers in the LHS should be sorted in descending order.

-- Write a solution to build the equation.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Terms table:
-- +-------+--------+
-- | power | factor |
-- +-------+--------+
-- | 2     | 1      |
-- | 1     | -4     |
-- | 0     | 2      |
-- +-------+--------+
-- Output: 
-- +--------------+
-- | equation     |
-- +--------------+
-- | +1X^2-4X+2=0 |
-- +--------------+

-- Example 2:
-- Input: 
-- Terms table:
-- +-------+--------+
-- | power | factor |
-- +-------+--------+
-- | 4     | -4     |
-- | 2     | 1      |
-- | 1     | -1     |
-- +-------+--------+
-- Output: 
-- +-----------------+
-- | equation        |
-- +-----------------+
-- | -4X^4+1X^2-1X=0 |
-- +-----------------+
 
-- Follow up: What will be changed in your solution if the power is not a primary key but each power should be unique in the answer?

-- Create table If Not Exists Terms (power int, factor int)
-- Truncate table Terms
-- insert into Terms (power, factor) values ('2', '1')
-- insert into Terms (power, factor) values ('1', '-4')
-- insert into Terms (power, factor) values ('0', '2')

-- Write your MySQL query statement below
WITH t0 AS ( -- factor 补全 + 号
    SELECT 
        power, 
        CASE 
            WHEN factor > 0 THEN CONCAT("+", factor) 
            ELSE factor 
        END AS factor
    FROM  
        Terms
),
t1 AS ( -- 处理 power 为 0 & 1 的特殊情况
    SELECT 
        *,
        CASE 
            WHEN power = 0 THEN factor   -- 2X^0 = 2
            WHEN power = 1 THEN CONCAT(factor,"X") -- 2X^1 = 2X
            ELSE CONCAT(factor, "X^", power) -- 2X^2 
        END AS term
    FROM 
        t0
    ORDER BY power DESC -- 需要根据 power 从高到低输出 如: 2X^3 + X^2 - X - 4
)

SELECT 
    CONCAT(GROUP_CONCAT(term ORDER BY power DESC separator ""),  "=0") AS equation
FROM 
    t1