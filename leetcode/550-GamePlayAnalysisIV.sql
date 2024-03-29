-- 550. Game Play Analysis IV
-- Table: Activity
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | player_id    | int     |
-- | device_id    | int     |
-- | event_date   | date    |
-- | games_played | int     |
-- +--------------+---------+
-- (player_id, event_date) is the primary key of this table.
-- This table shows the activity of players of some games.
-- Each row is a record of a player who logged in and played a number of games (possibly 0) before logging out on someday using some device.
-- Write an SQL query to report the fraction of players that logged in again on the day after the day they first logged in,
-- rounded to 2 decimal places. In other words, you need to count the number of players that logged in for at least two consecutive days starting from their first login date,
-- then divide that number by the total number of players.
-- The query result format is in the following example.

-- Example 1:
-- Input:
-- Activity table:
-- +-----------+-----------+------------+--------------+
-- | player_id | device_id | event_date | games_played |
-- +-----------+-----------+------------+--------------+
-- | 1         | 2         | 2016-03-01 | 5            |
-- | 1         | 2         | 2016-03-02 | 6            |
-- | 2         | 3         | 2017-06-25 | 1            |
-- | 3         | 1         | 2016-03-02 | 0            |
-- | 3         | 4         | 2018-07-03 | 5            |
-- +-----------+-----------+------------+--------------+
-- Output:
-- +-----------+
-- | fraction  |
-- +-----------+
-- | 0.33      |
-- +-----------+
-- Explanation:
-- Only the player with id 1 logged back in after the first day he had logged in so the answer is 1/3 = 0.33

-- Create table If Not Exists Activity (player_id int, device_id int, event_date date, games_played int)
-- Truncate table Activity
-- insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-03-01', '5')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-03-02', '6')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('2', '3', '2017-06-25', '1')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('3', '1', '2016-03-02', '0')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('3', '4', '2018-07-03', '5')


-- Write your MySQL query statement below
-- 首日注册后第二天连续登录的.不是任意两天连续登录
# SELECT
#     ROUND((
#         SELECT
#             COUNT(DISTINCT a.player_id)
#         FROM
#             `Activity` AS a,
#             `Activity` AS b
#         WHERE
#             a.player_id = b.player_id AND
#             a.event_date = DATE_ADD(b.event_date,INTERVAL -1 DAY)
#     )
#     /
#     (
#         SELECT COUNT(DISTINCT player_id) FROM `Activity`
#     ),2) AS fraction

SELECT
	ROUND(
        COUNT(DISTINCT player_id) /
        (SELECT COUNT(distinct player_id) FROM Activity), -- 总人数
	2) AS fraction
FROM
    Activity
WHERE
	(
        player_id,
        event_date
    ) IN (
        SELECT
            player_id,
            Date(min(event_date)+1) -- 第二日期  min(event_date) 注册日
	    FROM
            Activity
	    GROUP BY
            player_id
    );

-- best solution
SELECT
    ROUND( AVG(a.event_date is not null), 2) AS fraction -- AVG(a.event_date is not null) 这个太赞了
FROM
    ( -- 取用户首次注册日期&用记编号
        SELECT
            player_id,
            MIN(event_date) AS login
        FROM activity
        GROUP BY player_id
    ) AS p
LEFT JOIN
    activity AS a
ON
    p.player_id = a.player_id AND
    DATEDIFF(a.event_date, p.login) = 1