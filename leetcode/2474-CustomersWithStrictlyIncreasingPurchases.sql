-- 2474. Customers With Strictly Increasing Purchases
-- Table: Orders
-- +--------------+------+
-- | Column Name  | Type |
-- +--------------+------+
-- | order_id     | int  |
-- | customer_id  | int  |
-- | order_date   | date |
-- | price        | int  |
-- +--------------+------+
-- order_id is the column with unique values for this table.
-- Each row contains the id of an order, the id of customer that ordered it, the date of the order, and its price.
 
-- Write a solution to report the IDs of the customers with the total purchases strictly increasing yearly.
--     The total purchases of a customer in one year is the sum of the prices of their orders in that year. If for some year the customer did not make any order, we consider the total purchases 0.
--     The first year to consider for each customer is the year of their first order.
--     The last year to consider for each customer is the year of their last order.

-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Orders table:
-- +----------+-------------+------------+-------+
-- | order_id | customer_id | order_date | price |
-- +----------+-------------+------------+-------+
-- | 1        | 1           | 2019-07-01 | 1100  |
-- | 2        | 1           | 2019-11-01 | 1200  |
-- | 3        | 1           | 2020-05-26 | 3000  |
-- | 4        | 1           | 2021-08-31 | 3100  |
-- | 5        | 1           | 2022-12-07 | 4700  |
-- | 6        | 2           | 2015-01-01 | 700   |
-- | 7        | 2           | 2017-11-07 | 1000  |
-- | 8        | 3           | 2017-01-01 | 900   |
-- | 9        | 3           | 2018-11-07 | 900   |
-- +----------+-------------+------------+-------+
-- Output: 
-- +-------------+
-- | customer_id |
-- +-------------+
-- | 1           |
-- +-------------+
-- Explanation: 
-- Customer 1: The first year is 2019 and the last year is 2022
--   - 2019: 1100 + 1200 = 2300
--   - 2020: 3000
--   - 2021: 3100
--   - 2022: 4700
--   We can see that the total purchases are strictly increasing yearly, so we include customer 1 in the answer.
-- Customer 2: The first year is 2015 and the last year is 2017
--   - 2015: 700
--   - 2016: 0
--   - 2017: 1000
--   We do not include customer 2 in the answer because the total purchases are not strictly increasing. Note that customer 2 did not make any purchases in 2016.
-- Customer 3: The first year is 2017, and the last year is 2018
--   - 2017: 900
--   - 2018: 900
--  We do not include customer 3 in the answer because the total purchases are not strictly increasing.

-- Create table If Not Exists Orders (order_id int, customer_id int, order_date date, price int)
-- Truncate table Orders
-- insert into Orders (order_id, customer_id, order_date, price) values ('1', '1', '2019-07-01', '1100')
-- insert into Orders (order_id, customer_id, order_date, price) values ('2', '1', '2019-11-01', '1200')
-- insert into Orders (order_id, customer_id, order_date, price) values ('3', '1', '2020-05-26', '3000')
-- insert into Orders (order_id, customer_id, order_date, price) values ('4', '1', '2021-08-31', '3100')
-- insert into Orders (order_id, customer_id, order_date, price) values ('5', '1', '2022-12-07', '4700')
-- insert into Orders (order_id, customer_id, order_date, price) values ('6', '2', '2015-01-01', '700')
-- insert into Orders (order_id, customer_id, order_date, price) values ('7', '2', '2017-11-07', '1000')
-- insert into Orders (order_id, customer_id, order_date, price) values ('8', '3', '2017-01-01', '900')
-- insert into Orders (order_id, customer_id, order_date, price) values ('9', '3', '2018-11-07', '900')

-- Write your MySQL query statement below
WITH t AS ( -- 每个客户每年的购物总和
    SELECT 
        customer_id,
        YEAR(order_date) AS year,
        SUM(price) AS amount
    FROM
        Orders 
    GROUP BY
        customer_id, YEAR(order_date)
)
SELECT 
    customer_id 
FROM 
    (
        SELECT 
            *,
            CASE 
                WHEN s < amount OR year = min_year THEN 1 
                ELSE 0 
            END AS flag  -- 标记,有那一年不连续则记为零,逐年递增,则两年的平均值要小于第二年
        FROM 
        (
            SELECT 
                *,
                AVG(amount) OVER(PARTITION BY customer_id ORDER BY year RANGE BETWEEN 1 PRECEDING AND current ROW) AS s,
                MIN(year) OVER(PARTITION BY customer_id) AS min_year -- 开始年份
            FROM 
                t
        ) AS a
    ) AS b
GROUP BY 
    customer_id 
HAVING 
    MIN(flag) != 0


-- best solution
SELECT 
    customer_id
FROM 
(
    SELECT 
        customer_id,
        year,
        amount,
        year - RANK() OVER(PARTITION BY customer_id ORDER BY amount) AS diff -- 这句是关键
    FROM 
    ( -- 每个客户每年消费
        SELECT 
            customer_id,
            YEAR(order_date) as year,
            SUM(price) as amount
        FROM 
            Orders
        GROUP BY
            customer_id, YEAR(order_date) 
    ) AS t
) AS d
GROUP BY 
    customer_id
HAVING 
    COUNT(DISTINCT diff) = 1