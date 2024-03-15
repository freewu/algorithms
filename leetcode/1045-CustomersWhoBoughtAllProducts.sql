-- 1045. Customers Who Bought All Products
-- Table: Customer
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | customer_id | int     |
-- | product_key | int     |
-- +-------------+---------+
-- There is no primary key for this table. It may contain duplicates.
-- product_key is a foreign key to Product table.
--  
-- Table: Product
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | product_key | int     |
-- +-------------+---------+
-- product_key is the primary key column for this table.
-- Write an SQL query to report the customer ids from the Customer table that bought all the products in the Product table.
-- Return the result table in any order.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- Customer table:
-- +-------------+-------------+
-- | customer_id | product_key |
-- +-------------+-------------+
-- | 1           | 5           |
-- | 2           | 6           |
-- | 3           | 5           |
-- | 3           | 6           |
-- | 1           | 6           |
-- +-------------+-------------+
-- Product table:
-- +-------------+
-- | product_key |
-- +-------------+
-- | 5           |
-- | 6           |
-- +-------------+
-- Output:
-- +-------------+
-- | customer_id |
-- +-------------+
-- | 1           |
-- | 3           |
-- +-------------+
-- Explanation:
-- The customers who bought all the products (5 and 6) are customers with IDs 1 and 3.

-- Create table If Not Exists Customer (customer_id int, product_key int)
-- Create table Product (product_key int)
-- Truncate table Customer
-- insert into Customer (customer_id, product_key) values ('1', '5')
-- insert into Customer (customer_id, product_key) values ('2', '6')
-- insert into Customer (customer_id, product_key) values ('3', '5')
-- insert into Customer (customer_id, product_key) values ('3', '6')
-- insert into Customer (customer_id, product_key) values ('1', '6')
-- Truncate table Product
-- insert into Product (product_key) values ('5')
-- insert into Product (product_key) values ('6')

-- Write your MySQL query statement below
SELECT
    customer_id
FROM
    (
        SELECT
            customer_id,
            COUNT(DISTINCT product_key) AS num
        FROM
            Customer
        GROUP BY
            customer_id
    ) AS a
WHERE
    a.num = (SELECT COUNT(*) FROM Product)

-- HAVING
SELECT 
    customer_id 
FROM 
    Customer 
GROUP BY 
    customer_id 
having 
    COUNT(DISTINCT product_key) = ( SELECT COUNT(*) FROM Product ) 
