-- 1972. First and Last Call On the Same Day
-- Table: Calls
-- +--------------+----------+
-- | Column Name  | Type     |
-- +--------------+----------+
-- | caller_id    | int      |
-- | recipient_id | int      |
-- | call_time    | datetime |
-- +--------------+----------+
-- (caller_id, recipient_id, call_time) is the primary key (combination of columns with unique values) for this table.
-- Each row contains information about the time of a phone call between caller_id and recipient_id.
 
-- Write a solution to report the IDs of the users whose first and last calls on any day were with the same person. 
-- Calls are counted regardless of being the caller or the recipient.

-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Calls table:
-- +-----------+--------------+---------------------+
-- | caller_id | recipient_id | call_time           |
-- +-----------+--------------+---------------------+
-- | 8         | 4            | 2021-08-24 17:46:07 |
-- | 4         | 8            | 2021-08-24 19:57:13 |
-- | 5         | 1            | 2021-08-11 05:28:44 |
-- | 8         | 3            | 2021-08-17 04:04:15 |
-- | 11        | 3            | 2021-08-17 13:07:00 |
-- | 8         | 11           | 2021-08-17 22:22:22 |
-- +-----------+--------------+---------------------+
-- Output: 
-- +---------+
-- | user_id |
-- +---------+
-- | 1       |
-- | 4       |
-- | 5       |
-- | 8       |
-- +---------+
-- Explanation: 
-- On 2021-08-24, the first and last call of this day for user 8 was with user 4. User 8 should be included in the answer.
-- Similarly, user 4 on 2021-08-24 had their first and last call with user 8. User 4 should be included in the answer.
-- On 2021-08-11, user 1 and 5 had a call. This call was the only call for both of them on this day. Since this call is the first and last call of the day for both of them, they should both be included in the answer.

-- Create table If Not Exists Calls (caller_id int, recipient_id int, call_time datetime)
-- Truncate table Calls
-- insert into Calls (caller_id, recipient_id, call_time) values ('8', '4', '2021-08-24 17:46:07')
-- insert into Calls (caller_id, recipient_id, call_time) values ('4', '8', '2021-08-24 19:57:13')
-- insert into Calls (caller_id, recipient_id, call_time) values ('5', '1', '2021-08-11 05:28:44')
-- insert into Calls (caller_id, recipient_id, call_time) values ('8', '3', '2021-08-17 04:04:15')
-- insert into Calls (caller_id, recipient_id, call_time) values ('11', '3', '2021-08-17 13:07:00')
-- insert into Calls (caller_id, recipient_id, call_time) values ('8', '11', '2021-08-17 22:22:22')

-- Write your MySQL query statement below
WITH c AS ( 
    -- 所有用户电话记录
    SELECT caller_id AS u1,recipient_id AS u2, call_time FROM Calls
    UNION ALL
    SELECT recipient_id AS u1,caller_id AS u2, call_time FROM Calls
),
t AS (
    -- 每一天每个用户的第一个电话和最后一个电话编号
    SELECT 
        u1,
        u2,
        DATE(call_time) AS d,
        ROW_NUMBER() OVER(PARTITION BY DATE(call_time), u1 ORDER BY call_time) AS rn1, -- 早 -> 晚 1为最早
        ROW_NUMBER() OVER(PARTITION BY DATE(call_time), u1 ORDER BY call_time DESC) AS rn2 -- 晚 -> 早 1为最晚
    FROM 
        c
)
SELECT 
    DISTINCT a.u1 as user_id 
FROM 
    (SELECT * FROM t WHERE rn1 = 1) AS a -- 最早
JOIN 
    (SELECT * FROM t WHERE rn2 = 1) AS b -- 最晚
ON 
    a.u1 = b.u1 AND a.u2 = b.u2