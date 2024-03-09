-- 1809. Ad-Free Sessions
-- Table: Playback
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | session_id  | int  |
-- | customer_id | int  |
-- | start_time  | int  |
-- | end_time    | int  |
-- +-------------+------+
-- session_id is the column with unique values for this table.
-- customer_id is the ID of the customer watching this session.
-- The session runs during the inclusive interval between start_time and end_time.
-- It is guaranteed that start_time <= end_time and that two sessions for the same customer do not intersect.

-- Table: Ads
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | ad_id       | int  |
-- | customer_id | int  |
-- | timestamp   | int  |
-- +-------------+------+
-- ad_id is the column with unique values for this table.
-- customer_id is the ID of the customer viewing this ad.
-- timestamp is the moment of time at which the ad was shown.

-- Write a solution to report all the sessions that did not get shown any ads.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Playback table:
-- +------------+-------------+------------+----------+
-- | session_id | customer_id | start_time | end_time |
-- +------------+-------------+------------+----------+
-- | 1          | 1           | 1          | 5        |
-- | 2          | 1           | 15         | 23       |
-- | 3          | 2           | 10         | 12       |
-- | 4          | 2           | 17         | 28       |
-- | 5          | 2           | 2          | 8        |
-- +------------+-------------+------------+----------+
-- Ads table:
-- +-------+-------------+-----------+
-- | ad_id | customer_id | timestamp |
-- +-------+-------------+-----------+
-- | 1     | 1           | 5         |
-- | 2     | 2           | 17        |
-- | 3     | 2           | 20        |
-- +-------+-------------+-----------+
-- Output: 
-- +------------+
-- | session_id |
-- +------------+
-- | 2          |
-- | 3          |
-- | 5          |
-- +------------+
-- Explanation: 
-- The ad with ID 1 was shown to user 1 at time 5 while they were in session 1.
-- The ad with ID 2 was shown to user 2 at time 17 while they were in session 4.
-- The ad with ID 3 was shown to user 2 at time 20 while they were in session 4.
-- We can see that sessions 1 and 4 had at least one ad. Sessions 2, 3, and 5 did not have any ads, so we return them.

-- Create table If Not Exists Playback(session_id int,customer_id int,start_time int,end_time int)
-- Create table If Not Exists Ads (ad_id int, customer_id int, timestamp int)
-- Truncate table Playback
-- insert into Playback (session_id, customer_id, start_time, end_time) values ('1', '1', '1', '5')
-- insert into Playback (session_id, customer_id, start_time, end_time) values ('2', '1', '15', '23')
-- insert into Playback (session_id, customer_id, start_time, end_time) values ('3', '2', '10', '12')
-- insert into Playback (session_id, customer_id, start_time, end_time) values ('4', '2', '17', '28')
-- insert into Playback (session_id, customer_id, start_time, end_time) values ('5', '2', '2', '8')
-- Truncate table Ads
-- insert into Ads (ad_id, customer_id, timestamp) values ('1', '1', '5')
-- insert into Ads (ad_id, customer_id, timestamp) values ('2', '2', '17')
-- insert into Ads (ad_id, customer_id, timestamp) values ('3', '2', '20')

-- Write your MySQL query statement below
-- 超时了
SELECT 
    session_id
FROM
    Playback 
WHERE
    session_id NOT IN ( -- 出现过广告的剧集并被看过的
        SELECT
            p.session_id
        FROM
            Playback AS p 
        LEFT JOIN 
            Ads AS a 
        ON 
            p.customer_id = a.customer_id AND 
            a.timestamp BETWEEN p.start_time AND p.end_time 
        WHERE
            a.timestamp IS NOT NULL -- 不为空表示有广告记录存在
    )

SELECT 
    DISTINCT p.session_id AS session_id
FROM 
    Playback AS p
LEFT JOIN 
    Ads AS a 
ON 
    p.customer_id = a.customer_id AND 
    a.timestamp BETWEEN p.start_time AND p.end_time
WHERE 
    -- 因为有 p.customer_id = a.customer_id AND a.timestamp BETWEEN p.start_time AND p.end_time 限制，
    -- 但凡该 customer_id 有一次是有广告的都会被收进去，而没有广告的会被排除在外，
    -- 因此最终的联合表中字段 ad_id、timestamp 是null的数据都是一次广告都没有出现的
    a.timestamp IS NULL
