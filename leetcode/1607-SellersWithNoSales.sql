-- 1607. Sellers With No Sales
-- Table: Customer
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | customer_id   | int     |
-- | customer_name | varchar |
-- +---------------+---------+
-- customer_id is the column with unique values for this table.
-- Each row of this table contains the information of each customer in the WebStore.

-- Table: Orders
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | order_id      | int     |
-- | sale_date     | date    |
-- | order_cost    | int     |
-- | customer_id   | int     |
-- | seller_id     | int     |
-- +---------------+---------+
-- order_id is the column with unique values for this table.
-- Each row of this table contains all orders made in the webstore.
-- sale_date is the date when the transaction was made between the customer (customer_id) and the seller (seller_id).
 
-- Table: Seller
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | seller_id     | int     |
-- | seller_name   | varchar |
-- +---------------+---------+
-- seller_id is the column with unique values for this table.
-- Each row of this table contains the information of each seller.

-- Write a solution to report the names of all sellers who did not make any sales in 2020.
-- Return the result table ordered by seller_name in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Customer table:
-- +--------------+---------------+
-- | customer_id  | customer_name |
-- +--------------+---------------+
-- | 101          | Alice         |
-- | 102          | Bob           |
-- | 103          | Charlie       |
-- +--------------+---------------+
-- Orders table:
-- +-------------+------------+--------------+-------------+-------------+
-- | order_id    | sale_date  | order_cost   | customer_id | seller_id   |
-- +-------------+------------+--------------+-------------+-------------+
-- | 1           | 2020-03-01 | 1500         | 101         | 1           |
-- | 2           | 2020-05-25 | 2400         | 102         | 2           |
-- | 3           | 2019-05-25 | 800          | 101         | 3           |
-- | 4           | 2020-09-13 | 1000         | 103         | 2           |
-- | 5           | 2019-02-11 | 700          | 101         | 2           |
-- +-------------+------------+--------------+-------------+-------------+
-- Seller table:
-- +-------------+-------------+
-- | seller_id   | seller_name |
-- +-------------+-------------+
-- | 1           | Daniel      |
-- | 2           | Elizabeth   |
-- | 3           | Frank       |
-- +-------------+-------------+
-- Output: 
-- +-------------+
-- | seller_name |
-- +-------------+
-- | Frank       |
-- +-------------+
-- Explanation: 
-- Daniel made 1 sale in March 2020.
-- Elizabeth made 2 sales in 2020 and 1 sale in 2019.
-- Frank made 1 sale in 2019 but no sales in 2020.

-- Create table If Not Exists Customer (customer_id int, customer_name varchar(20))
-- Create table If Not Exists Orders (order_id int, sale_date date, order_cost int, customer_id int, seller_id int)
-- Create table If Not Exists Seller (seller_id int, seller_name varchar(20))
-- Truncate table Customer
-- insert into Customer (customer_id, customer_name) values ('101', 'Alice')
-- insert into Customer (customer_id, customer_name) values ('102', 'Bob')
-- insert into Customer (customer_id, customer_name) values ('103', 'Charlie')
-- Truncate table Orders
-- insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('1', '2020-03-01', '1500', '101', '1')
-- insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('2', '2020-05-25', '2400', '102', '2')
-- insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('3', '2019-05-25', '800', '101', '3')
-- insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('4', '2020-09-13', '1000', '103', '2')
-- insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('5', '2019-02-11', '700', '101', '2')
-- Truncate table Seller
-- insert into Seller (seller_id, seller_name) values ('1', 'Daniel')
-- insert into Seller (seller_id, seller_name) values ('2', 'Elizabeth')
-- insert into Seller (seller_id, seller_name) values ('3', 'Frank')

SELECT 
    seller_name 
FROM 
    Seller 
WHERE 
    seller_id NOT IN (
        SELECT 
            DISTINCT seller_id
        FROM 
            Orders 
        WHERE
            DATE_FORMAT(sale_date,"%Y") = "2020"
    )
ORDER BY
    seller_name

-- best solution
SELECT 
    s.seller_name AS seller_name
FROM 
    Seller AS s
LEFT JOIN 
    Orders AS o
ON s.seller_id = o.seller_id AND YEAR(o.sale_date) = 2020
WHERE 
    o.order_id IS NULL
ORDER BY 
    s.seller_name ASC;