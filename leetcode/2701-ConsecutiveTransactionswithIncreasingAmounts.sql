-- 2701. Consecutive Transactions with Increasing Amounts
-- Table: Transactions
-- +------------------+------+
-- | Column Name      | Type |
-- +------------------+------+
-- | transaction_id   | int  |
-- | customer_id      | int  |
-- | transaction_date | date |
-- | amount           | int  |
-- +------------------+------+
-- transaction_id is the primary key of this table. 
-- Each row contains information about transactions that includes unique (customer_id, transaction_date) along with the corresponding customer_id and amount.  
-- Write an SQL query to find the customers who have made consecutive transactions with increasing amount for at least three consecutive days. Include the customer_id, start date of the consecutive transactions period and the end date of the consecutive transactions period. There can be multiple consecutive transactions by a customer.

-- Return the result table ordered by customer_id in ascending order.
-- The query result format is in the following example.

-- Example 1:
-- Input: 
-- Transactions table:
-- +----------------+-------------+------------------+--------+
-- | transaction_id | customer_id | transaction_date | amount |
-- +----------------+-------------+------------------+--------+
-- | 1              | 101         | 2023-05-01       | 100    |
-- | 2              | 101         | 2023-05-02       | 150    |
-- | 3              | 101         | 2023-05-03       | 200    |
-- | 4              | 102         | 2023-05-01       | 50     |
-- | 5              | 102         | 2023-05-03       | 100    |
-- | 6              | 102         | 2023-05-04       | 200    |
-- | 7              | 105         | 2023-05-01       | 100    |
-- | 8              | 105         | 2023-05-02       | 150    |
-- | 9              | 105         | 2023-05-03       | 200    |
-- | 10             | 105         | 2023-05-04       | 300    |
-- | 11             | 105         | 2023-05-12       | 250    |
-- | 12             | 105         | 2023-05-13       | 260    |
-- | 13             | 105         | 2023-05-14       | 270    |
-- +----------------+-------------+------------------+--------+
-- Output: 
-- +-------------+-------------------+-----------------+
-- | customer_id | consecutive_start | consecutive_end | 
-- +-------------+-------------------+-----------------+
-- | 101         |  2023-05-01       | 2023-05-03      | 
-- | 105         |  2023-05-01       | 2023-05-04      |
-- | 105         |  2023-05-12       | 2023-05-14      | 
-- +-------------+-------------------+-----------------+
-- Explanation: 
-- - customer_id 101 has made consecutive transactions with increasing amounts from May 1st, 2023, to May 3rd, 2023
-- - customer_id 102 does not have any consecutive transactions for at least 3 days. 
-- - customer_id 105 has two sets of consecutive transactions: from May 1st, 2023, to May 4th, 2023, and from May 12th, 2023, to May 14th, 2023. 
-- customer_id is sorted in ascending order.

-- Create table If Not Exists Transactions (transaction_id int, customer_id int, transaction_date date, amount int)
-- Truncate table Transactions
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('1', '101', '2023-05-01', '100')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('2', '101', '2023-05-02', '150')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('3', '101', '2023-05-03', '200')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('4', '102', '2023-05-01', '50')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('5', '102', '2023-05-03', '100')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('6', '102', '2023-05-04', '200')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('7', '105', '2023-05-01', '100')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('8', '105', '2023-05-02', '150')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('9', '105', '2023-05-03', '200')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('10', '105', '2023-05-04', '300')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('11', '105', '2023-05-12', '250')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('12', '105', '2023-05-13', '260')
-- insert into Transactions (transaction_id, customer_id, transaction_date, amount) values ('13', '105', '2023-05-14', '270')

WITH t AS (
    SELECT 
        customer_id,
        transaction_date,
        amount,
        increase_status
    FROM
        (
            SELECT 
                customer_id,
                transaction_date,
                amount,
                COUNT(transaction_date) OVER(PARTITION BY customer_id, DATE_SUB(transaction_date,interval rn day)) AS cnt,
                CASE 
                    WHEN amount > LAG(amount, 1) OVER(PARTITION BY customer_id,date_sub(transaction_date,interval rn day) order by transaction_date) THEN 1
                    ELSE 0
                END AS increase_status
            FROM
                ( -- 按每个客户每单消费时间编号
                    SELECT 
                        *,
                        ROW_NUMBER() OVER(PARTITION BY customer_id ORDER BY transaction_date) AS rn
                    FROM 
                        Transactions
                ) AS r
        ) AS a
    WHERE cnt >= 3
)

SELECT 
    customer_id,
    DATE_SUB(MIN(transaction_date), interval 1 day) AS consecutive_start,
    MAX(transaction_date) AS consecutive_end
FROM
    (
        SELECT 
            *,
            ROW_NUMBER() OVER(PARTITION BY customer_id ORDER BY transaction_date) AS rk
        FROM 
            t
        WHERE 
            increase_status = 1
    ) AS a
GROUP BY
    customer_id,DATE_SUB(transaction_date,interval rk day)
HAVING 
    COUNT(*) >= 2
ORDER BY
    customer_id

-- best solution
SELECT 
    customer_id, 
    MIN(transaction_date) AS consecutive_start, 
    MAX(transaction_date) AS consecutive_end
FROM
(
    SELECT 
        SUM(IF(amount > pre1, 0 ,1)) OVER(PARTITION BY customer_id ORDER BY transaction_date) AS flag_amount, 
        diff, 
        customer_id,
        transaction_date
    FROM 
    (
        SELECT 
            *,
            DATE_SUB(transaction_date, interval ROW_NUMBER() OVER(PARTITION BY customer_id ORDER BY transaction_date) day) AS diff,
            LAG(amount,1,0) OVER(PARTITION BY customer_id ORDER BY transaction_date) AS pre1
        FROM 
            Transactions
    ) AS t 
) AS a 
GROUP BY
    diff, flag_amount, customer_id
HAVING
    COUNT(*) > 2 
ORDER BY 
    customer_id
