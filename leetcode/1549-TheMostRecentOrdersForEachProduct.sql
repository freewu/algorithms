-- 1549. The Most Recent Orders for Each Product
-- Table: Customers
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | customer_id   | int     |
-- | name          | varchar |
-- +---------------+---------+
-- customer_id is the column with unique values for this table.
-- This table contains information about the customers.
 
-- Table: Orders
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | order_id      | int     |
-- | order_date    | date    |
-- | customer_id   | int     |
-- | product_id    | int     |
-- +---------------+---------+
-- order_id is the column with unique values for this table.
-- This table contains information about the orders made by customer_id.
-- There will be no product ordered by the same user more than once in one day.
 
-- Table: Products
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | product_id    | int     |
-- | product_name  | varchar |
-- | price         | int     |
-- +---------------+---------+
-- product_id is the column with unique values for this table.
-- This table contains information about the Products.
-- Write a solution to find the most recent order(s) of each product.
-- Return the result table ordered by product_name in ascending order and in case of a tie by the product_id in ascending order. 
-- If there still a tie, order them by order_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Customers table:
-- +-------------+-----------+
-- | customer_id | name      |
-- +-------------+-----------+
-- | 1           | Winston   |
-- | 2           | Jonathan  |
-- | 3           | Annabelle |
-- | 4           | Marwan    |
-- | 5           | Khaled    |
-- +-------------+-----------+
-- Orders table:
-- +----------+------------+-------------+------------+
-- | order_id | order_date | customer_id | product_id |
-- +----------+------------+-------------+------------+
-- | 1        | 2020-07-31 | 1           | 1          |
-- | 2        | 2020-07-30 | 2           | 2          |
-- | 3        | 2020-08-29 | 3           | 3          |
-- | 4        | 2020-07-29 | 4           | 1          |
-- | 5        | 2020-06-10 | 1           | 2          |
-- | 6        | 2020-08-01 | 2           | 1          |
-- | 7        | 2020-08-01 | 3           | 1          |
-- | 8        | 2020-08-03 | 1           | 2          |
-- | 9        | 2020-08-07 | 2           | 3          |
-- | 10       | 2020-07-15 | 1           | 2          |
-- +----------+------------+-------------+------------+
-- Products table:
-- +------------+--------------+-------+
-- | product_id | product_name | price |
-- +------------+--------------+-------+
-- | 1          | keyboard     | 120   |
-- | 2          | mouse        | 80    |
-- | 3          | screen       | 600   |
-- | 4          | hard disk    | 450   |
-- +------------+--------------+-------+
-- Output: 
-- +--------------+------------+----------+------------+
-- | product_name | product_id | order_id | order_date |
-- +--------------+------------+----------+------------+
-- | keyboard     | 1          | 6        | 2020-08-01 |
-- | keyboard     | 1          | 7        | 2020-08-01 |
-- | mouse        | 2          | 8        | 2020-08-03 |
-- | screen       | 3          | 3        | 2020-08-29 |
-- +--------------+------------+----------+------------+
-- Explanation: 
-- keyboard's most recent order is in 2020-08-01, it was ordered two times this day.
-- mouse's most recent order is in 2020-08-03, it was ordered only once this day.
-- screen's most recent order is in 2020-08-29, it was ordered only once this day.
-- The hard disk was never ordered and we do not include it in the result table.

-- Create table If Not Exists Customers (customer_id int, name varchar(10))
-- Create table If Not Exists Orders (order_id int, order_date date, customer_id int, product_id int)
-- Create table If Not Exists Products (product_id int, product_name varchar(20), price int)
-- Truncate table Customers
-- insert into Customers (customer_id, name) values ('1', 'Winston')
-- insert into Customers (customer_id, name) values ('2', 'Jonathan')
-- insert into Customers (customer_id, name) values ('3', 'Annabelle')
-- insert into Customers (customer_id, name) values ('4', 'Marwan')
-- insert into Customers (customer_id, name) values ('5', 'Khaled')
-- Truncate table Orders
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('1', '2020-07-31', '1', '1')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('2', '2020-7-30', '2', '2')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('3', '2020-08-29', '3', '3')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('4', '2020-07-29', '4', '1')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('5', '2020-06-10', '1', '2')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('6', '2020-08-01', '2', '1')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('7', '2020-08-01', '3', '1')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('8', '2020-08-03', '1', '2')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('9', '2020-08-07', '2', '3')
-- insert into Orders (order_id, order_date, customer_id, product_id) values ('10', '2020-07-15', '1', '2')
-- Truncate table Products
-- insert into Products (product_id, product_name, price) values ('1', 'keyboard', '120')
-- insert into Products (product_id, product_name, price) values ('2', 'mouse', '80')
-- insert into Products (product_id, product_name, price) values ('3', 'screen', '600')
-- insert into Products (product_id, product_name, price) values ('4', 'hard disk', '450')

-- Write your MySQL query statement below
SELECT 
    p.product_name AS product_name,
    p.product_id AS product_id,
    o.order_id AS order_id,
    o.order_date AS order_date
FROM
    Products AS p,
    Orders AS o,
    ( -- 取每个商品最近的下单时间
        SELECT
            product_id,
            MAX(order_date) AS order_date 
        FROM
            Orders
        GROUP BY 
            product_id
    ) AS o1
WHERE
    p.product_id = o.product_id AND
    o.product_id = o1.product_id AND
    o.order_date = o1.order_date
ORDER BY
    product_name ASC, -- 结果以 product_name 升序排列
    product_id ASC, -- 如果有排序相同, 再以 product_id 升序排列
    order_id ASC -- 如果还有排序相同, 再以 order_id 升序排列


SELECT
    t1.product_name,
    t2.product_id,
    t2.order_id,
    t2.order_date
FROM
    Products AS t1
INNER JOIN 
    Orders AS t2 ON t1.product_id = t2.product_id
WHERE (t2.product_id, t2.order_date) IN (
    SELECT
        product_id,
        MAX(order_date) AS 'order_date'
    FROM
        Orders
    GROUP BY product_id
)
ORDER BY 
    t1.product_name, t2.product_id, t2.order_id