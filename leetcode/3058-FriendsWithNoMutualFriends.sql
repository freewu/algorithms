-- 3058. Friends With No Mutual Friends
-- Table: Friends
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | user_id1    | int  |
-- | user_id2    | int  |
-- +-------------+------+
-- (user_id1, user_id2) is the primary key (combination of columns with unique values) for this table.
-- Each row contains user id1, user id2, both of whom are friends with each other.
-- Write a solution to find all pairs of users who are friends with each other and have no mutual friends.
-- Return the result table ordered by user_id1, user_id2 in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Friends table:
-- +----------+----------+
-- | user_id1 | user_id2 | 
-- +----------+----------+
-- | 1        | 2        | 
-- | 2        | 3        | 
-- | 2        | 4        | 
-- | 1        | 5        | 
-- | 6        | 7        | 
-- | 3        | 4        | 
-- | 2        | 5        | 
-- | 8        | 9        | 
-- +----------+----------+
-- Output: 
-- +----------+----------+
-- | user_id1 | user_id2 | 
-- +----------+----------+
-- | 6        | 7        | 
-- | 8        | 9        | 
-- +----------+----------+
-- Explanation: 
-- - Users 1 and 2 are friends with each other, but they share a mutual friend with user ID 5, so this pair is not included.
-- - Users 2 and 3 are friends, they both share a mutual friend with user ID 4, resulting in exclusion, similarly for users 2 and 4 who share a mutual friend with user ID 3, hence not included.
-- - Users 1 and 5 are friends with each other, but they share a mutual friend with user ID 2, so this pair is not included.
-- - Users 6 and 7, as well as users 8 and 9, are friends with each other, and they don't have any mutual friends, hence included.
-- - Users 3 and 4 are friends with each other, but their mutual connection with user ID 2 means they are not included, similarly for users 2 and 5 are friends but are excluded due to their mutual connection with user ID 1.
-- Output table is ordered by user_id1 in ascending order.

-- Create Table if Not Exists Friends( user_id1 int, user_id2 int)
-- Truncate table Friends
-- insert into Friends (user_id1, user_id2) values ('1', '2')
-- insert into Friends (user_id1, user_id2) values ('2', '3')
-- insert into Friends (user_id1, user_id2) values ('2', '4')
-- insert into Friends (user_id1, user_id2) values ('1', '5')
-- insert into Friends (user_id1, user_id2) values ('6', '7')
-- insert into Friends (user_id1, user_id2) values ('3', '4')
-- insert into Friends (user_id1, user_id2) values ('2', '5')
-- insert into Friends (user_id1, user_id2) values ('8', '9')

# Write your MySQL query statement below
-- SELECT 
--     f.*
-- FROM
--     Friends AS f, 
--     (
--         (
--             SELECT 
--                 user_id1,
--                 user_id2
--             FROM
--                 Friends 
--         ) 
--         UNION ALL 
--         (
--             SELECT 
--                 user_id2 AS user_id1,
--                 user_id1 AS user_id2
--             FROM
--                 Friends 
--         )
--     ) AS t 
-- WHERE
--     f.user_id1 = t.user_id1 AND
--     f.user_id2 = t.user_id2

WITH t1 AS 
( -- 所有关系
    (SELECT user_id1, user_id2 FROM Friends)
    UNION ALL
    (SELECT user_id2, user_id1 FROM Friends)  
),
t2 AS 
(
    SELECT 
        a.user_id2 AS u1,
        b.user_id2 AS u2,
        a.*
    FROM 
        t1 a 
    LEFT JOIN 
        t1 b 
    ON 
        a.user_id1 = b.user_id1 and a.user_id2 != b.user_id2
    WHERE 
        b.user_id2 is not null
)

-- select * from t2

SELECT 
    * 
FROM 
    Friends 
WHERE 
    (user_id1, user_id2) NOT IN (SELECT u1,u2 FROM t2) AND
    (user_id1, user_id2) NOT IN (SELECT u2,u1 FROM t2)
ORDER BY 
    user_id1 -- 输出表以 user_id1 升序排列