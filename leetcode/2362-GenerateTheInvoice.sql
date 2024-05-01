-- 2362. Generate the Invoice
-- Table: Products
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | product_id  | int  |
-- | price       | int  |
-- +-------------+------+
-- product_id contains unique values.
-- Each row in this table shows the ID of a product and the price of one unit.
 
-- Table: Purchases
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | invoice_id  | int  |
-- | product_id  | int  |
-- | quantity    | int  |
-- +-------------+------+
-- (invoice_id, product_id) is the primary key (combination of columns with unique values) for this table.
-- Each row in this table shows the quantity ordered from one product in an invoice. 
 
-- Write a solution to show the details of the invoice with the highest price. If two or more invoices have the same price, return the details of the one with the smallest invoice_id.

-- Return the result table in any order.
-- The result format is shown in the following example.

-- Example 1:
-- Input: 
-- Products table:
-- +------------+-------+
-- | product_id | price |
-- +------------+-------+
-- | 1          | 100   |
-- | 2          | 200   |
-- +------------+-------+
-- Purchases table:
-- +------------+------------+----------+
-- | invoice_id | product_id | quantity |
-- +------------+------------+----------+
-- | 1          | 1          | 2        |
-- | 3          | 2          | 1        |
-- | 2          | 2          | 3        |
-- | 2          | 1          | 4        |
-- | 4          | 1          | 10       |
-- +------------+------------+----------+
-- Output: 
-- +------------+----------+-------+
-- | product_id | quantity | price |
-- +------------+----------+-------+
-- | 2          | 3        | 600   |
-- | 1          | 4        | 400   |
-- +------------+----------+-------+
-- Explanation: 
-- Invoice 1: price = (2 * 100) = $200
-- Invoice 2: price = (4 * 100) + (3 * 200) = $1000
-- Invoice 3: price = (1 * 200) = $200
-- Invoice 4: price = (10 * 100) = $1000
-- The highest price is $1000, and the invoices with the highest prices are 2 and 4. We return the details of the one with the smallest ID, which is invoice 2.

-- Create table If Not Exists Products (product_id int, price int)
-- Create table If Not Exists Purchases (invoice_id int, product_id int, quantity int)
-- Truncate table Products
-- insert into Products (product_id, price) values ('1', '100')
-- insert into Products (product_id, price) values ('2', '200')
-- Truncate table Purchases
-- insert into Purchases (invoice_id, product_id, quantity) values ('1', '1', '2')
-- insert into Purchases (invoice_id, product_id, quantity) values ('3', '2', '1')
-- insert into Purchases (invoice_id, product_id, quantity) values ('2', '2', '3')
-- insert into Purchases (invoice_id, product_id, quantity) values ('2', '1', '4')
-- insert into Purchases (invoice_id, product_id, quantity) values ('4', '1', '10')

WITH t AS ( -- 获取最大消费金额的发票
    SELECT 
        c.invoice_id,
        SUM(c.quantity * p.price) AS amount 
    FROM 
        Purchases AS c 
    LEFT JOIN 
        Products p 
    ON 
        c.product_id = p.product_id 
    GROUP BY 
        c.invoice_id
    ORDER BY 
        amount DESC, c.invoice_id  -- 如果两个或多个发票具有相同的价格，则返回 invoice_id 最小的发票的详细信息
    LIMIT 1
)

--  消费明细
SELECT 
    p.product_id AS product_id,
    a.quantity,
    a.quantity * p.price AS price
FROM 
    (-- 消费最高的发票
        SELECT 
            * 
        FROM 
            Purchases 
        WHERE 
            invoice_id IN (SELECT invoice_id FROM t)
    ) AS a 
LEFT JOIN 
    Products p 
ON 
    a.product_id = p.product_id
