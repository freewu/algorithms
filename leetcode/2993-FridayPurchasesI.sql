-- 2993. Friday Purchases I
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
-- Write a solution to calculate the total spending by users on each Friday of every week in November 2023. Output only weeks that include at least one purchase on a Friday.

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
-- | 4             | 2023-11-24    | 21692        |
-- +---------------+---------------+--------------+ 
-- Explanation: 
-- - During the first week of November 2023, transactions amounting to $5,117 occurred on Friday, 2023-11-03.
-- - For the second week of November 2023, there were no transactions on Friday, 2023-11-10.
-- - Similarly, during the third week of November 2023, there were no transactions on Friday, 2023-11-17.
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

SELECT
    week_of_month + 1 AS week_of_month,
    purchase_date,
    SUM(amount_spend)  AS total_amount
FROM
(
    SELECT
        (WEEKOFYEAR(purchase_date) - WEEKOFYEAR("2023-11-01"))  AS week_of_month,
        purchase_date,
        amount_spend
    FROM 
        Purchases 
    WHERE
        -- DATE_FORMAT(purchase_date,"%Y%m") = "202311" AND -- 只取 2023年11月的
        purchase_date >= "2023-11-01" AND -- 只取 2023年11月的
        purchase_date <= "2023-11-30" AND
        DAYOFWEEK(purchase_date) = 6 -- 只取周5的消费
) AS t
GROUP BY
    week_of_month, purchase_date
ORDER BY 
    week_of_month ASC -- the result table ordered by week of month in ascending order.


-- best solution
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
ORDER BY 
    purchase_date