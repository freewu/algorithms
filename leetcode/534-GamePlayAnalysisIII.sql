-- 534. Game Play Analysis III
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
-- Write an SQL query to report for each player and date, how many games played so far by the player.
-- That is, the total number of games played by the player until that date. Check the example for clarity.
-- Return the result table in any order.
-- The query result format is in the following example.

-- Example 1:
-- Input:
-- Activity table:
-- +-----------+-----------+------------+--------------+
-- | player_id | device_id | event_date | games_played |
-- +-----------+-----------+------------+--------------+
-- | 1         | 2         | 2016-03-01 | 5            |
-- | 1         | 2         | 2016-05-02 | 6            |
-- | 1         | 3         | 2017-06-25 | 1            |
-- | 3         | 1         | 2016-03-02 | 0            |
-- | 3         | 4         | 2018-07-03 | 5            |
-- +-----------+-----------+------------+--------------+
-- Output:
-- +-----------+------------+---------------------+
-- | player_id | event_date | games_played_so_far |
-- +-----------+------------+---------------------+
-- | 1         | 2016-03-01 | 5                   |
-- | 1         | 2016-05-02 | 11                  |
-- | 1         | 2017-06-25 | 12                  |
-- | 3         | 2016-03-02 | 0                   |
-- | 3         | 2018-07-03 | 5                   |
-- +-----------+------------+---------------------+
-- Explanation:
-- For the player with id 1, 5 + 6 = 11 games played by 2016-05-02, and 5 + 6 + 1 = 12 games played by 2017-06-25.
-- For the player with id 3, 0 + 5 = 5 games played by 2018-07-03.

-- Create table If Not Exists Activity (player_id int, device_id int, event_date date, games_played int)
-- Truncate table Activity
-- insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-03-01', '5')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-05-02', '6')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('1', '3', '2017-06-25', '1')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('3', '1', '2016-03-02', '0')
-- insert into Activity (player_id, device_id, event_date, games_played) values ('3', '4', '2018-07-03', '5')

-- Write your MySQL query statement below
SELECT
    a.player_id,
    a.event_date,
    SUM(b.games_played) AS games_played_so_far  -- 在此日期之前玩家所玩的游戏总数
FROM
    Activity AS a,
    Activity AS b
WHERE
    a.player_id = b.player_id AND
    a.event_date >= b.event_date -- >= 很重要，要统计当天的
GROUP BY
    a.player_id,
    a.event_date

-- best solution
SELECT 
    player_id,
    event_date,
    SUM(games_played) OVER(PARTITION BY player_id ORDER BY player_id, event_date ASC) AS games_played_so_far
FROM 
    Activity
ORDER BY
    player_id, event_date DESC