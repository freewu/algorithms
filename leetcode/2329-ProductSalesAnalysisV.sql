-- 2329. Product Sales Analysis V
-- Table: Sales
-- +-------------+-------+
-- | Column Name | Type  |
-- +-------------+-------+
-- | sale_id     | int   |
-- | product_id  | int   |
-- | user_id     | int   |
-- | quantity    | int   |
-- +-------------+-------+
-- sale_id contains unique values.
-- product_id is a foreign key (column with unique values) to Product table.
-- Each row of this table shows the ID of the product and the quantity purchased by a user.

-- Table: Product
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | product_id  | int  |
-- | price       | int  |
-- +-------------+------+
-- product_id contains unique values.
-- Each row of this table indicates the price of each product.

-- Write a solution to report the spending of each user.
-- Return the resulting table ordered by spending in descending order. In case of a tie, order them by user_id in ascending order.
-- The result format is in the following example.
 
-- Example 1:
-- Input: 
-- Sales table:
-- +---------+------------+---------+----------+
-- | sale_id | product_id | user_id | quantity |
-- +---------+------------+---------+----------+
-- | 1       | 1          | 101     | 10       |
-- | 2       | 2          | 101     | 1        |
-- | 3       | 3          | 102     | 3        |
-- | 4       | 3          | 102     | 2        |
-- | 5       | 2          | 103     | 3        |
-- +---------+------------+---------+----------+
-- Product table:
-- +------------+-------+
-- | product_id | price |
-- +------------+-------+
-- | 1          | 10    |
-- | 2          | 25    |
-- | 3          | 15    |
-- +------------+-------+
-- Output: 
-- +---------+----------+
-- | user_id | spending |
-- +---------+----------+
-- | 101     | 125      |
-- | 102     | 75       |
-- | 103     | 75       |
-- +---------+----------+
-- Explanation: 
-- User 101 spent 10 * 10 + 1 * 25 = 125.
-- User 102 spent 3 * 15 + 2 * 15 = 75.
-- User 103 spent 3 * 25 = 75.
-- Users 102 and 103 spent the same amount and we break the tie by their ID while user 101 is on the top.

-- Create table If Not Exists Sales (sale_id int, product_id int, user_id int, quantity int)
-- Create table If Not Exists Product (product_id int, price int)
-- Truncate table Sales
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('1', '1', '101', '10')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('2', '2', '101', '1')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('3', '3', '102', '3')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('4', '3', '102', '2')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('5', '2', '103', '3')
-- Truncate table Product
-- insert into Product (product_id, price) values ('1', '10')
-- insert into Product (product_id, price) values ('2', '25')
-- insert into Product (product_id, price) values ('3', '15')

-- Write your MySQL query statement below
SELECT
    s.user_id AS user_id,
    SUM(p.price  *  s.quantity) AS spending
FROM
    Sales AS s
LEFT JOIN 
    Product AS p 
ON 
    s.product_id = p.product_id 
GROUP BY
    s.user_id
ORDER BY 
    -- 按用户消费额 spending 递减的顺序返回结果。在消费额相等的情况下，以 user_id 递增的顺序将其排序
    spending DESC, user_id ASC