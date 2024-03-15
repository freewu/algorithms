-- 1907. Count Salary Categories
-- Table: Accounts
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | account_id  | int  |
-- | income      | int  |
-- +-------------+------+
-- account_id is the primary key (column with unique values) for this table.
-- Each row contains information about the monthly income for one bank account.
 
-- Write a solution to calculate the number of bank accounts for each salary category. The salary categories are:
--      "Low Salary": All the salaries strictly less than $20000.
--      "Average Salary": All the salaries in the inclusive range [$20000, $50000].
--      "High Salary": All the salaries strictly greater than $50000.

-- The result table must contain all three categories. If there are no accounts in a category, return 0.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Accounts table:
-- +------------+--------+
-- | account_id | income |
-- +------------+--------+
-- | 3          | 108939 |
-- | 2          | 12747  |
-- | 8          | 87709  |
-- | 6          | 91796  |
-- +------------+--------+
-- Output: 
-- +----------------+----------------+
-- | category       | accounts_count |
-- +----------------+----------------+
-- | Low Salary     | 1              |
-- | Average Salary | 0              |
-- | High Salary    | 3              |
-- +----------------+----------------+
-- Explanation: 
-- Low Salary: Account 2.
-- Average Salary: No accounts.
-- High Salary: Accounts 3, 6, and 8.

-- Create table If Not Exists Accounts (account_id int, income int)
-- Truncate table Accounts
-- insert into Accounts (account_id, income) values ('3', '108939')
-- insert into Accounts (account_id, income) values ('2', '12747')
-- insert into Accounts (account_id, income) values ('8', '87709')
-- insert into Accounts (account_id, income) values ('6', '91796')

-- # Write your MySQL query statement below
-- SELECT 
--     category,
--     count(1) AS accounts_count 
-- FROM 
-- (
--     SELECT 
--     CASE 
--     WHEN income > 50000 THEN 'High Salary'
--     WHEN income <= 50000 AND income >= 20000 THEN 'Average Salary'
--     WHEN income < 20000 THEN 'Low Salary'
--     END AS category,
--     account_id 
-- FROM 
--     Accounts
-- ) AS a
-- GROUP BY 
--     category

-- UNION
( -- "Low Salary": All the salaries strictly less than $20000.
    SELECT 
        "Low Salary" AS category, 
        COUNT(*) AS accounts_count 
    FROM 
        accounts 
    WHERE income < 20000
)
UNION
( -- "Average Salary": All the salaries in the inclusive range [$20000, $50000].
    SELECT 
        "Average Salary" AS category, 
        COUNT(*) AS accounts_count 
    FROM    
        accounts 
    WHERE 
        income BETWEEN 20000 AND 50000
)
UNION
( -- "High Salary": All the salaries strictly greater than $50000.
    SELECT 
        "High Salary" AS category, 
        COUNT(*) AS accounts_count 
    FROM 
        accounts 
    WHERE 
        income > 50000
)

-- LEFT JOIN
SELECT 
    t_1.f_type AS category,
    IFNULL(t2.accounts_count,0) AS accounts_count
FROM 
(
    SELECT 'Low Salary' AS category UNION  
    SELECT 'Average Salary' AS  category UNION 
    SELECT 'High Salary' AS category
) AS c
LEFT JOIN 
(
    SELECT
        t.category,
        COUNT(DISTINCT account_id) AS accounts_count
    FROM
    (
        SELECT 
            account_id, 
            CASE -- 按薪水范围分类
                WHEN income < 20000 THEN 'Low Salary'
                WHEN income >= 20000 AND income <= 50000 THEN 'Average Salary'
                WHEN income > 50000 THEN 'High Salary'
            END AS category
        FROM 
            Accounts
    ) AS t
    GROUP BY
        t.category
) AS t2
ON 
    c.category = t2.category
