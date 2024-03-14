-- 1285. Find the Start and End Number of Continuous Ranges
-- Table: Logs
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | log_id        | int     |
-- +---------------+---------+
-- log_id is the column of unique values for this table.
-- Each row of this table contains the ID in a log Table.
 
-- Write a solution to find the start and end number of continuous ranges in the table Logs.
-- Return the result table ordered by start_id.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Logs table:
-- +------------+
-- | log_id     |
-- +------------+
-- | 1          |
-- | 2          |
-- | 3          |
-- | 7          |
-- | 8          |
-- | 10         |
-- +------------+
-- Output: 
-- +------------+--------------+
-- | start_id   | end_id       |
-- +------------+--------------+
-- | 1          | 3            |
-- | 7          | 8            |
-- | 10         | 10           |
-- +------------+--------------+
-- Explanation: 
-- The result table should contain all ranges in table Logs.
-- From 1 to 3 is contained in the table.
-- From 4 to 6 is missing in the table
-- From 7 to 8 is contained in the table.
-- Number 9 is missing from the table.
-- Number 10 is contained in the table.

-- Create table If Not Exists Logs (log_id int)
-- Truncate table Logs
-- insert into Logs (log_id) values ('1')
-- insert into Logs (log_id) values ('2')
-- insert into Logs (log_id) values ('3')
-- insert into Logs (log_id) values ('7')
-- insert into Logs (log_id) values ('8')
-- insert into Logs (log_id) values ('10')

-- 解题思路:
-- 1 先给每个数进行排名 ②用这些数减去自己的排名，如果减了之后的结果是一样的，说明这几个数是连续的
-- 2 用 logid 减去排名得出来的数进行 group by
-- 3 再把连续的数全都放在一个一个小组里面，求出每个小组的最大值和最小值就可以了
-- SELECT 
--     a.*,
--     (a.log_id - a.rk)
-- FROM
-- (
--     SELECT 
--         log_id,
--         rank() OVER(ORDER BY log_id ASC) AS rk
--     FROM 
--         Logs
-- ) AS a
-- | log_id | rk | (a.log_id - a.rk) |
-- | ------ | -- | ----------------- |
-- | 1      | 1  | 0                 |
-- | 2      | 2  | 0                 |
-- | 3      | 3  | 0                 |
-- | 7      | 4  | 3                 |
-- | 8      | 5  | 3                 |
-- | 10     | 6  | 4                 |

SELECT 
    MIN(a.log_id) AS start_id,
    MAX(a.log_id) AS end_id
FROM
(
    SELECT 
        log_id,
        rank() OVER(ORDER BY log_id ASC) AS rk
    FROM 
        Logs
) AS a
GROUP BY
    (a.log_id - a.rk)
ORDER BY
    start_id ASC; -- Return the result table ordered by start_id