-- 602. Friend Requests II: Who Has the Most Friends
-- Table: RequestAccepted
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | requester_id   | int     |
-- | accepter_id    | int     |
-- | accept_date    | date    |
-- +----------------+---------+
-- (requester_id, accepter_id) is the primary key for this table.
-- This table contains the ID of the user who sent the request, the ID of the user who received the request, and the date when the request was accepted.
--  
-- Write an SQL query to find the people who have the most friends and the most friends number.
-- The test cases are generated so that only one person has the most friends.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- RequestAccepted table:
-- +--------------+-------------+-------------+
-- | requester_id | accepter_id | accept_date |
-- +--------------+-------------+-------------+
-- | 1            | 2           | 2016/06/03  |
-- | 1            | 3           | 2016/06/08  |
-- | 2            | 3           | 2016/06/08  |
-- | 3            | 4           | 2016/06/09  |
-- +--------------+-------------+-------------+
-- Output:
-- +----+-----+
-- | id | num |
-- +----+-----+
-- | 3  | 3   |
-- +----+-----+
-- Explanation:
-- The person with id 3 is a friend of people 1, 2, and 4, so he has three friends in total, which is the most number than any others.
--  
-- Follow up: In the real world, multiple people could have the same most number of friends. Could you find all these people in this case?

-- Create table If Not Exists RequestAccepted (requester_id int not null, accepter_id int null, accept_date date null)
-- Truncate table RequestAccepted
-- insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('1', '2', '2016/06/03')
-- insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('1', '3', '2016/06/08')
-- insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('2', '3', '2016/06/08')
-- insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('3', '4', '2016/06/09')

-- Write your MySQL query statement below
SELECT
    a.id,
    (a.num + IFNULL(b.num,0)) AS num
FROM
    (
        SELECT
            requester_id AS id,
            COUNT(*) AS num
        FROM
            RequestAccepted
        GROUP BY
            requester_id
    ) AS a
LEFT JOIN
    (
        SELECT
            accepter_id AS id,
            COUNT(*) AS num
        FROM
            RequestAccepted
        GROUP BY
            accepter_id
    ) AS b
ON
    a.id = b.id
ORDER BY
    num DESC
LIMIT 1

-- UNION best solution
SELECT
    id,
    COUNT(*) as num
FROM
    (
        SELECT requester_id AS id FROM RequestAccepted
        UNION ALL
        SELECT  accepter_id AS id FROM RequestAccepted
    ) AS t
GROUP BY
    id
ORDER BY
    num DESC
LIMIT 1