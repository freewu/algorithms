-- 2720. Popularity Percentage
-- Table: Friends
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | user1       | int  |
-- | user2       | int  |
-- +-------------+------+
-- (user1, user2) is the primary key (combination of unique values) of this table.
-- Each row contains information about friendship where user1 and user2 are friends.
-- Write a solution to find the popularity percentage for each user on Meta/Facebook. The popularity percentage is defined as the total number of friends the user has divided by the total number of users on the platform, then converted into a percentage by multiplying by 100, rounded to 2 decimal places.

-- Return the result table ordered by user1 in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Friends table:
-- +-------+-------+
-- | user1 | user2 | 
-- +-------+-------+
-- | 2     | 1     | 
-- | 1     | 3     | 
-- | 4     | 1     | 
-- | 1     | 5     | 
-- | 1     | 6     |
-- | 2     | 6     | 
-- | 7     | 2     | 
-- | 8     | 3     | 
-- | 3     | 9     |  
-- +-------+-------+
-- Output: 
-- +-------+-----------------------+
-- | user1 | percentage_popularity |
-- +-------+-----------------------+
-- | 1     | 55.56                 |
-- | 2     | 33.33                 |
-- | 3     | 33.33                 |
-- | 4     | 11.11                 |
-- | 5     | 11.11                 |
-- | 6     | 22.22                 |
-- | 7     | 11.11                 |
-- | 8     | 11.11                 |
-- | 9     | 11.11                 |
-- +-------+-----------------------+
-- Explanation: 
-- There are total 9 users on the platform.
-- - User "1" has friendships with 2, 3, 4, 5 and 6. Therefore, the percentage popularity for user 1 would be calculated as (5/9) * 100 = 55.56.
-- - User "2" has friendships with 1, 6 and 7. Therefore, the percentage popularity for user 2 would be calculated as (3/9) * 100 = 33.33.
-- - User "3" has friendships with 1, 8 and 9. Therefore, the percentage popularity for user 3 would be calculated as (3/9) * 100 = 33.33.
-- - User "4" has friendships with 1. Therefore, the percentage popularity for user 4 would be calculated as (1/9) * 100 = 11.11.
-- - User "5" has friendships with 1. Therefore, the percentage popularity for user 5 would be calculated as (1/9) * 100 = 11.11.
-- - User "6" has friendships with 1 and 2. Therefore, the percentage popularity for user 6 would be calculated as (2/9) * 100 = 22.22.
-- - User "7" has friendships with 2. Therefore, the percentage popularity for user 7 would be calculated as (1/9) * 100 = 11.11.
-- - User "8" has friendships with 3. Therefore, the percentage popularity for user 8 would be calculated as (1/9) * 100 = 11.11.
-- - User "9" has friendships with 3. Therefore, the percentage popularity for user 9 would be calculated as (1/9) * 100 = 11.11.
-- user1 is sorted in ascending order.

-- Create table if not exists Friends (user1 int, user2 int)
-- Truncate table Friends
-- insert into Friends (user1, user2) values ('2', '1')
-- insert into Friends (user1, user2) values ('1', '3')
-- insert into Friends (user1, user2) values ('4', '1')
-- insert into Friends (user1, user2) values ('1', '5')
-- insert into Friends (user1, user2) values ('1', '6')
-- insert into Friends (user1, user2) values ('2', '6')
-- insert into Friends (user1, user2) values ('7', '2')
-- insert into Friends (user1, user2) values ('8', '3')
-- insert into Friends (user1, user2) values ('3', '9')

-- # Write your MySQL query statement below
-- WITH r AS ( -- 所有的关系
--     SELECT 
--         user1,
--         user2
--     FROM
--     (
--         SELECT user1,user2 FROM Friends 
--         UNION ALL
--         SELECT user2 AS user1,user1 AS user2 FROM Friends 
--     ) AS t  
--     GROUP BY 
--         user1,user2
-- )

-- SELECT 
--     user1,
--     ROUND(
--         COUNT(distinct user1)  / (SELECT COUNT(*) FROM Friends ) * 100
--         ,
--         2
--     ) AS percentage_popularity 
-- FROM 
--     r
-- GROUP BY
--     user1
-- ORDER BY 
--     user1

-- SELECT
--     user1,
--     ROUND(
--         (SELECT count(1) FROM r WHERE user1 = t.user1) / (select count(distinct user1) from r) * 100,
--         2
--     ) AS percentage_popularity 
-- FROM
--     (
--         SELECT distinct user1 FROM r
--     ) AS t 
-- ORDER BY 
--     user1


WITH r AS (
    SELECT user1 AS u1,user2 AS u2 FROM Friends 
    UNION ALL
    SELECT user2 AS u1,user1 AS u2 FROM Friends 
)

SELECT
    t.u1 AS user1,
    ROUND(
        (SELECT count(1) FROM r where u1 = t.u1) / (SELECT COUNT(distinct u1) FROM r) * 100,
        2
    ) AS  percentage_popularity 
FROM
    r
GROUP BY 
    u1


WITH r AS (
    SELECT user1,user2 FROM Friends 
    UNION ALL
    SELECT user2 AS user1,user1 AS user2 FROM Friends 
)

SELECT 
    user1,
    ROUND(
        COUNT(DISTINCT user2) / (select COUNT(DISTINCT user1) FROM r) * 100.00,
        2
    ) AS percentage_popularity
FROM
    r
GROUP BY 
    1