-- 2893. Calculate Orders Within Each Interval
-- Table: Orders

-- +-------------+------+ 
-- | Column Name | Type | 
-- +-------------+------+ 
-- | minute      | int  | 
-- | order_count | int  |
-- +-------------+------+
-- minute is the primary key for this table.
-- Each row of this table contains the minute and number of orders received during that specific minute. The total number of rows will be a multiple of 6.
-- Write a query to calculate total orders within each interval. Each interval is defined as a combination of 6 minutes.
--     Minutes 1 to 6 fall within interval 1, while minutes 7 to 12 belong to interval 2, and so forth.

-- Return the result table ordered by interval_no in ascending order.
-- The result format is in the following example.
 
-- Example 1:
-- Input: 
-- Orders table:
-- +--------+-------------+
-- | minute | order_count | 
-- +--------+-------------+
-- | 1      | 0           |
-- | 2      | 2           | 
-- | 3      | 4           | 
-- | 4      | 6           | 
-- | 5      | 1           | 
-- | 6      | 4           | 
-- | 7      | 1           | 
-- | 8      | 2           | 
-- | 9      | 4           | 
-- | 10     | 1           | 
-- | 11     | 4           | 
-- | 12     | 6           | 
-- +--------+-------------+
-- Output: 
-- +-------------+--------------+
-- | interval_no | total_orders | 
-- +-------------+--------------+
-- | 1           | 17           | 
-- | 2           | 18           |    
-- +-------------+--------------+
-- Explanation: 
-- - Interval number 1 comprises minutes from 1 to 6. The total orders in these six minutes are (0 + 2 + 4 + 6 + 1 + 4) = 17.
-- - Interval number 2 comprises minutes from 7 to 12. The total orders in these six minutes are (1 + 2 + 4 + 1 + 4 + 6) = 18.
-- Returning table orderd by interval_no in ascending order.

-- Create table if not exists Orders(minute int, order_count int)
-- Truncate table Orders
-- insert into Orders (minute, order_count) values ('1', '0')
-- insert into Orders (minute, order_count) values ('2', '2')
-- insert into Orders (minute, order_count) values ('3', '4')
-- insert into Orders (minute, order_count) values ('4', '6')
-- insert into Orders (minute, order_count) values ('5', '1')
-- insert into Orders (minute, order_count) values ('6', '4')
-- insert into Orders (minute, order_count) values ('7', '1')
-- insert into Orders (minute, order_count) values ('8', '2')
-- insert into Orders (minute, order_count) values ('9', '4')
-- insert into Orders (minute, order_count) values ('10', '1')
-- insert into Orders (minute, order_count) values ('11', '4')
-- insert into Orders (minute, order_count) values ('12', '6')

-- Write your MySQL query statement below
-- SELECT 
--     interval_no,
--     SUM(order_count) AS total_orders 
-- FROM 
-- (
--     SELECT 
--         CASE 
--             WHEN minute >= 1 AND minute < 7 THEN 1
--             WHEN minute >= 7 AND minute < 13 THEN 2
--             WHEN minute >= 13 AND minute < 19 THEN 3
--             WHEN minute >= 19 AND minute < 25 THEN 4
--             WHEN minute >= 25 AND minute < 31 THEN 5
--             WHEN minute >= 31 AND minute < 37 THEN 6
--             WHEN minute >= 37 AND minute < 43 THEN 7
--             WHEN minute >= 43 AND minute < 49 THEN 8
--             WHEN minute >= 49 AND minute < 55 THEN 9
--             ELSE 10
--         END AS interval_no,
--         order_count
--     FROM
--         Orders
-- ) AS t 
-- GROUP BY 
--     interval_no
-- ORDER BY 
--     interval_no

WITH t AS (
    SELECT
        minute,
        SUM(order_count) OVER(ORDER BY minute ROWS BETWEEN 5 PRECEDING AND current ROW) AS total_orders
    FROM 
        Orders
)

-- SELECT * FROM t
SELECT 
    ROW_NUMBER() OVER(ORDER BY minute) AS interval_no,
    total_orders
FROM 
    t
WHERE
    minute % 6 = 0

-- ROUND
SELECT 
    ROUND(( minute + 5) /6 -0.5,0) AS interval_no,
    SUM(order_count) as total_orders
FROM 
    orders
GROUP BY 
    interval_no
ORDER BY 
    interval_no