-- 2994. Friday Purchases II
-- Table: Purchases
-- +---------------+------+
-- | Column Name   | Type |
-- +---------------+------+
-- | user_id       | int  |
-- | purchase_date | date |
-- | amount_spend  | int  |
-- +---------------+------+
-- (user_id, purchase_date, amount_spend) is the primary key (combination of columns with unique values) for this table.
-- purchase_date will range from November 1, 2023, to November 30, 2023, inclusive of both dates.
-- Each row contains user id, purchase date, and amount spend.
-- Write a solution to calculate the total spending by users on each Friday of every week in November 2023. If there are no purchases on a particular Friday of a week, it will be considered as 0.

-- Return the result table ordered by week of month in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Purchases table:
-- +---------+---------------+--------------+
-- | user_id | purchase_date | amount_spend |
-- +---------+---------------+--------------+
-- | 11      | 2023-11-07    | 1126         |
-- | 15      | 2023-11-30    | 7473         |
-- | 17      | 2023-11-14    | 2414         |
-- | 12      | 2023-11-24    | 9692         |
-- | 8       | 2023-11-03    | 5117         |
-- | 1       | 2023-11-16    | 5241         |
-- | 10      | 2023-11-12    | 8266         |
-- | 13      | 2023-11-24    | 12000        |
-- +---------+---------------+--------------+
-- Output: 
-- +---------------+---------------+--------------+
-- | week_of_month | purchase_date | total_amount |
-- +---------------+---------------+--------------+
-- | 1             | 2023-11-03    | 5117         |
-- | 2             | 2023-11-10    | 0            |
-- | 3             | 2023-11-17    | 0            |
-- | 4             | 2023-11-24    | 21692        |
-- +---------------+---------------+--------------+ 
-- Explanation: 
-- - During the first week of November 2023, transactions amounting to $5,117 occurred on Friday, 2023-11-03.
-- - For the second week of November 2023, there were no transactions on Friday, 2023-11-10, resulting in a value of 0 in the output table for that day.
-- - Similarly, during the third week of November 2023, there were no transactions on Friday, 2023-11-17, reflected as 0 in the output table for that specific day.
-- - In the fourth week of November 2023, two transactions took place on Friday, 2023-11-24, amounting to $12,000 and $9,692 respectively, summing up to a total of $21,692.
-- Output table is ordered by week_of_month in ascending order.

-- Create Table if Not Exists Purchases( user_id int, purchase_date date, amount_spend int)
-- Truncate table Purchases
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('11', '2023-11-07', '1126')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('15', '2023-11-30', '7473')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('17', '2023-11-14', '2414')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('12', '2023-11-24', '9692')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('8', '2023-11-03', '5117')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('1', '2023-11-16', '5241')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('10', '2023-11-12', '8266')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('13', '2023-11-24', '12000')

WITH RECURSIVE cte(date) AS (
    SELECT '2023-11-01'
    UNION ALL
    SELECT DATE_ADD(date, INTERVAL 1 DAY)
    FROM cte
    WHERE date <= '2023-11-29'
), t AS (
    SELECT *
    FROM cte
    LEFT JOIN purchases p
    ON cte.date = p.purchase_date
)
SELECT 
    CEILING(DAY(date) / 7) AS week_of_month,
    date AS purchase_date,
    IFNULL(SUM(amount_spend), 0) AS total_amount
FROM 
    t
WHERE 
    DAYNAME(date) = 'Friday'
GROUP BY 
    date
ORDER BY 
    week_of_month

-- 使用递归得到2023年每一个周五的日期
WITH RECURSIVE dates AS ( 
    SELECT '2023-11-01' AS date
    UNION ALL
    SELECT 
        adddate(date,  1 )
    FROM 
        dates
    WHERE 
        DATE_FORMAT(DATE_ADD(date, INTERVAL 1 DAY), '%Y-%m') = '2023-11'
), 
t AS (
    SELECT 
        row_number() over(order by date) as week_of_month, 
        date
    FROM 
        dates
    WHERE 
        DAYOFWEEK(date) = 6
),
r AS (
    SELECT 
        WEEK(purchase_date, 1) - WEEK('2023-11-01',1) + 1 AS week_of_month,
        purchase_date,
        SUM(amount_spend) AS total_amount
    FROM
        Purchases
    WHERE 
        WEEKDAY(purchase_date) = 4
    GROUP BY
        purchase_date
)

SELECT 
    t.week_of_month,
    t.date AS purchase_date
    IFNULL(r.total_amount, 0) AS total_amount
FROM
    t 
LEFT JOIN 
    r 
ON 
    t.week_of_month = r.week_of_month


