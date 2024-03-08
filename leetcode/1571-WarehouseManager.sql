-- 1571. Warehouse Manager
-- Table: Warehouse
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | name         | varchar |
-- | product_id   | int     |
-- | units        | int     |
-- +--------------+---------+
-- (name, product_id) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the information of the products in each warehouse.

-- Table: Products
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | product_id    | int     |
-- | product_name  | varchar |
-- | Width         | int     |
-- | Length        | int     |
-- | Height        | int     |
-- +---------------+---------+
-- product_id is the primary key (column with unique values) for this table.
-- Each row of this table contains information about the product dimensions (Width, Lenght, and Height) in feets of each product.
-- Write a solution to report the number of cubic feet of volume the inventory occupies in each warehouse.
-- Return the result table in any order.
-- The query result format is in the following example.

-- Example 1:
-- Input: 
-- Warehouse table:
-- +------------+--------------+-------------+
-- | name       | product_id   | units       |
-- +------------+--------------+-------------+
-- | LCHouse1   | 1            | 1           |
-- | LCHouse1   | 2            | 10          |
-- | LCHouse1   | 3            | 5           |
-- | LCHouse2   | 1            | 2           |
-- | LCHouse2   | 2            | 2           |
-- | LCHouse3   | 4            | 1           |
-- +------------+--------------+-------------+
-- Products table:
-- +------------+--------------+------------+----------+-----------+
-- | product_id | product_name | Width      | Length   | Height    |
-- +------------+--------------+------------+----------+-----------+
-- | 1          | LC-TV        | 5          | 50       | 40        |
-- | 2          | LC-KeyChain  | 5          | 5        | 5         |
-- | 3          | LC-Phone     | 2          | 10       | 10        |
-- | 4          | LC-T-Shirt   | 4          | 10       | 20        |
-- +------------+--------------+------------+----------+-----------+
-- Output: 
-- +----------------+------------+
-- | warehouse_name | volume     | 
-- +----------------+------------+
-- | LCHouse1       | 12250      | 
-- | LCHouse2       | 20250      |
-- | LCHouse3       | 800        |
-- +----------------+------------+
-- Explanation: 
-- Volume of product_id = 1 (LC-TV), 5x50x40 = 10000
-- Volume of product_id = 2 (LC-KeyChain), 5x5x5 = 125 
-- Volume of product_id = 3 (LC-Phone), 2x10x10 = 200
-- Volume of product_id = 4 (LC-T-Shirt), 4x10x20 = 800
-- LCHouse1: 1 unit of LC-TV + 10 units of LC-KeyChain + 5 units of LC-Phone.
--           Total volume: 1*10000 + 10*125  + 5*200 = 12250 cubic feet
-- LCHouse2: 2 units of LC-TV + 2 units of LC-KeyChain.
--           Total volume: 2*10000 + 2*125 = 20250 cubic feet
-- LCHouse3: 1 unit of LC-T-Shirt.
--           Total volume: 1*800 = 800 cubic feet.

-- Create table If Not Exists Warehouse (name varchar(50), product_id int, units int)
-- Create table If Not Exists Products (product_id int, product_name varchar(50), Width int,Length int,Height int)
-- Truncate table Warehouse
-- insert into Warehouse (name, product_id, units) values ('LCHouse1', '1', '1')
-- insert into Warehouse (name, product_id, units) values ('LCHouse1', '2', '10')
-- insert into Warehouse (name, product_id, units) values ('LCHouse1', '3', '5')
-- insert into Warehouse (name, product_id, units) values ('LCHouse2', '1', '2')
-- insert into Warehouse (name, product_id, units) values ('LCHouse2', '2', '2')
-- insert into Warehouse (name, product_id, units) values ('LCHouse3', '4', '1')
-- Truncate table Products
-- insert into Products (product_id, product_name, Width, Length, Height) values ('1', 'LC-TV', '5', '50', '40')
-- insert into Products (product_id, product_name, Width, Length, Height) values ('2', 'LC-KeyChain', '5', '5', '5')
-- insert into Products (product_id, product_name, Width, Length, Height) values ('3', 'LC-Phone', '2', '10', '10')
-- insert into Products (product_id, product_name, Width, Length, Height) values ('4', 'LC-T-Shirt', '4', '10', '20')

-- Write your MySQL query statement below
SELECT
    w.name AS warehouse_name,
    SUM(w.units * p.Width * p.Length * p.Height) AS volume -- 计算货物体积
FROM 
    Warehouse AS w
LEFT JOIN
    Products AS p 
ON 
    w.product_id = p.product_id
GROUP BY
    1