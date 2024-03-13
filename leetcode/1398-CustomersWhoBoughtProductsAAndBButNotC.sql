-- 1398. Customers Who Bought Products A and B but Not C
-- Table: Customers
-- +---------------------+---------+
-- | Column Name         | Type    |
-- +---------------------+---------+
-- | customer_id         | int     |
-- | customer_name       | varchar |
-- +---------------------+---------+
-- customer_id is the column with unique values for this table.
-- customer_name is the name of the customer.
 
-- Table: Orders
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | order_id      | int     |
-- | customer_id   | int     |
-- | product_name  | varchar |
-- +---------------+---------+
-- order_id is the column with unique values for this table.
-- customer_id is the id of the customer who bought the product "product_name".
 
-- Write a solution to report the customer_id and customer_name of customers who bought products "A", "B" but did not buy the product "C" since we want to recommend them to purchase this product.
-- Return the result table ordered by customer_id.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Customers table:
-- +-------------+---------------+
-- | customer_id | customer_name |
-- +-------------+---------------+
-- | 1           | Daniel        |
-- | 2           | Diana         |
-- | 3           | Elizabeth     |
-- | 4           | Jhon          |
-- +-------------+---------------+
-- Orders table:
-- +------------+--------------+---------------+
-- | order_id   | customer_id  | product_name  |
-- +------------+--------------+---------------+
-- | 10         |     1        |     A         |
-- | 20         |     1        |     B         |
-- | 30         |     1        |     D         |
-- | 40         |     1        |     C         |
-- | 50         |     2        |     A         |
-- | 60         |     3        |     A         |
-- | 70         |     3        |     B         |
-- | 80         |     3        |     D         |
-- | 90         |     4        |     C         |
-- +------------+--------------+---------------+
-- Output: 
-- +-------------+---------------+
-- | customer_id | customer_name |
-- +-------------+---------------+
-- | 3           | Elizabeth     |
-- +-------------+---------------+
-- Explanation: Only the customer_id with id 3 bought the product A and B but not the product C.

-- Create table If Not Exists Customers (customer_id int, customer_name varchar(30))
-- Create table If Not Exists Orders (order_id int, customer_id int, product_name varchar(30))
-- Truncate table Customers
-- insert into Customers (customer_id, customer_name) values ('1', 'Daniel')
-- insert into Customers (customer_id, customer_name) values ('2', 'Diana')
-- insert into Customers (customer_id, customer_name) values ('3', 'Elizabeth')
-- insert into Customers (customer_id, customer_name) values ('4', 'Jhon')
-- Truncate table Orders
-- insert into Orders (order_id, customer_id, product_name) values ('10', '1', 'A')
-- insert into Orders (order_id, customer_id, product_name) values ('20', '1', 'B')
-- insert into Orders (order_id, customer_id, product_name) values ('30', '1', 'D')
-- insert into Orders (order_id, customer_id, product_name) values ('40', '1', 'C')
-- insert into Orders (order_id, customer_id, product_name) values ('50', '2', 'A')
-- insert into Orders (order_id, customer_id, product_name) values ('60', '3', 'A')
-- insert into Orders (order_id, customer_id, product_name) values ('70', '3', 'B')
-- insert into Orders (order_id, customer_id, product_name) values ('80', '3', 'D')
-- insert into Orders (order_id, customer_id, product_name) values ('90', '4', 'C')

SELECT
    *
FROM
    Customers
WHERE
    customer_id IN (
        --  买了 A 和 B 的用户
        SELECT 
            a.customer_id  
        FROM
        (
            SELECT 
                customer_id
            FROM    
                Orders 
            WHERE
                product_name  = 'A'
        ) AS a,
        (
            SELECT 
                customer_id
            FROM    
                Orders 
            WHERE
                product_name  = 'B'
        ) AS b
        WHERE 
            a.customer_id = b.customer_id 
    )
    AND 
    customer_id NOT IN (
        -- 买了 C 的用户
        SELECT 
            DISTINCT customer_id  
        FROM
            Orders 
        WHERE
            product_name = "C"
    )

-- best solution
SELECT 
    a.customer_id AS customer_id,
    b.customer_name AS customer_name
FROM 
    Orders AS a 
JOIN 
    Customers as b 
USING(customer_id) 
GROUP BY 
    customer_id
HAVING 
    SUM( a.product_name = 'C' ) = 0 AND -- 没有买 C
    SUM( a.product_name = 'A' ) * SUM(a.product_name = 'B') != 0 -- 买了 A 和 B 如果没买就为 0 用 * 太牛逼了