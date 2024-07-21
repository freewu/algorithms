-- 3230. Customer Purchasing Behavior Analysis
-- Table: Transactions
-- +------------------+---------+
-- | Column Name      | Type    |
-- +------------------+---------+
-- | transaction_id   | int     |
-- | customer_id      | int     |
-- | product_id       | int     |
-- | transaction_date | date    |
-- | amount           | decimal |
-- +------------------+---------+
-- transaction_id is the unique identifier for this table.
-- Each row of this table contains information about a transaction, including the customer ID, product ID, date, and amount spent.

-- Table: Products
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | product_id  | int     |
-- | category    | varchar |
-- | price       | decimal |
-- +-------------+---------+
-- product_id is the unique identifier for this table.
-- Each row of this table contains information about a product, including its category and price.
-- Write a solution to analyze customer purchasing behavior. For each customer, calculate:

-- The total amount spent.
-- The number of transactions.
-- The number of unique product categories purchased.
-- The average amount spent. 
-- The most frequently purchased product category (if there is a tie, choose the one with the most recent transaction).
-- A loyalty score defined as: (Number of transactions * 10) + (Total amount spent / 100).
-- Round total_amount, avg_transaction_amount, and loyalty_score to 2 decimal places.

-- Return the result table ordered by loyalty_score in descending order, then by customer_id in ascending order.
-- The query result format is in the following example.

-- Example:
-- Input:
-- Transactions table:
-- +----------------+-------------+------------+------------------+--------+
-- | transaction_id | customer_id | product_id | transaction_date | amount |
-- +----------------+-------------+------------+------------------+--------+
-- | 1              | 101         | 1          | 2023-01-01       | 100.00 |
-- | 2              | 101         | 2          | 2023-01-15       | 150.00 |
-- | 3              | 102         | 1          | 2023-01-01       | 100.00 |
-- | 4              | 102         | 3          | 2023-01-22       | 200.00 |
-- | 5              | 101         | 3          | 2023-02-10       | 200.00 |
-- +----------------+-------------+------------+------------------+--------+
-- Products table:
-- +------------+----------+--------+
-- | product_id | category | price  |
-- +------------+----------+--------+
-- | 1          | A        | 100.00 |
-- | 2          | B        | 150.00 |
-- | 3          | C        | 200.00 |
-- +------------+----------+--------+
-- Output:
-- +-------------+--------------+-------------------+-------------------+------------------------+--------------+---------------+
-- | customer_id | total_amount | transaction_count | unique_categories | avg_transaction_amount | top_category | loyalty_score |
-- +-------------+--------------+-------------------+-------------------+------------------------+--------------+---------------+
-- | 101         | 450.00       | 3                 | 3                 | 150.00                 | C            | 34.50         |
-- | 102         | 300.00       | 2                 | 2                 | 150.00                 | C            | 23.00         |
-- +-------------+--------------+-------------------+-------------------+------------------------+--------------+---------------+
-- Explanation:
-- For customer 101:
-- Total amount spent: 100.00 + 150.00 + 200.00 = 450.00
-- Number of transactions: 3
-- Unique categories: A, B, C (3 categories)
-- Average transaction amount: 450.00 / 3 = 150.00
-- Top category: C (Customer 101 made 1 purchase each in categories A, B, and C. Since the count is the same for all categories, we choose the most recent transaction, which is category C on 2023-02-10)
-- Loyalty score: (3 * 10) + (450.00 / 100) = 34.50
-- For customer 102:
-- Total amount spent: 100.00 + 200.00 = 300.00
-- Number of transactions: 2
-- Unique categories: A, C (2 categories)
-- Average transaction amount: 300.00 / 2 = 150.00
-- Top category: C (Customer 102 made 1 purchase each in categories A and C. Since the count is the same for both categories, we choose the most recent transaction, which is category C on 2023-01-22)
-- Loyalty score: (2 * 10) + (300.00 / 100) = 23.00
-- Note: The output is ordered by loyalty_score in descending order, then by customer_id in ascending order.

-- CREATE TABLE if not exists Transactions (
--     transaction_id INT,
--     customer_id INT,
--     product_id INT,
--     transaction_date DATE,
--     amount DECIMAL(10, 2)
-- )
-- CREATE TABLE if not exists Products (
--     product_id INT ,
--     category VARCHAR(255),
--     price DECIMAL(10, 2)
-- )

-- Truncate table Transactions
-- insert into Transactions (transaction_id, customer_id, product_id, transaction_date, amount) values ('1', '101', '1', '2023-01-01', '100.0')
-- insert into Transactions (transaction_id, customer_id, product_id, transaction_date, amount) values ('2', '101', '2', '2023-01-15', '150.0')
-- insert into Transactions (transaction_id, customer_id, product_id, transaction_date, amount) values ('3', '102', '1', '2023-01-01', '100.0')
-- insert into Transactions (transaction_id, customer_id, product_id, transaction_date, amount) values ('4', '102', '3', '2023-01-22', '200.0')
-- insert into Transactions (transaction_id, customer_id, product_id, transaction_date, amount) values ('5', '101', '3', '2023-02-10', '200.0')
-- Truncate table Products
-- insert into Products (product_id, category, price) values ('1', 'A', '100.0')
-- insert into Products (product_id, category, price) values ('2', 'B', '150.0')
-- insert into Products (product_id, category, price) values ('3', 'C', '200.0')

--  Write your MySQL query statement below
-- WITH t AS (
--     SELECT
--         a.transaction_id,
--         a.customer_id,
--         a.product_id,
--         a.transaction_date,
--         a.amount,
--         p.category,
--         p.price
--     FROM
--         Transactions AS a
--     LEFT JOIN
--         Products AS p
--     ON 
--         a.product_id  = p.product_id 
-- )

-- SELECT * FROM t;
-- | transaction_id | customer_id | product_id | transaction_date | amount | category | price |
-- | -------------- | ----------- | ---------- | ---------------- | ------ | -------- | ----- |
-- | 1              | 101         | 1          | 2023-01-01       | 100    | A        | 100   |
-- | 2              | 101         | 2          | 2023-01-15       | 150    | B        | 150   |
-- | 3              | 102         | 1          | 2023-01-01       | 100    | A        | 100   |
-- | 4              | 102         | 3          | 2023-01-22       | 200    | C        | 200   |
-- | 5              | 101         | 3          | 2023-02-10       | 200    | C        | 200   |

-- Write your MySQL query statement below
WITH t AS ( -- 合并两个表
    SELECT
        a.transaction_id,
        a.customer_id,
        a.product_id,
        a.transaction_date,
        a.amount,
        p.category,
        p.price
    FROM
        Transactions AS a
    LEFT JOIN
        Products AS p
    ON 
        a.product_id  = p.product_id 
),
`d` as ( -- 统计出最每个 customer_id 每个 category 的出现频次排名
    SELECT 
        customer_id,
        category,
        RANK() OVER(PARTITION BY customer_id ORDER BY num DESC,transaction_date DESC ) AS rk
    FROM
    (
        SELECT
            customer_id,
            category,
            COUNT(*) AS num,
            MAX(transaction_date) AS transaction_date
        FROM
            t 
        GROUP BY
            customer_id, category
    ) AS aa
)

-- SELECT * FROM d

-- SELECT
--     customer_id,
--     RANK(PARTITION BY category ORDER BY COUNT(*) DESC,transaction_date DESC) AS rk
-- FROM
--     t
-- GROUP BY
--     customer_id, category
-- ORDER BY 
--     COUNT(*), transaction_date DESC

SELECT
    *
FROM
(
    SELECT 
        t.customer_id,
        SUM(t.amount) AS total_amount,
        COUNT(t.transaction_id) AS transaction_count,
        COUNT(DISTINCT t.category) AS unique_categories,
        ROUND(SUM(t.amount) / COUNT(t.transaction_id), 2)  AS avg_transaction_amount,
        d.category AS top_category,
        ROUND(COUNT(t.transaction_id) * 10  + (SUM(t.amount) / 100), 2)  AS loyalty_score 
    FROM 
        t
    LEFT JOIN
        d 
    ON
        t.customer_id = d.customer_id AND d.rk = 1
    GROUP BY
        t.customer_id
) AS k
ORDER BY
    k.loyalty_score DESC, k.customer_id

