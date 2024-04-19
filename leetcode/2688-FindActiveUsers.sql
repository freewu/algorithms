-- 2688. Find Active Users
-- Table: Users
-- +-------------+----------+ 
-- | Column Name | Type     | 
-- +-------------+----------+ 
-- | user_id     | int      | 
-- | item        | varchar  |
-- | created_at  | datetime |
-- | amount      | int      |
-- +-------------+----------+
-- This table may contain duplicate records. 
-- Each row includes the user ID, the purchased item, the date of purchase, and the purchase amount.
-- Write a solution to identify active users. An active user is a user that has made a second purchase within 7 days of any other of their purchases.

-- For example, if the ending date is May 31, 2023. So any date between May 31, 2023, and June 7, 2023 (inclusive) would be considered "within 7 days" of May 31, 2023.
-- Return a list of user_id which denotes the list of active users in any order.
-- The result format is in the following example.

-- Example 1:
-- Input:
-- Users table:
-- +---------+-------------------+------------+--------+ 
-- | user_id | item              | created_at | amount |  
-- +---------+-------------------+------------+--------+
-- | 5       | Smart Crock Pot   | 2021-09-18 | 698882 |
-- | 6       | Smart Lock        | 2021-09-14 | 11487  |
-- | 6       | Smart Thermostat  | 2021-09-10 | 674762 |
-- | 8       | Smart Light Strip | 2021-09-29 | 630773 |
-- | 4       | Smart Cat Feeder  | 2021-09-02 | 693545 |
-- | 4       | Smart Bed         | 2021-09-13 | 170249 |
-- +---------+-------------------+------------+--------+ 
-- Output:
-- +---------+
-- | user_id | 
-- +---------+
-- | 6       | 
-- +---------+
-- Explanation: 
-- - User with user_id 5 has only one transaction, so he is not an active user.
-- - User with user_id 6 has two transaction his first transaction was on 2021-09-10 and second transation was on 2021-09-14. The distance between the first and second transactions date is <= 7 days. So he is an active user. 
-- - User with user_id 8 has only one transaction, so he is not an active user.  
-- - User with user_id 4 has two transaction his first transaction was on 2021-09-02 and second transation was on 2021-09-13. The distance between the first and second transactions date is > 7 days. So he is not an active user. 

-- Create table If Not Exists Users (user_id int, item varchar(100),created_at date,amount int)
-- Truncate table Users
-- insert into Users (user_id, item, created_at, amount) values ('5', 'Smart Crock Pot', '2021-09-18', '698882')
-- insert into Users (user_id, item, created_at, amount) values ('6', 'Smart Lock', '2021-09-14', '11487')
-- insert into Users (user_id, item, created_at, amount) values ('6', 'Smart Thermostat', '2021-09-10', '674762')
-- insert into Users (user_id, item, created_at, amount) values ('8', 'Smart Light Strip', '2021-09-29', '630773')
-- insert into Users (user_id, item, created_at, amount) values ('4', 'Smart Cat Feeder', '2021-09-02', '693545')
-- insert into Users (user_id, item, created_at, amount) values ('4', 'Smart Bed', '2021-09-13', '170249')

-- # DATE_ADD(date, INTERVAL expr type)
--     date 参数是合法的日期表达式。expr 参数是我们希望添加的时间间隔。
--     type 参数可以是下列值：  
-- Type值	说明
-- MICROSECOND	返回时间或日期时间表达式expr的微秒，这个数字范围为 0 到 999999
-- SECOND	返回时间秒值，范围为0〜59。
-- MINUTE	返回时间的分钟，范围为0至59。
-- HOUR	返回时间的小时部分。返回值的范围为0至23的小时值。然而，TIME值的范围实际上要大得多，所以HOUR可以返回大于23的值。
-- DAY	返回给定日期的月份的日期部分。
-- WEEK	返回日期的星期数
-- MONTH	返回日期的月份，取值范围为0〜12。
-- QUARTER	返回年份日期的季度值，范围为1〜4
-- YEAR	返回年份日期，范围为1000〜9999或0
-- SECOND_MICROSECOND	 
-- MINUTE_MICROSECOND	 
-- MINUTE_SECOND	 
-- HOUR_MICROSECOND	 
-- HOUR_SECOND	 
-- HOUR_MINUTE	 
-- DAY_MICROSECOND	 
-- DAY_SECOND	 
-- DAY_MINUTE	 
-- DAY_HOUR	 
-- YEAR_MONTH	 

-- Write your MySQL query statement below
-- SELECT 
--     DISTINCT a.user_id AS user_id 
-- FROM
--     Users AS a 
-- LEFT JOIN
--     Users AS b
-- ON 
--     a.user_id = b.user_id AND 
--     a.created_at != b.created_at AND -- 不能当天
--     b.created_at BETWEEN a.created_at AND DATE_ADD(a.created_at, INTERVAL 7 DAY)
-- WHERE  
--     b.user_id IS NOT NULL

WITH t AS (
    SELECT
        *,
        ROW_NUMBER() OVER(PARTITION BY user_id ORDER BY created_at) AS "rk" -- 每条件记录编号上号
    FROM
        Users
)

-- SELECT * FROM t
SELECT 
    DISTINCT a.user_id AS user_id 
FROM
    t AS a 
LEFT JOIN
    t AS b
ON 
    a.user_id = b.user_id AND 
    a.rk != b.rk AND -- 本条数据
    b.created_at BETWEEN a.created_at AND DATE_ADD(a.created_at, INTERVAL 7 DAY)
WHERE  
    b.user_id IS NOT NULL