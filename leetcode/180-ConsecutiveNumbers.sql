-- 180. Consecutive Numbers
-- Table: Logs
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | num         | varchar |
-- +-------------+---------+
-- id is the primary key for this table.
-- id is an autoincrement column.
--  
--
-- Write an SQL query to find all numbers that appear at least three times consecutively.
--
-- Return the result table in any order.
--
-- The query result format is in the following example.
--
--  
--
-- Example 1:
--
-- Input:
-- Logs table:
-- +----+-----+
-- | id | num |
-- +----+-----+
-- | 1  | 1   |
-- | 2  | 1   |
-- | 3  | 1   |
-- | 4  | 2   |
-- | 5  | 1   |
-- | 6  | 2   |
-- | 7  | 2   |
-- +----+-----+
-- Output:
-- +-----------------+
-- | ConsecutiveNums |
-- +-----------------+
-- | 1               |
-- +-----------------+
-- Explanation: 1 is the only number that appears consecutively for at least three times.
--
SELECT
    DISTINCT a.num  AS ConsecutiveNums
FROM
    Logs AS a,
    Logs AS b,
    Logs AS c
WHERE
    a.num = b.num AND
    b.num = c.num AND
    a.id + 1 = b.id AND
    a.id + 2 = c.id

-- best solution
select
    distinct Num as ConsecutiveNums
from (
    select
        Num,
        lead(Num,1) over() as Num2,
        lead(Num,2) over() as Num3
    from Logs
    ) t
where
    Num = Num2 and
    Num = Num3

-- 一、LEAD(expr [, N[, default]] over(partition by order by )
-- 功能：会返回分区内当前行后边N行的字段值，如果没有这样的行，会返回你设置的default （如果N和default省略了，则默认为 1和null)
-- 注意：
-- ①N必须是一个非负整数，如果N=0，expr是当前行的字段值
-- ②从mysql8.0.22开始，N不能为空，且必须是1-2^63的整数（可以是变量）
-- 二、LAG(expr [, N[, default]]) over(partition by order by )
-- 功能：会返回分组内当前行前边N行的字段值，如果没有这样的行，会返回你设置的default
-- 注意：当使用了多个窗口函数，且这些窗口函数中的over()的内容都相同，可使用命名窗口函数。
