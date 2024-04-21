-- 2986. Find Third Transaction
-- Table: Transactions
-- +------------------+----------+
-- | Column Name      | Type     |
-- +------------------+----------+
-- | user_id          | int      |
-- | spend            | decimal  |
-- | transaction_date | datetime |
-- +------------------+----------+
-- (user_id, transaction_date) is column of unique values for this table.
-- This table contains user_id, spend, and transaction_date.
-- Write a solution to find the third transaction (if they have at least three transactions) of every user, where the spending on the preceding two transactions is lower than the spending on the third transaction.

-- Return the result table by user_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Transactions table:
-- +---------+--------+---------------------+
-- | user_id | spend  | transaction_date    | 
-- +---------+--------+---------------------+
-- | 1       | 65.56  | 2023-11-18 13:49:42 | 
-- | 1       | 96.0   | 2023-11-30 02:47:26 |     
-- | 1       | 7.44   | 2023-11-02 12:15:23 | 
-- | 1       | 49.78  | 2023-11-12 00:13:46 | 
-- | 2       | 40.89  | 2023-11-21 04:39:15 |  
-- | 2       | 100.44 | 2023-11-20 07:39:34 | 
-- | 3       | 37.33  | 2023-11-03 06:22:02 | 
-- | 3       | 13.89  | 2023-11-11 16:00:14 | 
-- | 3       | 7.0    | 2023-11-29 22:32:36 | 
-- +---------+--------+---------------------+
-- Output
-- +---------+-------------------------+------------------------+
-- | user_id | third_transaction_spend | third_transaction_date | 
-- +---------+-------------------------+------------------------+
-- | 1       | 65.56                   | 2023-11-18 13:49:42    |  
-- +---------+-------------------------+------------------------+
-- Explanation
-- - For user_id 1, their third transaction occurred on 2023-11-18 at 13:49:42 with an amount of $65.56, surpassing the expenditures of the previous two transactions which were $7.44 on 2023-11-02 at 12:15:23 and $49.78 on 2023-11-12 at 00:13:46. Thus, this third transaction will be included in the output table.
-- - user_id 2 only has a total of 2 transactions, so there isn't a third transaction to consider.
-- - For user_id 3, the amount of $7.0 for their third transaction is less than that of the preceding two transactions, so it won't be included.
-- Output table is ordered by user_id in ascending order.

-- Create Table if Not Exists Transactions (user_id int, spend decimal(5,2), transaction_date datetime) 
-- Truncate table Transactions
-- insert into Transactions (user_id, spend, transaction_date) values ('1', '65.56', '2023-11-18 13:49:42')
-- insert into Transactions (user_id, spend, transaction_date) values ('1', '96.0', '2023-11-30 02:47:26')
-- insert into Transactions (user_id, spend, transaction_date) values ('1', '7.44', '2023-11-02 12:15:23')
-- insert into Transactions (user_id, spend, transaction_date) values ('1', '49.78', '2023-11-12 00:13:46')
-- insert into Transactions (user_id, spend, transaction_date) values ('2', '40.89', '2023-11-21 04:39:15')
-- insert into Transactions (user_id, spend, transaction_date) values ('2', '100.44', '2023-11-20 07:39:34')
-- insert into Transactions (user_id, spend, transaction_date) values ('3', '37.33', '2023-11-03 06:22:02')
-- insert into Transactions (user_id, spend, transaction_date) values ('3', '13.89', '2023-11-11 16:00:14')
-- insert into Transactions (user_id, spend, transaction_date) values ('3', '7.0', '2023-11-29 22:32:36')

# Write your MySQL query statement below
WITH t AS (
    SELECT 
        *,
        RANK() OVER(PARTITION BY user_id ORDER BY transaction_date) AS rk
    FROM
        Transactions
)

SELECT 
    a.user_id AS user_id,
    a.spend AS third_transaction_spend,
    a.transaction_date AS third_transaction_date
FROM 
    t AS a,
    t AS b,
    t AS c 
WHERE
    a.user_id = b.user_id AND a.user_id = c.user_id AND 
    a.spend > b.spend AND a.spend > c.spend AND --  前两笔交易 的花费 低于 第三笔交易的花费。
    a.rk = 3 AND b.rk = 2 AND c.rk = 1
ORDER BY 
    a.user_id -- 按 升序 user_id 排序的结果表


-- SELECT
--     a.user_id AS user_id,
--     a.spend AS third_transaction_spend,
--     a.transaction_date AS third_transaction_date
-- FROM 
--     (-- 取第三笔数据
--         SELECT 
--             *
--         FROM 
--             t 
--         WHERE 
--             rk = 3
--     ) AS a 
-- JOIN 
--     (-- 每个客户前两笔
--         SELECT 
--             user_id,
--             SUM(spend) AS spend
--         FROM 
--             t 
--         WHERE
--             t.rk IN (1,2)
--         GROUP BY
--             user_id
--     ) AS b 
-- ON
--     a.user_id = b.user_id AND 
--     a.spend >= b.spend -- 满足 前两笔交易 的花费 低于 第三笔交易的花费
-- ORDER BY 
--     a.user_id

-- -- -- SELECT * FROM t
-- SELECT 
--     a.user_id AS user_id,
--     a.spend AS third_transaction_spend,
--     -- a.transaction_date AS third_transaction_date,
--     a.rk,
--     -- -- b.*
--     -- a.*,
--     b.spend,
--     b.rk
-- FROM 
--     t AS a 
-- JOIN 
--     t AS b 
-- ON 
--     a.user_id = b.user_id AND 
--     a.rk = 3 AND --  第三笔交易 （如果他们有至少三笔交易）
--     a.rk > b.rk
--     --a.rk > b.rk AND  a.spend > b.spend  --  前两笔交易 的花费 低于 第三笔交易的花费
--     -- b.rk = 1 -- 一般能匹配两条取和1条即可
-- ORDER BY 
--     user_id 


select 
	user_id,
    spend as third_transaction_spend,
    transaction_date as third_transaction_date
from 
(
    select 
        *,
        row_number() over(partition by user_id order by transaction_date) as rn,
        lag(spend,1,0) over(partition by user_id order by transaction_date) as last_spend,
        lag(spend,2,0) over(partition by user_id order by transaction_date) as first_spend
    from 
        Transactions
)   t1
where 
    rn = 3 and 
    spend >last_spend and spend > first_spend
order by 
    user_id 