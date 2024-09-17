-- 3293. Calculate Product Final Price
-- Table: Products
-- +------------+---------+ 
-- | Column Name| Type    | 
-- +------------+---------+ 
-- | product_id | int     | 
-- | category   | varchar |
-- | price      | decimal |
-- +------------+---------+
-- product_id is the unique key for this table.
-- Each row includes the product's ID, its category, and its price.
-- Table: Discounts
-- +------------+---------+ 
-- | Column Name| Type    | 
-- +------------+---------+ 
-- | category   | varchar |
-- | discount   | int     |
-- +------------+---------+
-- category is the primary key for this table.
-- Each row contains a product category and the percentage discount applied to that category (values range from 0 to 100).
-- Write a solution to find the final price of each product after applying the category discount. If a product's category has no associated discount, its price remains unchanged.

-- Return the result table ordered by product_id in ascending order.
-- The result format is in the following example.

-- Example:
-- Input:
-- Products table:
-- +------------+-------------+-------+
-- | product_id | category    | price |
-- +------------+-------------+-------+
-- | 1          | Electronics | 1000  |
-- | 2          | Clothing    | 50    |
-- | 3          | Electronics | 1200  | 
-- | 4          | Home        | 500   |
-- +------------+-------------+-------+
-- Discounts table:
-- +------------+----------+
-- | category   | discount |
-- +------------+----------+
-- | Electronics| 10       |
-- | Clothing   | 20       |
-- +------------+----------+
-- Output:
-- +------------+------------+-------------+
-- | product_id | final_price| category    |
-- +------------+------------+-------------+
-- | 1          | 900        | Electronics |
-- | 2          | 40         | Clothing    |
-- | 3          | 1080       | Electronics |
-- | 4          | 500        | Home        |
-- +------------+------------+-------------+
-- Explanation:
-- For product 1, it belongs to the Electronics category which has a 10% discount, so the final price is 1000 - (10% of 1000) = 900.
-- For product 2, it belongs to the Clothing category which has a 20% discount, so the final price is 50 - (20% of 50) = 40.
-- For product 3, it belongs to the Electronics category and receives a 10% discount, so the final price is 1200 - (10% of 1200) = 1080.
-- For product 4, no discount is available for the Home category, so the final price remains 500.
-- Result table is ordered by product_id in ascending order.

-- Create table if not exists Products (product_id int, category varchar(50), price int)
-- Create table if not exists Discounts(category varchar(50), discount int)
-- Truncate table Products
-- insert into Products (product_id, category, price) values ('1', 'Electronics', '1000')
-- insert into Products (product_id, category, price) values ('2', 'Clothing', '50')
-- insert into Products (product_id, category, price) values ('3', 'Electronics', '1200')
-- insert into Products (product_id, category, price) values ('4', 'Home', '500')
-- Truncate table Discounts
-- insert into Discounts (category, discount) values ('Electronics', '10')
-- insert into Discounts (category, discount) values ('Clothing', '20')

-- SELECT
--     p.*,
--     d.*
-- FROM
--     Products AS p 
-- LEFT JOIN
--     Discounts AS d 
-- ON
--     p.category = d.category
-- ORDER BY 
--     p.product_id 

SELECT
    p.product_id AS  product_id,
    IF(
        d.discount = null,
        p.price,
        ROUND( (100 - d.discount) / 100 * p.price )
    ) AS final_price
    p.category AS category
FROM
    Products AS p 
LEFT JOIN
    Discounts AS d 
ON
    p.category = d.category
ORDER BY 
    p.product_id 