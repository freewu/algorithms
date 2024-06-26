-- 3052. Maximize Items
-- Table: Inventory
-- +----------------+---------+ 
-- | Column Name    | Type    | 
-- +----------------+---------+ 
-- | item_id        | int     | 
-- | item_type      | varchar |
-- | item_category  | varchar |
-- | square_footage | decimal |
-- +----------------+---------+
-- item_id is the column of unique values for this table.
-- Each row includes item id, item type, item category and sqaure footage.
-- Leetcode warehouse wants to maximize the number of items it can stock in a 500,000 square feet warehouse. It wants to stock as many prime items as possible, and afterwards use the remaining square footage to stock the most number of non-prime items.

-- Write a solution to find the number of prime and non-prime items that can be stored in the 500,000 square feet warehouse. Output the item type with prime_eligible followed by not_prime and the maximum number of items that can be stocked.

-- Note:
--     Item count must be a whole number (integer).
--     If the count for the not_prime category is 0, you should output 0 for that particular category.

-- Return the result table ordered by item count in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Inventory table:
-- +---------+----------------+---------------+----------------+
-- | item_id | item_type      | item_category | square_footage | 
-- +---------+----------------+---------------+----------------+
-- | 1374    | prime_eligible | Watches       | 68.00          | 
-- | 4245    | not_prime      | Art           | 26.40          | 
-- | 5743    | prime_eligible | Software      | 325.00         | 
-- | 8543    | not_prime      | Clothing      | 64.50          |  
-- | 2556    | not_prime      | Shoes         | 15.00          |
-- | 2452    | prime_eligible | Scientific    | 85.00          |
-- | 3255    | not_prime      | Furniture     | 22.60          | 
-- | 1672    | prime_eligible | Beauty        | 8.50           |  
-- | 4256    | prime_eligible | Furniture     | 55.50          |
-- | 6325    | prime_eligible | Food          | 13.20          | 
-- +---------+----------------+---------------+----------------+
-- Output: 
-- +----------------+-------------+
-- | item_type      | item_count  | 
-- +----------------+-------------+
-- | prime_eligible | 5400        | 
-- | not_prime      | 8           | 
-- +----------------+-------------+
-- Explanation: 
-- - The prime-eligible category comprises a total of 6 items, amounting to a combined square footage of 555.20 (68 + 325 + 85 + 8.50 + 55.50 + 13.20). It is possible to store 900 combinations of these 6 items, totaling 5400 items and occupying 499,680 square footage.
-- - In the not_prime category, there are a total of 4 items with a combined square footage of 128.50. After deducting the storage used by prime-eligible items (500,000 - 499,680 = 320), there is room for 2 combinations of non-prime items, accommodating a total of 8 non-prime items within the available 320 square footage.
-- Output table is ordered by item count in descending order.

-- Create table If Not Exists Inventory ( item_id int, item_type varchar(50), item_category varchar(50), square_footage decimal(10,2))
-- Truncate table Inventory
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('1374', 'prime_eligible', 'Watches', '68.0')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('4245', 'not_prime', 'Art', '26.4')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('5743', 'prime_eligible', 'Software', '325.0')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('8543', 'not_prime', 'Clothing', '64.5')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('2556', 'not_prime', 'Shoes', '15.0')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('2452', 'prime_eligible', 'Scientific', '85.0')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('3255', 'not_prime', 'Furniture', '22.6')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('1672', 'prime_eligible', 'Beauty', '8.5')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('4256', 'prime_eligible', 'Furniture', '55.5')
-- insert into Inventory (item_id, item_type, item_category, square_footage) values ('6325', 'prime_eligible', 'Food', '13.2')

WITH t AS ( -- 汇总出不同 item_type 需要占用的总面积和数量
    SELECT
        item_type,
        SUM(square_footage) as square,
        COUNT(item_category) as cnt
    FROM 
        Inventory
    GROUP BY 
        item_type 
),
t1 AS ( --  主要商品 所占仓库面积 & 数量
    SELECT 
        item_type,
        FLOOR(500000 / square) * cnt as item_count,
        square * FLOOR(500000 / square) as total
    FROM 
        t
    WHERE
        item_type = "prime_eligible"
),
t2 AS ( --  非主要商品数量
    SELECT 
        item_type,
        FLOOR(((500000 - (SELECT total FROM t1)) / square)) * cnt AS item_count -- 除去 prime_eligible 占完的还余下的
    FROM 
        t
    WHERE
        item_type = "not_prime"
)


SELECT item_type, item_count FROM t1
UNION ALL
SELECT item_type, item_count FROM t2

-- best solution
WITH pr_b AS (  -- 主要商品
    SELECT item_id,square_footage FROM Inventory WHERE item_type = 'prime_eligible'
), 
no_b AS ( -- 非主要商品
    SELECT item_id,square_footage FROM Inventory WHERE item_type = 'not_prime'
) 

(
    SELECT 
        'prime_eligible' AS item_type,
        FLOOR(500000 / SUM(square_footage)) * COUNT(item_id) AS item_count -- 固定面积下主要最大数量
    FROM 
        pr_b
)
UNION ALL
(
    SELECT 
        'not_prime' AS item_type,
        FLOOR (
            ( 500000 -  (SELECT FLOOR(500000 / SUM(square_footage))  * SUM(square_footage) FROM pr_b)) / SUM(square_footage) -- 先减去主要商品的占地面积
        ) * COUNT(item_id) AS item_count 
    FROM 
        no_b
)
