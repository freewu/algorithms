-- 1097. Game Play Analysis V
-- Table: Activity
--
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
--  
-- The install date of a player is the first login day of that player.
-- We define day one retention of some date x to be the number of players whose install date is x and they logged back in on the day right after x,
-- divided by the number of players whose install date is x, rounded to 2 decimal places.
-- Write an SQL query to report for each install date, the number of players that installed the game on that day, and the day one retention.
--
-- Return the result table in any order.
-- The query result format is in the following example.
--
-- Example 1:
--
-- Input:
-- Activity table:
-- +-----------+-----------+------------+--------------+
-- | player_id | device_id | event_date | games_played |
-- +-----------+-----------+------------+--------------+
-- | 1         | 2         | 2016-03-01 | 5            |
-- | 1         | 2         | 2016-03-02 | 6            |
-- | 2         | 3         | 2017-06-25 | 1            |
-- | 3         | 1         | 2016-03-01 | 0            |
-- | 3         | 4         | 2016-07-03 | 5            |
-- +-----------+-----------+------------+--------------+
-- Output:
-- +------------+----------+----------------+
-- | install_dt | installs | Day1_retention |
-- +------------+----------+----------------+
-- | 2016-03-01 | 2        | 0.50           |
-- | 2017-06-25 | 1        | 0.00           |
-- +------------+----------+----------------+
-- Explanation:
-- Player 1 and 3 installed the game on 2016-03-01 but only player 1 logged back in on 2016-03-02 so the day 1 retention of 2016-03-01 is 1 / 2 = 0.50
-- Player 2 installed the game on 2017-06-25 but didn't log back in on 2017-06-26 so the day 1 retention of 2017-06-25 is 0 / 1 = 0.00
--
# Write your MySQL query statement below
# SELECT
#     a.event_date AS install_dt,
#     a.num AS installs,
#     ROUND(IFNULL(b.num / a.num,0),2) AS Day1_retention
# FROM
#     (
#         -- 每天新用户数
#         SELECT
#             event_date,
#             COUNT(*) AS num
#         FROM
#             (
#                 SELECT
#                     player_id,
#                     MIN(event_date) AS event_date
#                 FROM
#                     Activity
#                 GROUP BY
#                     player_id
#             ) AS a1
#         GROUP BY
#             event_date
#     ) AS a
#  LEFT JOIN
#     (
#         -- 留存数量
#         SELECT
#             b1.event_date,
#             COUNT(*) AS num
#         FROM
#             ( -- 用户 & 第一次时间
#                 SELECT
#                     player_id,
#                     MIN(event_date) AS event_date
#                 FROM
#                     Activity
#                 GROUP BY
#                     player_id
#             ) AS b1,
#             Activity AS b2
#         WHERE
#            b1.event_date = DATE_ADD(b2.event_date,interval -1 day) AND
#            b1.player_id = b2.player_id
#     ) AS b
# ON
#     a.event_date = b.event_date
# ORDER BY
#     a.event_date

select a1.install_dt,
       count(*) installs,
       round(count(a2.event_date)/count(*),2) Day1_retention
from(
    select player_id,min(event_date) install_dt
    from Activity
    group by player_id
) a1
left join Activity a2
    on a1.player_id = a2.player_id and datediff(a2.event_date,a1.install_dt)=1
group by a1.install_dt