-- 1164. Product Price at a Given Date
-- Table: Products

-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | product_id    | int     |
-- | new_price     | int     |
-- | change_date   | date    |
-- +---------------+---------+
-- (product_id, change_date) is the primary key (combination of columns with unique values) of this table.
-- Each row of this table indicates that the price of some product was changed to a new price at some date.
 

-- Write a solution to find the prices of all products on 2019-08-16. Assume the price of all products before any change is 10.

-- Return the result table in any order.

-- The result format is in the following example.

 

-- Example 1:

-- Input: 
-- Products table:
-- +------------+-----------+-------------+
-- | product_id | new_price | change_date |
-- +------------+-----------+-------------+
-- | 1          | 20        | 2019-08-14  |
-- | 2          | 50        | 2019-08-14  |
-- | 1          | 30        | 2019-08-15  |
-- | 1          | 35        | 2019-08-16  |
-- | 2          | 65        | 2019-08-17  |
-- | 3          | 20        | 2019-08-18  |
-- +------------+-----------+-------------+
-- Output: 
-- +------------+-------+
-- | product_id | price |
-- +------------+-------+
-- | 2          | 50    |
-- | 1          | 35    |
-- | 3          | 10    |
-- +------------+-------+

-- Write your MySQL query statement below
SELECT 
    p.product_id,
    IFNULL(d.new_price,10) AS price -- 处理 null 的问题
 FROM 
    (
        SELECT
            product_id
        FROM 
            Products 
        GROUP BY
            product_id 
    ) AS p -- 取唯一的  product_id
LEFT JOIN 
    (
       SELECT
            d1.product_id,
            d1.new_price
        FROM 
            (
                SELECT
                    product_id,
                    MAX(change_date) as change_date 
                FROM 
                    Products
                WHERE
                    change_date <= '2019-08-16'
                GROUP BY 
                    product_id
            ) AS d2 -- 取每个产品在 2019-08-16 内最近的一个更新日期
        LEFT JOIN 
            Products as d1
        ON d1.product_id = d2.product_id AND d1.change_date = d2.change_date  
    ) AS d -- 取每个产品在 2019-08-16 内最近的一个更新日期的价格
ON p.product_id = d.product_id

-- 简洁处理
SELECT
     distinct p1.product_id,
     coalesce((
         select 
            p2.new_price
        from
            Products p2
        where
            p2.product_id=p1.product_id AND p2.change_date <= '2019-08-16'
        order by
            p2.change_date DESC
        limit 1
     ),10) as price
from    
    Products p1
