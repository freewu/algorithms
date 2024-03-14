-- 1831. Maximum Transaction Each Day
-- Table: Transactions
-- +----------------+----------+
-- | Column Name    | Type     |
-- +----------------+----------+
-- | transaction_id | int      |
-- | day            | datetime |
-- | amount         | int      |
-- +----------------+----------+
-- transaction_id is the column with unique values for this table.
-- Each row contains information about one transaction.
 
-- Write a solution to report the IDs of the transactions with the maximum amount on their respective day. If in one day there are multiple such transactions, return all of them.
-- Return the result table ordered by transaction_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Transactions table:
-- +----------------+--------------------+--------+
-- | transaction_id | day                | amount |
-- +----------------+--------------------+--------+
-- | 8              | 2021-4-3 15:57:28  | 57     |
-- | 9              | 2021-4-28 08:47:25 | 21     |
-- | 1              | 2021-4-29 13:28:30 | 58     |
-- | 5              | 2021-4-28 16:39:59 | 40     |
-- | 6              | 2021-4-29 23:39:28 | 58     |
-- +----------------+--------------------+--------+
-- Output: 
-- +----------------+
-- | transaction_id |
-- +----------------+
-- | 1              |
-- | 5              |
-- | 6              |
-- | 8              |
-- +----------------+
-- Explanation: 
-- "2021-4-3"  --> We have one transaction with ID 8, so we add 8 to the result table.
-- "2021-4-28" --> We have two transactions with IDs 5 and 9. The transaction with ID 5 has an amount of 40, while the transaction with ID 9 has an amount of 21. We only include the transaction with ID 5 as it has the maximum amount this day.
-- "2021-4-29" --> We have two transactions with IDs 1 and 6. Both transactions have the same amount of 58, so we include both in the result table.
-- We order the result table by transaction_id after collecting these IDs.
 
-- Follow up: Could you solve it without using the MAX() function?

-- use MAX
SELECT 
    t.transaction_id
FROM
    Transactions AS t,
    (-- 每天最大的金额
        SELECT 
            DATE_FORMAT(day,"%Y-%m-%d") AS date,
            MAX(amount) AS amount
        FROM 
            Transactions
        GROUP BY
            date
    ) as r
WHERE
    DATE_FORMAT(t.day,"%Y-%m-%d") = r.date AND 
    t.amount = r.amount
ORDER BY
    t.transaction_id ASC -- 结果根据 transaction_id 升序排列


-- SELECT 
--     DATE_FORMAT(day,"%Y-%m-%d") AS date,
--     amount,
--     RANK() OVER(PARTITION BY DATE_FORMAT(day,"%Y-%m-%d") ORDER BY amount DESC) AS rk
-- FROM
--     Transactions
-- | date       | amount | rk |
-- | ---------- | ------ | -- |
-- | 2021-04-03 | 57     | 1  |
-- | 2021-04-28 | 40     | 1  |
-- | 2021-04-28 | 21     | 2  |
-- | 2021-04-29 | 58     | 1  |
-- | 2021-04-29 | 58     | 1  |

-- use rank
SELECT 
    a.transaction_id
FROM 
    (
        SELECT
            transaction_id,
            RANK() OVER(PARTITION BY DATE_FORMAT(day,"%Y-%m-%d") ORDER BY amount DESC) AS rk
        FROM
            Transactions
    ) AS a
WHERE
    a.rk = 1 -- 只取第一的
ORDER BY
    a.transaction_id ASC -- 结果根据 transaction_id 升序排列
