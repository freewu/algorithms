-- 1777. Product's Price for Each Store
-- Table: Products
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | product_id  | int     |
-- | store       | enum    |
-- | price       | int     |
-- +-------------+---------+
-- In SQL, (product_id, store) is the primary key for this table.
-- store is a category of type ('store1', 'store2', 'store3') where each represents the store this product is available at.
-- price is the price of the product at this store.

-- Find the price of each product in each store.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Products table:
-- +-------------+--------+-------+
-- | product_id  | store  | price |
-- +-------------+--------+-------+
-- | 0           | store1 | 95    |
-- | 0           | store3 | 105   |
-- | 0           | store2 | 100   |
-- | 1           | store1 | 70    |
-- | 1           | store3 | 80    |
-- +-------------+--------+-------+
-- Output: 
-- +-------------+--------+--------+--------+
-- | product_id  | store1 | store2 | store3 |
-- +-------------+--------+--------+--------+
-- | 0           | 95     | 100    | 105    |
-- | 1           | 70     | null   | 80     |
-- +-------------+--------+--------+--------+
-- Explanation: 
-- Product 0 price's are 95 for store1, 100 for store2 and, 105 for store3.
-- Product 1 price's are 70 for store1, 80 for store3 and, it's not sold in store2.

-- Create table If Not Exists Products (product_id int, store ENUM('store1', 'store2', 'store3'), price int)
-- Truncate table Products
-- insert into Products (product_id, store, price) values ('0', 'store1', '95')
-- insert into Products (product_id, store, price) values ('0', 'store3', '105')
-- insert into Products (product_id, store, price) values ('0', 'store2', '100')
-- insert into Products (product_id, store, price) values ('1', 'store1', '70')
-- insert into Products (product_id, store, price) values ('1', 'store3', '80')

-- IF
SELECT
    product_id,
    SUM(IF(store = 'store1', price, null)) AS store1,
    SUM(IF(store = 'store2', price, null)) AS store2,
    SUM(IF(store = 'store3', price, null)) AS store3
FROM
    Products
GROUP BY
    product_id

-- CASE WHEN
SELECT
    product_id,
    MAX(CASE WHEN store = 'store1' THEN price ELSE NULL END) AS store1,
    MAX(CASE WHEN store = 'store2' THEN price ELSE NULL END) AS store2,
    MAX(CASE WHEN store = 'store3' THEN price ELSE NULL END) AS store3
FROM
    Products
GROUP BY
    product_id;

-- LEFT JOIN
SELECT
    distinct p.product_id AS product_id, 
    a.price AS store1, 
    b.price AS store2, 
    c.price AS store3
FROM
    Products p
LEFT JOIN 
    (
        SELECT * FROM Products WHERE store = 'store1'
    ) AS a 
ON p.product_id = a.product_id
LEFT JOIN 
    (
        SELECT * FROM Products WHERE store = 'store2'
    ) AS b 
ON p.product_id = b.product_id
LEFT JOIN 
    (
        SELECT * FROM Products WHERE store = 'store3'
    ) AS c 
ON p.product_id = c.product_id