-- 1083. Sales Analysis II
-- Table: Product
--
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | product_id   | int     |
-- | product_name | varchar |
-- | unit_price   | int     |
-- +--------------+---------+
-- product_id is the primary key of this table.
-- Each row of this table indicates the name and the price of each product.
-- Table: Sales
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | seller_id   | int     |
-- | product_id  | int     |
-- | buyer_id    | int     |
-- | sale_date   | date    |
-- | quantity    | int     |
-- | price       | int     |
-- +-------------+---------+
-- This table has no primary key, it can have repeated rows.
-- product_id is a foreign key to the Product table.
-- Each row of this table contains some information about one sale.
--  
-- Write an SQL query that reports the buyers who have bought S8 but not iPhone. Note that S8 and iPhone are products present in the Product table.
-- Return the result table in any order.
-- The query result format is in the following example. 
--
-- Example 1:
--
-- Input:
-- Product table:
-- +------------+--------------+------------+
-- | product_id | product_name | unit_price |
-- +------------+--------------+------------+
-- | 1          | S8           | 1000       |
-- | 2          | G4           | 800        |
-- | 3          | iPhone       | 1400       |
-- +------------+--------------+------------+
-- Sales table:
-- +-----------+------------+----------+------------+----------+-------+
-- | seller_id | product_id | buyer_id | sale_date  | quantity | price |
-- +-----------+------------+----------+------------+----------+-------+
-- | 1         | 1          | 1        | 2019-01-21 | 2        | 2000  |
-- | 1         | 2          | 2        | 2019-02-17 | 1        | 800   |
-- | 2         | 1          | 3        | 2019-06-02 | 1        | 800   |
-- | 3         | 3          | 3        | 2019-05-13 | 2        | 2800  |
-- +-----------+------------+----------+------------+----------+-------+
-- Output:
-- +-------------+
-- | buyer_id    |
-- +-------------+
-- | 1           |
-- +-------------+
-- Explanation: The buyer with id 1 bought an S8 but did not buy an iPhone. The buyer with id 3 bought both.
--
-- Write your MySQL query statement below
SELECT
    DISTINCT s.buyer_id AS buyer_id
FROM
    Sales AS s,
    Product AS p
WHERE
    s.product_id = p.product_id AND
    p.product_name = "S8" AND
    buyer_id NOT IN ( -- 购买了 iPhone  的用户
        SELECT
            s1.buyer_id
        FROM
            Sales AS s1,
            Product AS p1
        WHERE
            p1.product_id = s1.product_id AND
            p1.product_name = "iPhone"
    )
ORDER BY
    buyer_id