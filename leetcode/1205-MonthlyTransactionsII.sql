-- 1205. Monthly Transactions II
-- Table: Transactions
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | id             | int     |
-- | country        | varchar |
-- | state          | enum    |
-- | amount         | int     |
-- | trans_date     | date    |
-- +----------------+---------+
-- id is the column of unique values of this table.
-- The table has information about incoming transactions.
-- The state column is an ENUM (category) of type ["approved", "declined"].

-- Table: Chargebacks
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | trans_id       | int     |
-- | trans_date     | date    |
-- +----------------+---------+
-- Chargebacks contains basic information regarding incoming chargebacks from some transactions placed in Transactions table.
-- trans_id is a foreign key (reference column) to the id column of Transactions table.
-- Each chargeback corresponds to a transaction made previously even if they were not approved.

-- Write a solution to find for each month and country: 
--     the number of approved transactions and their total amount, the number of chargebacks, and their total amount.
-- Note: In your solution, given the month and country, ignore rows with all zeros.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Transactions table:
-- +-----+---------+----------+--------+------------+
-- | id  | country | state    | amount | trans_date |
-- +-----+---------+----------+--------+------------+
-- | 101 | US      | approved | 1000   | 2019-05-18 |
-- | 102 | US      | declined | 2000   | 2019-05-19 |
-- | 103 | US      | approved | 3000   | 2019-06-10 |
-- | 104 | US      | declined | 4000   | 2019-06-13 |
-- | 105 | US      | approved | 5000   | 2019-06-15 |
-- +-----+---------+----------+--------+------------+
-- Chargebacks table:
-- +----------+------------+
-- | trans_id | trans_date |
-- +----------+------------+
-- | 102      | 2019-05-29 |
-- | 101      | 2019-06-30 |
-- | 105      | 2019-09-18 |
-- +----------+------------+
-- Output: 
-- +---------+---------+----------------+-----------------+------------------+-------------------+
-- | month   | country | approved_count | approved_amount | chargeback_count | chargeback_amount |
-- +---------+---------+----------------+-----------------+------------------+-------------------+
-- | 2019-05 | US      | 1              | 1000            | 1                | 2000              |
-- | 2019-06 | US      | 2              | 8000            | 1                | 1000              |
-- | 2019-09 | US      | 0              | 0               | 1                | 5000              |
-- +---------+---------+----------------+-----------------+------------------+-------------------+

-- Create table If Not Exists Transactions (id int, country varchar(4), state enum('approved', 'declined'), amount int, trans_date date)
-- Create table If Not Exists Chargebacks (trans_id int, trans_date date)
-- Truncate table Transactions
-- insert into Transactions (id, country, state, amount, trans_date) values ('101', 'US', 'approved', '1000', '2019-05-18')
-- insert into Transactions (id, country, state, amount, trans_date) values ('102', 'US', 'declined', '2000', '2019-05-19')
-- insert into Transactions (id, country, state, amount, trans_date) values ('103', 'US', 'approved', '3000', '2019-06-10')
-- insert into Transactions (id, country, state, amount, trans_date) values ('104', 'US', 'declined', '4000', '2019-06-13')
-- insert into Transactions (id, country, state, amount, trans_date) values ('105', 'US', 'approved', '5000', '2019-06-15')
-- Truncate table Chargebacks
-- insert into Chargebacks (trans_id, trans_date) values ('102', '2019-05-29')
-- insert into Chargebacks (trans_id, trans_date) values ('101', '2019-06-30')
-- insert into Chargebacks (trans_id, trans_date) values ('105', '2019-09-18')

SELECT 
    month,
    country, 
    SUM(IF(approved_amount > 0, 1, 0)) AS approved_count, 
    SUM(approved_amount) AS approved_amount,
    SUM(IF(chargeback_amount > 0, 1, 0)) AS chargeback_count,
    SUM(chargeback_amount) AS chargeback_amount
FROM
(
    (-- 已批准交易的数量及其总金额
        SELECT 
            DATE_FORMAT(t.trans_date,'%Y-%m') AS month,
            t.country AS country, 
            t.amount AS approved_amount,
            0 AS chargeback_amount
        FROM Transactions AS t
        WHERE state = 'approved'
    )
    UNION ALL
    ( -- 退单的数量及其总金额
        SELECT
            DATE_FORMAT(c.trans_date,'%Y-%m') AS month, -- 要使用 退单 的日期
            t.country AS country, 
            0 AS approved_amount,
            t.amount AS chargeback_amount
        FROM Chargebacks AS c
        LEFT JOIN Transactions AS t
        ON t.id = c.trans_id
    )
) AS r
GROUP BY 
    month, country