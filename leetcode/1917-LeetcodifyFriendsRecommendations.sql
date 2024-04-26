-- 1917. Leetcodify Friends Recommendations
-- Table: Listens
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | user_id     | int     |
-- | song_id     | int     |
-- | day         | date    |
-- +-------------+---------+
-- This table may contain duplicates (In other words, there is no primary key for this table in SQL).
-- Each row of this table indicates that the user user_id listened to the song song_id on the day day.
 
-- Table: Friendship
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | user1_id      | int     |
-- | user2_id      | int     |
-- +---------------+---------+
-- In SQL,(user1_id, user2_id) is the primary key for this table.
-- Each row of this table indicates that the users user1_id and user2_id are friends.
-- Note that user1_id < user2_id.
 
-- Recommend friends to Leetcodify users. We recommend user x to user y if:
--     Users x and y are not friends, and
--     Users x and y listened to the same three or more different songs on the same day.

-- Note that friend recommendations are unidirectional, meaning if user x and user y should be recommended to each other, the result table should have both user x recommended to user y and user y recommended to user x. Also, note that the result table should not contain duplicates (i.e., user y should not be recommended to user x multiple times.).
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Listens table:
-- +---------+---------+------------+
-- | user_id | song_id | day        |
-- +---------+---------+------------+
-- | 1       | 10      | 2021-03-15 |
-- | 1       | 11      | 2021-03-15 |
-- | 1       | 12      | 2021-03-15 |
-- | 2       | 10      | 2021-03-15 |
-- | 2       | 11      | 2021-03-15 |
-- | 2       | 12      | 2021-03-15 |
-- | 3       | 10      | 2021-03-15 |
-- | 3       | 11      | 2021-03-15 |
-- | 3       | 12      | 2021-03-15 |
-- | 4       | 10      | 2021-03-15 |
-- | 4       | 11      | 2021-03-15 |
-- | 4       | 13      | 2021-03-15 |
-- | 5       | 10      | 2021-03-16 |
-- | 5       | 11      | 2021-03-16 |
-- | 5       | 12      | 2021-03-16 |
-- +---------+---------+------------+
-- Friendship table:
-- +----------+----------+
-- | user1_id | user2_id |
-- +----------+----------+
-- | 1        | 2        |
-- +----------+----------+
-- Output: 
-- +---------+----------------+
-- | user_id | recommended_id |
-- +---------+----------------+
-- | 1       | 3              |
-- | 2       | 3              |
-- | 3       | 1              |
-- | 3       | 2              |
-- +---------+----------------+
-- Explanation: 
-- Users 1 and 2 listened to songs 10, 11, and 12 on the same day, but they are already friends.
-- Users 1 and 3 listened to songs 10, 11, and 12 on the same day. Since they are not friends, we recommend them to each other.
-- Users 1 and 4 did not listen to the same three songs.
-- Users 1 and 5 listened to songs 10, 11, and 12, but on different days.
-- Similarly, we can see that users 2 and 3 listened to songs 10, 11, and 12 on the same day and are not friends, so we recommend them to each other.

-- Create table If Not Exists Listens (user_id int, song_id int, day date)
-- Create table If Not Exists Friendship (user1_id int, user2_id int)
-- Truncate table Listens
-- insert into Listens (user_id, song_id, day) values ('1', '10', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('1', '11', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('1', '12', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('2', '10', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('2', '11', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('2', '12', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('3', '10', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('3', '11', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('3', '12', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('4', '10', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('4', '11', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('4', '13', '2021-03-15')
-- insert into Listens (user_id, song_id, day) values ('5', '10', '2021-03-16')
-- insert into Listens (user_id, song_id, day) values ('5', '11', '2021-03-16')
-- insert into Listens (user_id, song_id, day) values ('5', '12', '2021-03-16')
-- Truncate table Friendship
-- insert into Friendship (user1_id, user2_id) values ('1', '2')


-- # Write your MySQL query statement below
-- -- SELECT 
-- --     a.*,
-- --     b.*
-- -- FROM
-- --     Listens AS a,
-- --     Listens AS b
-- -- WHERE
-- --     a.user_id != b.user_id AND -- 先排除自己
-- --     a.song_id = b.song_id AND a.day = b.day -- 当天听了相同歌

-- WITH f AS (
--     SELECT user1_id,user2_id FROM Friendship
--     UNION
--     SELECT user2_id AS user1_id, user1_id AS user2_id FROM Friendship
-- ),
-- r AS (
--     SELECT 
--         a.user_id,
--         b.user_id AS recommended_id 
--     FROM
--         Listens AS a,
--         Listens AS b
--     WHERE
--         a.user_id != b.user_id AND -- 先排除自己
--         (a.song_id, a.day) = (b.song_id, b.day) AND -- 当天听了相同歌
--         (a.user_id, b.user_id) NOT IN ( SELECT user1_id,user2_id FROM f ) -- 排除已是好友的
--     GROUP BY
--         a.user_id,b.user_id
--     HAVING 
--         COUNT(DISTINCT a.song_id) >= 3 -- 用户 x 和 y 在同一天收听了相同的三首或更多不同歌曲
-- )

-- SELECT user_id,recommended_id FROM r 
-- UNION
-- SELECT recommended_id AS user_id, user_id AS recommended_id FROM r 

# Write your MySQL query statement below

-- 所有好友关系
with f(user1_id, user2_id) as (
    select user1_id, user2_id from Friendship
    union All 
    select user2_id, user1_id from Friendship
)

select 
    distinct l1.user_id,
    l2.user_id as recommended_id 
from 
    Listens l1
inner join 
    Listens l2
on 
    l1.song_id = l2.song_id and l1.day = l2.day and -- 当天听了相同歌
    l1.user_id <> l2.user_id and -- 先排除自己
    not exists (select 1 from f where user1_id = l1.user_id and l2.user_id = user2_id) -- 排除已是好友的
group by 
    l1.user_id, l2.user_id, l1.day
having 
    count(distinct l2.song_id) >= 3 -- 用户 x 和 y 在同一天收听了相同的三首或更多不同歌曲

-- best solution
with t1 as  (
    select 
        distinct user_id,
        song_id,
        day
    from 
        Listens 
),
t2 as (
    select 
        a.user_id as user1_id,
        b.user_id as user2_id
    from 
        t1 AS a 
    join 
        t1 AS b 
    on 
        a.user_id < b.user_id and
        a.song_id = b.song_id and a.day = b.day
    where 
        not exists (select user1_id,user2_id from Friendship where user1_id=a.user_id and user2_id=b.user_id) 
    group by 
        1,2,b.day
    having 
        count(distinct b.song_id)>=3 
)
select user1_id as user_id,user2_id as recommended_id from t2
union 
select user2_id as user_id,user1_id as recommended_id from t2