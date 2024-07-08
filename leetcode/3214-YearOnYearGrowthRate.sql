-- 3214. Year on Year Growth Rate
-- Table: user_transactions
-- +------------------+----------+
-- | Column Name      | Type     | 
-- +------------------+----------+
-- | transaction_id   | integer  |
-- | product_id       | integer  |
-- | spend            | decimal  |
-- | transaction_date | datetime |
-- +------------------+----------+
-- The transaction_id column uniquely identifies each row in this table.
-- Each row of this table contains the transaction ID, product ID, the spend amount, and the transaction date.
-- Write a solution to calculate the year-on-year growth rate for the total spend for each product.

-- The result table should include the following columns:
--     year: The year of the transaction.
--     product_id: The ID of the product.
--     curr_year_spend: The total spend for the current year.
--     prev_year_spend: The total spend for the previous year.
--     yoy_rate: The year-on-year growth rate percentage, rounded to 2 decimal places.

-- Return the result table ordered by product_id,year in ascending order.
-- The result format is in the following example.

-- Example:
-- Input:
-- user_transactions table:
-- +----------------+------------+---------+---------------------+
-- | transaction_id | product_id | spend   | transaction_date    |
-- +----------------+------------+---------+---------------------+
-- | 1341           | 123424     | 1500.60 | 2019-12-31 12:00:00 |
-- | 1423           | 123424     | 1000.20 | 2020-12-31 12:00:00 |
-- | 1623           | 123424     | 1246.44 | 2021-12-31 12:00:00 |
-- | 1322           | 123424     | 2145.32 | 2022-12-31 12:00:00 |
-- +----------------+------------+---------+---------------------+
-- Output:
-- +------+------------+----------------+----------------+----------+
-- | year | product_id | curr_year_spend| prev_year_spend| yoy_rate |
-- +------+------------+----------------+----------------+----------+
-- | 2019 | 123424     | 1500.60        | NULL           | NULL     |
-- | 2020 | 123424     | 1000.20        | 1500.60        | -33.35   |
-- | 2021 | 123424     | 1246.44        | 1000.20        | 24.62    |
-- | 2022 | 123424     | 2145.32        | 1246.44        | 72.12    |
-- +------+------------+----------------+----------------+----------+
-- Explanation:
-- For product ID 123424:
-- In 2019:
-- Current year's spend is 1500.60
-- No previous year's spend recorded
-- YoY growth rate: NULL
-- In 2020:
-- Current year's spend is 1000.20
-- Previous year's spend is 1500.60
-- YoY growth rate: ((1000.20 - 1500.60) / 1500.60) * 100 = -33.35%
-- In 2021:
-- Current year's spend is 1246.44
-- Previous year's spend is 1000.20
-- YoY growth rate: ((1246.44 - 1000.20) / 1000.20) * 100 = 24.62%
-- In 2022:
-- Current year's spend is 2145.32
-- Previous year's spend is 1246.44
-- YoY growth rate: ((2145.32 - 1246.44) / 1246.44) * 100 = 72.12%
-- Note: Output table is ordered by product_id and year in ascending order.

-- Create Table if not exists user_transactions( transaction_id int, product_id int, spend decimal(10,2), 
-- transaction_date datetime)
-- Truncate table user_transactions
-- insert into user_transactions (transaction_id, product_id, spend, transaction_date) values ('1341', '123424', '1500.6', '2019-12-31 12:00:00')
-- insert into user_transactions (transaction_id, product_id, spend, transaction_date) values ('1423', '123424', '1000.2', '2020-12-31 12:00:00')
-- insert into user_transactions (transaction_id, product_id, spend, transaction_date) values ('1623', '123424', '1246.44', '2021-12-31 12:00:00')
-- insert into user_transactions (transaction_id, product_id, spend, transaction_date) values ('1322', '123424', '2145.32', '2022-12-31 12:00:00')

WITH t AS ( -- 按年统计消费总额
    SELECT 
        product_id,
        EXTRACT(YEAR FROM transaction_date) AS year,
        SUM(spend) AS total_spend
    FROM 
        user_transactions
    GROUP BY 
        product_id, 
        EXTRACT(YEAR FROM transaction_date)
),
tt AS ( -- 和上一年消费的
    SELECT 
        ys.product_id,
        ys.year,
        ys.total_spend AS curr_year_spend,
        py.total_spend AS prev_year_spend
    FROM 
        t ys
    LEFT JOIN 
        t py
    ON 
        ys.product_id = py.product_id
        AND ys.year = py.year + 1
)
SELECT
    year,
    product_id,
    curr_year_spend,
    prev_year_spend,
    ROUND(
        CASE 
            WHEN prev_year_spend = 0 THEN NULL
            ELSE ROUND(((curr_year_spend - prev_year_spend) / prev_year_spend) * 100, 2)
        END, 
        2
    ) AS yoy_rate
FROM 
    tt
ORDER BY 
    product_id,
    year