-- 3390. Longest Team Pass Streak
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
-- (pass_from, time_stamp) is the unique key for this table.
-- pass_from is a foreign key to player_id from Teams table.
-- Each row represents a pass made during a match, time_stamp represents the time in minutes (00:00-90:00) when the pass was made,
-- pass_to is the player_id of the player receiving the pass.
-- Write a solution to find the longest successful pass streak for each team during the match. 
-- The rules are as follows:
--     A successful pass streak is defined as consecutive passes where:
--         Both the pass_from and pass_to players belong to the same team
--     A streak breaks when either:
--         The pass is intercepted (received by a player from the opposing team)

-- Return the result table ordered by team_name in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- Teams table:
-- +-----------+-----------+
-- | player_id | team_name |
-- +-----------+-----------+
-- | 1         | Arsenal   |
-- | 2         | Arsenal   |
-- | 3         | Arsenal   |
-- | 4         | Arsenal   |
-- | 5         | Chelsea   |
-- | 6         | Chelsea   |
-- | 7         | Chelsea   |
-- | 8         | Chelsea   |
-- +-----------+-----------+
-- Passes table:
-- +-----------+------------+---------+
-- | pass_from | time_stamp | pass_to |
-- +-----------+------------+---------+
-- | 1         | 00:05      | 2       |
-- | 2         | 00:07      | 3       |
-- | 3         | 00:08      | 4       |
-- | 4         | 00:10      | 5       |
-- | 6         | 00:15      | 7       |
-- | 7         | 00:17      | 8       |
-- | 8         | 00:20      | 6       |
-- | 6         | 00:22      | 5       |
-- | 1         | 00:25      | 2       |
-- | 2         | 00:27      | 3       |
-- +-----------+------------+---------+
-- Output:
-- +-----------+----------------+
-- | team_name | longest_streak |
-- +-----------+----------------+
-- | Arsenal   | 3              |
-- | Chelsea   | 4              |
-- +-----------+----------------+
-- Explanation:
--     Arsenal's streaks:
--         First streak: 3 passes (1→2→3→4) ended when player 4 passed to Chelsea's player 5
--         Second streak: 2 passes (1→2→3)
--         Longest streak = 3
--     Chelsea's streaks:
--         First streak: 3 passes (6→7→8→6→5)
--         Longest streak = 4

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
-- insert into Teams (player_id, team_name) values ('4', 'Arsenal')
-- insert into Teams (player_id, team_name) values ('5', 'Chelsea')
-- insert into Teams (player_id, team_name) values ('6', 'Chelsea')
-- insert into Teams (player_id, team_name) values ('7', 'Chelsea')
-- insert into Teams (player_id, team_name) values ('8', 'Chelsea')
-- Truncate table Passes
-- insert into Passes (pass_from, time_stamp, pass_to) values ('1', '00:05', '2')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('2', '00:07', '3')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('3', '00:08', '4')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('4', '00:10', '5')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('6', '00:15', '7')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('7', '00:17', '8')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('8', '00:20', '6')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('6', '00:22', '5')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('1', '00:25', '2')
-- insert into Passes (pass_from, time_stamp, pass_to) values ('2', '00:27', '3')

-- Write your MySQL query statement below
WITH d1 AS 
(
    SELECT 
        SUM(t1.team_name != t2.team_name) OVER (PARTITION BY t1.team_name ORDER BY p.time_stamp) AS total,
        t1.team_name team_name
    FROM 
        Passes AS p
    JOIN Teams AS t1 ON p.pass_from = t1.player_id
    JOIN Teams AS t2 ON p.pass_to = t2.player_id
)
, d2 AS
(
    SELECT 
        team_name,
        IF(total = 0, COUNT(*), COUNT(*) - 1) AS count
        --CASE WHEN total = 0 THEN count(*) ELSE count(*) - 1 END AS count
    FROM 
        d1
    GROUP BY 
        team_name, total
    HAVING 
        count != 0
)

SELECT 
    team_name,
    MAX(count) AS longest_streak
FROM 
    d2
GROUP BY
    team_name
ORDER BY  
    team_name -- Return the result table ordered by team_name in ascending order.
