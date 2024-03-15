-- 1193. Monthly Transactions I
-- Table: Transactions
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | country       | varchar |
-- | state         | enum    |
-- | amount        | int     |
-- | trans_date    | date    |
-- +---------------+---------+
-- id is the primary key of this table.
-- The table has information about incoming transactions.
-- The state column is an enum of type ["approved", "declined"].
 
-- Write an SQL query to find for each month and country, the number of transactions and their total amount, the number of approved transactions and their total amount.
-- Return the result table in any order.
-- The query result format is in the following example.

-- Example 1:
-- Input: 
-- Transactions table:
-- +------+---------+----------+--------+------------+
-- | id   | country | state    | amount | trans_date |
-- +------+---------+----------+--------+------------+
-- | 121  | US      | approved | 1000   | 2018-12-18 |
-- | 122  | US      | declined | 2000   | 2018-12-19 |
-- | 123  | US      | approved | 2000   | 2019-01-01 |
-- | 124  | DE      | approved | 2000   | 2019-01-07 |
-- +------+---------+----------+--------+------------+
-- Output: 
-- +----------+---------+-------------+----------------+--------------------+-----------------------+
-- | month    | country | trans_count | approved_count | trans_total_amount | approved_total_amount |
-- +----------+---------+-------------+----------------+--------------------+-----------------------+
-- | 2018-12  | US      | 2           | 1              | 3000               | 1000                  |
-- | 2019-01  | US      | 1           | 1              | 2000               | 2000                  |
-- | 2019-01  | DE      | 1           | 1              | 2000               | 2000                  |
-- +----------+---------+-------------+----------------+--------------------+-----------------------+

-- Create table If Not Exists Transactions (id int, country varchar(4), state enum('approved', 'declined'), amount int, trans_date date)
-- Truncate table Transactions
-- insert into Transactions (id, country, state, amount, trans_date) values ('121', 'US', 'approved', '1000', '2018-12-18')
-- insert into Transactions (id, country, state, amount, trans_date) values ('122', 'US', 'declined', '2000', '2018-12-19')
-- insert into Transactions (id, country, state, amount, trans_date) values ('123', 'US', 'approved', '2000', '2019-01-01')
-- insert into Transactions (id, country, state, amount, trans_date) values ('124', 'DE', 'approved', '2000', '2019-01-07')

-- Write your MySQL query statement below
-- union
SELECT 
    month,
    country,
    SUM(declined_count + approved_count) AS trans_count,
    SUM(approved_count) AS approved_count,
    SUM(declined_amount + approved_amount) AS trans_total_amount,
    SUM(approved_amount) AS approved_total_amount
FROM
(
    (
        Select
            DATE_FORMAT(trans_date,'%Y-%m') AS month,
            country,
            count(1) AS declined_count,
            SUM(amount) AS declined_amount,
            0 AS approved_count,
            0 AS approved_amount
        FROM 
            Transactions 
        WHERE
            state = 'declined'
        GROUP BY
            DATE_FORMAT(trans_date,'%Y-%m'),country
    ) UNION  (
        Select
            DATE_FORMAT(trans_date,'%Y-%m') AS month,
            country,
            0 AS declined_count,
            0 AS declined_amount,
            count(1) AS approved_count,
            SUM(amount) AS approved_amount
        FROM 
            Transactions 
        WHERE
            state = 'approved'
        GROUP BY
            DATE_FORMAT(trans_date,'%Y-%m'),country
    )
) AS a
GROUP BY 
    month, country

-- 使用 IF 的处理
SELECT 
    DATE_FORMAT(trans_date, '%Y-%m') AS month,
    country,
    COUNT(*) AS trans_count,
    COUNT(IF(state = 'approved', 1, NULL)) AS approved_count,
    SUM(amount) AS trans_total_amount,
    SUM(IF(state = 'approved', amount, 0)) AS approved_total_amount
FROM 
    Transactions
GROUP BY 
    month, country