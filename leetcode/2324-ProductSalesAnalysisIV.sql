-- 2324. Product Sales Analysis IV
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
-- product_id is a foreign key (reference column) to Product table.
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
 
-- Write a solution that reports for each user the product id on which the user spent the most money. 
-- In case the same user spent the most money on two or more products, report all of them.
-- Return the resulting table in any order.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Sales table:
-- +---------+------------+---------+----------+
-- | sale_id | product_id | user_id | quantity |
-- +---------+------------+---------+----------+
-- | 1       | 1          | 101     | 10       |
-- | 2       | 3          | 101     | 7        |
-- | 3       | 1          | 102     | 9        |
-- | 4       | 2          | 102     | 6        |
-- | 5       | 3          | 102     | 10       |
-- | 6       | 1          | 102     | 6        |
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
-- +---------+------------+
-- | user_id | product_id |
-- +---------+------------+
-- | 101     | 3          |
-- | 102     | 1          |
-- | 102     | 2          |
-- | 102     | 3          |
-- +---------+------------+ 
-- Explanation: 
-- User 101:
--     - Spent 10 * 10 = 100 on product 1.
--     - Spent 7 * 15 = 105 on product 3.
-- User 101 spent the most money on product 3.
-- User 102:
--     - Spent (9 + 7) * 10 = 150 on product 1.
--     - Spent 6 * 25 = 150 on product 2.
--     - Spent 10 * 15 = 150 on product 3.
-- User 102 spent the most money on products 1, 2, and 3.

-- Create table If Not Exists Sales (sale_id int, product_id int, user_id int, quantity int)
-- Create table If Not Exists Product (product_id int, price int)
-- Truncate table Sales
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('1', '1', '101', '10')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('2', '3', '101', '7')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('3', '1', '102', '9')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('4', '2', '102', '6')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('5', '3', '102', '10')
-- insert into Sales (sale_id, product_id, user_id, quantity) values ('6', '1', '102', '6')
-- Truncate table Product
-- insert into Product (product_id, price) values ('1', '10')
-- insert into Product (product_id, price) values ('2', '25')
-- insert into Product (product_id, price) values ('3', '15')

-- Write your MySQL query statement below
WITH t AS ( -- 统计 用户,商品,总花费
    SELECT 
        s.user_id,
        s.product_id,
        SUM(s.quantity * p.price)  AS purchased
    FROM
        Sales AS s
    LEFT JOIN
        Product AS p 
    USING(product_id)
    GROUP BY
        user_id, product_id 
)

-- SELECT * FROM t

SELECT
    user_id,
    product_id
FROM
    t 
WHERE
    (user_id, purchased) IN ( -- 每个用户最大花费
        SELECT
            user_id,
            MAX(purchased)
        FROM
            t
        GROUP BY 
            user_id
    )