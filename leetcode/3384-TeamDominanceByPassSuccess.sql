-- 3384. Team Dominance by Pass Success
-- Table: Teams
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | player_id   | int     |
-- | team_name   | varchar | 
-- +-------------+---------+
-- player_id is the unique key for this table.
-- Each row contains the unique identifier for player and the name of one of the teams participating in that match.

-- Table: Passes
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | pass_from   | int     |
-- | time_stamp  | varchar |
-- | pass_to     | int     |
-- +-------------+---------+
-- (pass_from, time_stamp) is the primary key for this table.
-- pass_from is a foreign key to player_id from Teams table.
-- Each row represents a pass made during a match, time_stamp represents the time in minutes (00:00-90:00) when the pass was made,
-- pass_to is the player_id of the player receiving the pass.

-- Write a solution to calculate the dominance score for each team in both halves of the match. The rules are as follows:

-- A match is divided into two halves: first half (00:00-45:00 minutes) and second half (45:01-90:00 minutes)
-- The dominance score is calculated based on successful and intercepted passes:
-- When pass_to is a player from the same team: +1 point
-- When pass_to is a player from the opposing team (interception): -1 point
-- A higher dominance score indicates better passing performance
-- Return the result table ordered by team_name and half_number in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- Teams table:
-- +------------+-----------+
-- | player_id  | team_name |
-- +------------+-----------+
-- | 1          | Arsenal   |
-- | 2          | Arsenal   |
-- | 3          | Arsenal   |
-- | 4          | Chelsea   |
-- | 5          | Chelsea   |
-- | 6          | Chelsea   |
-- +------------+-----------+
-- Passes table:
-- +-----------+------------+---------+
-- | pass_from | time_stamp | pass_to |
-- +-----------+------------+---------+
-- | 1         | 00:15      | 2       |
-- | 2         | 00:45      | 3       |
-- | 3         | 01:15      | 1       |
-- | 4         | 00:30      | 1       |
-- | 2         | 46:00      | 3       |
-- | 3         | 46:15      | 4       |
-- | 1         | 46:45      | 2       |
-- | 5         | 46:30      | 6       |
-- +-----------+------------+---------+
-- Output:
-- +-----------+-------------+-----------+
-- | team_name | half_number | dominance |
-- +-----------+-------------+-----------+
-- | Arsenal   | 1           | 3         |
-- | Arsenal   | 2           | 1         |
-- | Chelsea   | 1           | -1        |
-- | Chelsea   | 2           | 1         |
-- +-----------+-------------+-----------+
-- Explanation:
-- First Half (00:00-45:00):
-- Arsenal's passes:
-- 1 → 2 (00:15): Successful pass (+1)
-- 2 → 3 (00:45): Successful pass (+1)
-- 3 → 1 (01:15): Successful pass (+1)
-- Chelsea's passes:
-- 4 → 1 (00:30): Intercepted by Arsenal (-1)
-- Second Half (45:01-90:00):
-- Arsenal's passes:
-- 2 → 3 (46:00): Successful pass (+1)
-- 3 → 4 (46:15): Intercepted by Chelsea (-1)
-- 1 → 2 (46:45): Successful pass (+1)
-- Chelsea's passes:
-- 5 → 6 (46:30): Successful pass (+1)
-- The results are ordered by team_name and then half_number

-- CREATE TABLE If not exists Teams (
--     player_id INT,
--     team_name VARCHAR(100)
-- )
-- CREATE TABLE if not exists Passes (
--     pass_from INT,
--     time_stamp VARCHAR(5),
--     pass_to INT
-- )
-- Truncate table Teams
-- insert into Teams (player_id, team_name) values ('1', 'Arsenal')
-- insert into Teams (player_id, team_name) values ('2', 'Arsenal')
-- insert into Teams (player_id, team_name) values ('3', 'Arsenal')
-- insert into Teams (player_id, team_name) values ('4', 'Chelsea')
-- insert into Teams (player_id, team_name) values ('5', 'Chelsea')
-- insert into Teams (player_id, team_name) values ('6', 'Chelsea')
-- Truncate table Passes
-- insert into Passes (pass_from, time_stamp, pass_to) values ('1', '00:15', '2')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('2', '00:45', '3')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('3', '01:15', '1')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('4', '00:30', '1')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('2', '46:00', '3')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('3', '46:15', '4')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('1', '46:45', '2')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('5', '46:30', '6')

-- Write your MySQL query statement below
SELECT 
    t1.team_name,
    IF(time_stamp <= '45:00', 1, 2) AS half_number, -- 判断时间是上半场还是下半场
    SUM(IF(t1.team_name = t2. team_name, 1, -1)) AS dominance -- 优势得分
FROM 
    Passes AS p
JOIN 
    Teams AS t1 ON pass_from = t1.player_id -- Passes表与Teams表连接 确定pass_from的队伍
JOIN 
    Teams AS t2 ON pass_to = t2.player_id -- Passes表与Teams表连接 确定pass_to的队伍
GROUP BY 
    t1.team_name, IF(time_stamp <= '45:00', 1, 2) -- 分组 pass_from 的队伍，和场次
ORDER BY 
    t1.team_name, half_number
