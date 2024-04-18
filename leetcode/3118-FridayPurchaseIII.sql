-- 3118. Friday Purchase III 
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
-- Each row contains user_id, purchase_date, and amount_spend.

-- Table: Users
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | user_id     | int  |
-- | membership  | enum |
-- +-------------+------+
-- user_id is the primary key for this table.
-- membership is an ENUM (category) type of ('Standard', 'Premium', 'VIP').
-- Each row of this table indicates the user_id, membership type.
-- Write a solution to calculate the total spending by Premium and VIP members on each Friday of every week in November 2023.  If there are no purchases on a particular Friday by Premium or VIP members, it should be considered as 0.

-- Return the result table ordered by week of the month,  and membership in ascending order.

-- The result format is in the following example.
 

-- Example:
-- Input:
-- Purchases table:
-- +---------+---------------+--------------+
-- | user_id | purchase_date | amount_spend |
-- +---------+---------------+--------------+
-- | 11      | 2023-11-03    | 1126         |
-- | 15      | 2023-11-10    | 7473         |
-- | 17      | 2023-11-17    | 2414         |
-- | 12      | 2023-11-24    | 9692         |
-- | 8       | 2023-11-24    | 5117         |
-- | 1       | 2023-11-24    | 5241         |
-- | 10      | 2023-11-22    | 8266         |
-- | 13      | 2023-11-21    | 12000        |
-- +---------+---------------+--------------+
-- Users table:
-- +---------+------------+
-- | user_id | membership |
-- +---------+------------+
-- | 11      | Premium    |
-- | 15      | VIP        |
-- | 17      | Standard   |
-- | 12      | VIP        |
-- | 8       | Premium    |
-- | 1       | VIP        |
-- | 10      | Standard   |
-- | 13      | Premium    |
-- +---------+------------+
-- Output:
-- +---------------+-------------+--------------+
-- | week_of_month | membership  | total_amount |
-- +---------------+-------------+--------------+
-- | 1             | Premium     | 1126         |
-- | 1             | VIP         | 0            |
-- | 2             | Premium     | 0            |
-- | 2             | VIP         | 7473         |
-- | 3             | Premium     | 0            |
-- | 3             | VIP         | 0            |
-- | 4             | Premium     | 5117         |
-- | 4             | VIP         | 14933        |
-- +---------------+-------------+--------------+    
-- Explanation:
-- During the first week of November 2023, a transaction occurred on Friday, 2023-11-03, by a Premium member amounting to $1,126. No transactions were made by VIP members on this day, resulting in a value of 0.
-- For the second week of November 2023, there was a transaction on Friday, 2023-11-10, and it was made by a VIP member, amounting to $7,473. Since there were no purchases by Premium members that Friday, the output shows 0 for Premium members.
-- Similarly, during the third week of November 2023, no transactions by Premium or VIP members occurred on Friday, 2023-11-17, which shows 0 for both categories in this week.
-- In the fourth week of November 2023, transactions occurred on Friday, 2023-11-24, involving one Premium member purchase of $5,117 and VIP member purchases totaling $14,933 ($9,692 from one and $5,241 from another).
-- Note: The output table is ordered by week_of_month and membership in ascending order.

-- Create Table if Not Exists Purchases( user_id int, purchase_date date, amount_spend int)
-- Create Table if Not Exists Users (user_id int, membership enum('Standard', 'Premium', 'VIP'))
-- Truncate table Purchases
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('11', '2023-11-03', '1126')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('15', '2023-11-10', '7473')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('17', '2023-11-17', '2414')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('12', '2023-11-24', '9692')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('8', '2023-11-24', '5117')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('1', '2023-11-24', '5241')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('10', '2023-11-22', '8266')
-- insert into Purchases (user_id, purchase_date, amount_spend) values ('13', '2023-11-21', '12000')
-- Truncate table Users
-- insert into Users (user_id, membership) values ('11', 'Premium')
-- insert into Users (user_id, membership) values ('15', 'VIP')
-- insert into Users (user_id, membership) values ('17', 'Standard')
-- insert into Users (user_id, membership) values ('12', 'VIP')
-- insert into Users (user_id, membership) values ('8', 'Premium')
-- insert into Users (user_id, membership) values ('1', 'VIP')
-- insert into Users (user_id, membership) values ('10', 'Standard')
-- insert into Users (user_id, membership) values ('13', 'Premium')


# Write your MySQL query statement below
WITH t AS -- 2023-11 Premium / VIP 用户 周五 消费明细总表
(
    SELECT
        p.*,
        (WEEKOFYEAR(purchase_date) - WEEKOFYEAR("2023-11-01")) AS week,
        u.membership 
    FROM 
        Purchases AS p
    LEFT JOIN
        Users AS u 
    ON 
        p.user_id = u.user_id 
    WHERE
        DATE_FORMAT(purchase_date,"%Y%m") = "202311" AND -- 只取 2023年11月的
        membership IN ("Premium","VIP") AND -- 只取 Premium / VIP 用户
        DAYOFWEEK(purchase_date) = 6 -- 只取周5的消费
),
r AS (
    SELECT 
        a.*,
        b.*
    FROM
        (
            SELECT 1 AS week_of_month 
            UNION ALL
            SELECT 2 AS week_of_month 
            UNION ALL
            SELECT 3 AS week_of_month 
            UNION ALL
            SELECT 4 AS week_of_month 
        ) AS a, 
        (
            SELECT "Premium" AS membership   
            UNION ALL
            SELECT "VIP" AS membership   
        ) AS b
) 

-- SELECT * FROM t
-- SELECT * FROM r
-- SELECT 
--     t.week + 1 AS week_of_month,
--     membership,
--     SUM(amount_spend) AS total_amount
-- FROM 
--     t
-- GROUP BY
--     t.week, membership

SELECT 
    r.*,
    IFNULL(x.total_amount, 0) AS total_amount
FROM 
    r
LEFT JOIN
    (
        SELECT 
            t.week + 1 AS week_of_month,
            membership,
            SUM(amount_spend) AS total_amount
        FROM 
            t
        GROUP BY
            t.week, membership
    ) AS x 
ON 
    x.week_of_month = r.week_of_month  AND
    x.membership = r.membership
ORDER BY
    week_of_month ASC, membership