-- 1892. Page Recommendations II
-- Table: Friendship
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | user1_id      | int     |
-- | user2_id      | int     |
-- +---------------+---------+
-- (user1_id, user2_id) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table indicates that the users user1_id and user2_id are friends.
 
-- Table: Likes
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | user_id     | int     |
-- | page_id     | int     |
-- +-------------+---------+
-- (user_id, page_id) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table indicates that user_id likes page_id.
 
-- You are implementing a page recommendation system for a social media website. Your system will recommend a page to user_id if the page is liked by at least one friend of user_id and is not liked by user_id.
-- Write a solution to find all the possible page recommendations for every user. Each recommendation should appear as a row in the result table with these columns:

--     user_id: The ID of the user that your system is making the recommendation to.
--     page_id: The ID of the page that will be recommended to user_id.
--     friends_likes: The number of the friends of user_id that like page_id.

-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:

-- Input: 
-- Friendship table:
-- +----------+----------+
-- | user1_id | user2_id |
-- +----------+----------+
-- | 1        | 2        |
-- | 1        | 3        |
-- | 1        | 4        |
-- | 2        | 3        |
-- | 2        | 4        |
-- | 2        | 5        |
-- | 6        | 1        |
-- +----------+----------+
-- Likes table:
-- +---------+---------+
-- | user_id | page_id |
-- +---------+---------+
-- | 1       | 88      |
-- | 2       | 23      |
-- | 3       | 24      |
-- | 4       | 56      |
-- | 5       | 11      |
-- | 6       | 33      |
-- | 2       | 77      |
-- | 3       | 77      |
-- | 6       | 88      |
-- +---------+---------+
-- Output: 
-- +---------+---------+---------------+
-- | user_id | page_id | friends_likes |
-- +---------+---------+---------------+
-- | 1       | 77      | 2             |
-- | 1       | 23      | 1             |
-- | 1       | 24      | 1             |
-- | 1       | 56      | 1             |
-- | 1       | 33      | 1             |
-- | 2       | 24      | 1             |
-- | 2       | 56      | 1             |
-- | 2       | 11      | 1             |
-- | 2       | 88      | 1             |
-- | 3       | 88      | 1             |
-- | 3       | 23      | 1             |
-- | 4       | 88      | 1             |
-- | 4       | 77      | 1             |
-- | 4       | 23      | 1             |
-- | 5       | 77      | 1             |
-- | 5       | 23      | 1             |
-- +---------+---------+---------------+
-- Explanation: 
-- Take user 1 as an example:
--   - User 1 is friends with users 2, 3, 4, and 6.
--   - Recommended pages are 23 (user 2 liked it), 24 (user 3 liked it), 56 (user 3 liked it), 33 (user 6 liked it), and 77 (user 2 and user 3 liked it).
--   - Note that page 88 is not recommended because user 1 already liked it.
-- Another example is user 6:
--   - User 6 is friends with user 1.
--   - User 1 only liked page 88, but user 6 already liked it. Hence, user 6 has no recommendations.
-- You can recommend pages for users 2, 3, 4, and 5 using a similar process.

-- Create table If Not Exists Friendship (user1_id int, user2_id int)
-- Create table If Not Exists Likes (user_id int, page_id int)
-- Truncate table Friendship
-- insert into Friendship (user1_id, user2_id) values ('1', '2')
-- insert into Friendship (user1_id, user2_id) values ('1', '3')
-- insert into Friendship (user1_id, user2_id) values ('1', '4')
-- insert into Friendship (user1_id, user2_id) values ('2', '3')
-- insert into Friendship (user1_id, user2_id) values ('2', '4')
-- insert into Friendship (user1_id, user2_id) values ('2', '5')
-- insert into Friendship (user1_id, user2_id) values ('6', '1')
-- Truncate table Likes
-- insert into Likes (user_id, page_id) values ('1', '88')
-- insert into Likes (user_id, page_id) values ('2', '23')
-- insert into Likes (user_id, page_id) values ('3', '24')
-- insert into Likes (user_id, page_id) values ('4', '56')
-- insert into Likes (user_id, page_id) values ('5', '11')
-- insert into Likes (user_id, page_id) values ('6', '33')
-- insert into Likes (user_id, page_id) values ('2', '77')
-- insert into Likes (user_id, page_id) values ('3', '77')
-- insert into Likes (user_id, page_id) values ('6', '88')

-- 超出时间限制 8 / 13
WITH f AS (
    SELECT  user1_id, user2_id FROM Friendship
    UNION ALL
    SELECT  user2_id AS user1_id, user1_id AS user2_id FROM Friendship
)

-- SELECT 
--     f.*,
--     l.*
-- FROM
--     f,
--     Likes AS l 
-- WHERE
--     f.user1_id != l.user_id AND f.user2_id = l.user_id -- 至少一个朋友喜欢 ，而 不被user_id喜欢

SELECT 
    user1_id AS user_id,
    a.page_id AS page_id,
    COUNT(DISTINCT f.user2_id) AS friends_likes 
FROM
    f,
    Likes AS a
WHERE
    f.user1_id != a.user_id AND f.user2_id = a.user_id AND -- 至少一个朋友喜欢
    (f.user1_id, a.page_id) NOT IN ( SELECT user_id, page_id FROM Likes ) -- 而不被user_id喜欢
GROUP BY
    f.user1_id, a.page_id


-- not exist
SELECT 
    f.user1_id AS user_id,
    a.page_id AS page_id,
    COUNT(DISTINCT f.user2_id) AS friends_likes 
FROM
    Likes AS a
INNER JOIN 
    f
ON 
    f.user2_id = a.user_id
WHERE 
    NOT EXISTS (
        SELECT 
            b.user_id 
        FROM 
            Likes b 
        WHERE 
            b.user_id = f.user1_id and b.page_id = a.page_id
    )
GROUP BY 
    user1_id, page_id


-- join a b c
SELECT
    a.user2_id AS user_id,
    b.page_id AS page_id,
    count(1) AS friends_likes
FROM 
    f AS a 
LEFT JOIN 
    Likes b 
ON 
    a.user1_id = b.user_id
LEFT JOIN 
    Likes c 
ON 
    a.user2_id = c.user_id AND b.page_id = c.page_id
WHERE 
    c.page_id IS NULL
GROUP BY
    a.user2_id,b.page_id
