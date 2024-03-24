-- 1369. Get the Second Most Recent Activity
-- Table: UserActivity
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | username      | varchar |
-- | activity      | varchar |
-- | startDate     | Date    |
-- | endDate       | Date    |
-- +---------------+---------+
-- This table may contain duplicates rows.
-- This table contains information about the activity performed by each user in a period of time.
-- A person with username performed an activity from startDate to endDate.
 
-- Write a solution to show the second most recent activity of each user.
-- If the user only has one activity, return that one. A user cannot perform more than one activity at the same time.
-- Return the result table in any order.

-- The result format is in the following example.

-- Example 1:
-- Input: 
-- UserActivity table:
-- +------------+--------------+-------------+-------------+
-- | username   | activity     | startDate   | endDate     |
-- +------------+--------------+-------------+-------------+
-- | Alice      | Travel       | 2020-02-12  | 2020-02-20  |
-- | Alice      | Dancing      | 2020-02-21  | 2020-02-23  |
-- | Alice      | Travel       | 2020-02-24  | 2020-02-28  |
-- | Bob        | Travel       | 2020-02-11  | 2020-02-18  |
-- +------------+--------------+-------------+-------------+
-- Output: 
-- +------------+--------------+-------------+-------------+
-- | username   | activity     | startDate   | endDate     |
-- +------------+--------------+-------------+-------------+
-- | Alice      | Dancing      | 2020-02-21  | 2020-02-23  |
-- | Bob        | Travel       | 2020-02-11  | 2020-02-18  |
-- +------------+--------------+-------------+-------------+
-- Explanation: 
-- The most recent activity of Alice is Travel from 2020-02-24 to 2020-02-28, before that she was dancing from 2020-02-21 to 2020-02-23.
-- Bob only has one record, we just take that one.

-- Create table If Not Exists UserActivity (username varchar(30), activity varchar(30), startDate date, endDate date)
-- Truncate table UserActivity
-- insert into UserActivity (username, activity, startDate, endDate) values ('Alice', 'Travel', '2020-02-12', '2020-02-20')
-- insert into UserActivity (username, activity, startDate, endDate) values ('Alice', 'Dancing', '2020-02-21', '2020-02-23')
-- insert into UserActivity (username, activity, startDate, endDate) values ('Alice', 'Travel', '2020-02-24', '2020-02-28')
-- insert into UserActivity (username, activity, startDate, endDate) values ('Bob', 'Travel', '2020-02-11', '2020-02-18')

WITH t AS 
( -- 只参与了一次的
    SELECT 
        username
    FROM
        (-- 按 每个用户参加的每个活动排序
            SELECT 
                *,
                RANK() OVER(PARTITION BY username ORDER BY startDate DESC) AS rk
            FROM 
                UserActivity 
        ) AS a
    GROUP BY 
        username
    HAVING
        COUNT(*) = 1
)

-- SELECT * FROM t

SELECT 
    username,
    activity,
    startDate,
    endDate
FROM
    (-- 按 每个用户参加的每个活动排序
        SELECT 
            *,
            RANK() OVER(PARTITION BY username ORDER BY startDate DESC) AS rk
        FROM 
            UserActivity 
    ) AS a
WHERE
    rk = 2 OR -- 每一位用户 最近第二次 的活动
    (username IN (SELECT username FROM t) AND rk = 1) -- 如果用户仅有一次活动，返回该活动

