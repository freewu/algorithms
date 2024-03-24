-- 1355. Activity Participants
-- Table: Friends
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | name          | varchar |
-- | activity      | varchar |
-- +---------------+---------+
-- id is the id of the friend and the primary key for this table in SQL.
-- name is the name of the friend.
-- activity is the name of the activity which the friend takes part in.

-- Table: Activities
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | name          | varchar |
-- +---------------+---------+
-- In SQL, id is the primary key for this table.
-- name is the name of the activity.
 
-- Find the names of all the activities with neither the maximum nor the minimum number of participants.
-- Each activity in the Activities table is performed by any person in the table Friends.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:

-- Input: 
-- Friends table:
-- +------+--------------+---------------+
-- | id   | name         | activity      |
-- +------+--------------+---------------+
-- | 1    | Jonathan D.  | Eating        |
-- | 2    | Jade W.      | Singing       |
-- | 3    | Victor J.    | Singing       |
-- | 4    | Elvis Q.     | Eating        |
-- | 5    | Daniel A.    | Eating        |
-- | 6    | Bob B.       | Horse Riding  |
-- +------+--------------+---------------+
-- Activities table:
-- +------+--------------+
-- | id   | name         |
-- +------+--------------+
-- | 1    | Eating       |
-- | 2    | Singing      |
-- | 3    | Horse Riding |
-- +------+--------------+
-- Output: 
-- +--------------+
-- | activity     |
-- +--------------+
-- | Singing      |
-- +--------------+
-- Explanation: 
-- Eating activity is performed by 3 friends, maximum number of participants, (Jonathan D. , Elvis Q. and Daniel A.)
-- Horse Riding activity is performed by 1 friend, minimum number of participants, (Bob B.)
-- Singing is performed by 2 friends (Victor J. and Jade W.)

-- Create table If Not Exists Friends (id int, name varchar(30), activity varchar(30))
-- Create table If Not Exists Activities (id int, name varchar(30))
-- Truncate table Friends
-- insert into Friends (id, name, activity) values ('1', 'Jonathan D.', 'Eating')
-- insert into Friends (id, name, activity) values ('2', 'Jade W.', 'Singing')
-- insert into Friends (id, name, activity) values ('3', 'Victor J.', 'Singing')
-- insert into Friends (id, name, activity) values ('4', 'Elvis Q.', 'Eating')
-- insert into Friends (id, name, activity) values ('5', 'Daniel A.', 'Eating')
-- insert into Friends (id, name, activity) values ('6', 'Bob B.', 'Horse Riding')
-- Truncate table Activities
-- insert into Activities (id, name) values ('1', 'Eating')
-- insert into Activities (id, name) values ('2', 'Singing')
-- insert into Activities (id, name) values ('3', 'Horse Riding')

-- SELECT
--     MAX(cnt) AS max_cnt,
--     MIN(cnt) AS min_cnt
-- FROM
-- (
--     SELECT
--         COUNT(*) AS cnt,
--         activity      
--     FROM
--         Friends 
--     GROUP BY
--         activity
-- ) AS a

-- SELECT 
--     a.activity AS activity
-- FROM 
-- (-- 统计活动& 参加人数
--     SELECT
--         COUNT(*) AS cnt,
--         activity      
--     FROM
--         Friends 
--     GROUP BY
--         activity
-- ) AS a,
-- ( -- 最大值 & 最小值 
--     SELECT
--         MAX(cnt) AS max_cnt,
--         MIN(cnt) AS min_cnt
--     FROM
--     (
--         SELECT
--             COUNT(*) AS cnt,
--             activity      
--         FROM
--             Friends 
--         GROUP BY
--             activity
--     ) AS t
-- ) AS b
-- WHERE
--     a.cnt > b.min_cnt AND
--     a.cnt < b.max_cnt

WITH t AS (
    -- 统计活动& 参加人数
    SELECT
        COUNT(*) AS cnt,
        activity      
    FROM
        Friends 
    GROUP BY
        activity
)

SELECT 
    a.activity AS activity
FROM 
    t AS a,
    ( -- 最大值 & 最小值 
        SELECT
            MAX(cnt) AS max_cnt,
            MIN(cnt) AS min_cnt
        FROM
            t
    ) AS b
WHERE
    a.cnt > b.min_cnt AND
    a.cnt < b.max_cnt


-- rank
select 
    name activity
from (
    select 
        a.name, 
        ifnull(cnt,0) as cnt, 
        rank() over(order by ifnull(cnt,0)) r1, 
        rank() over(order by ifnull(cnt,0) desc) r2
    from 
        activities a 
    left join (
        select 
            activity, 
            count(distinct id) as cnt 
        from 
            friends 
        group by activity
    ) f 
    on a.name=f.activity
) t 
where r1!=1 and r2!=1;