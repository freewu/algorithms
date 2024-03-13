-- 1212. Team Scores in Football Tournament
-- Table: Teams
-- +---------------+----------+
-- | Column Name   | Type     |
-- +---------------+----------+
-- | team_id       | int      |
-- | team_name     | varchar  |
-- +---------------+----------+
-- team_id is the column with unique values of this table.
-- Each row of this table represents a single football team.

-- Table: Matches
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | match_id      | int     |
-- | host_team     | int     |
-- | guest_team    | int     | 
-- | host_goals    | int     |
-- | guest_goals   | int     |
-- +---------------+---------+
-- match_id is the column of unique values of this table.
-- Each row is a record of a finished match between two different teams. 
-- Teams host_team and guest_team are represented by their IDs in the Teams table (team_id), and they scored host_goals and guest_goals goals, respectively.
 
-- You would like to compute the scores of all teams after all matches. Points are awarded as follows:
--         A team receives three points if they win a match (i.e., Scored more goals than the opponent team).
--         A team receives one point if they draw a match (i.e., Scored the same number of goals as the opponent team).
--         A team receives no points if they lose a match (i.e., Scored fewer goals than the opponent team).

-- Write a solution that selects the team_id, team_name and num_points of each team in the tournament after all described matches.
-- Return the result table ordered by num_points in decreasing order. 
-- In case of a tie, order the records by team_id in increasing order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Teams table:
-- +-----------+--------------+
-- | team_id   | team_name    |
-- +-----------+--------------+
-- | 10        | Leetcode FC  |
-- | 20        | NewYork FC   |
-- | 30        | Atlanta FC   |
-- | 40        | Chicago FC   |
-- | 50        | Toronto FC   |
-- +-----------+--------------+
-- Matches table:
-- +------------+--------------+---------------+-------------+--------------+
-- | match_id   | host_team    | guest_team    | host_goals  | guest_goals  |
-- +------------+--------------+---------------+-------------+--------------+
-- | 1          | 10           | 20            | 3           | 0            |
-- | 2          | 30           | 10            | 2           | 2            |
-- | 3          | 10           | 50            | 5           | 1            |
-- | 4          | 20           | 30            | 1           | 0            |
-- | 5          | 50           | 30            | 1           | 0            |
-- +------------+--------------+---------------+-------------+--------------+
-- Output: 
-- +------------+--------------+---------------+
-- | team_id    | team_name    | num_points    |
-- +------------+--------------+---------------+
-- | 10         | Leetcode FC  | 7             |
-- | 20         | NewYork FC   | 3             |
-- | 50         | Toronto FC   | 3             |
-- | 30         | Atlanta FC   | 1             |
-- | 40         | Chicago FC   | 0             |
-- +------------+--------------+---------------+

-- Create table If Not Exists Teams (team_id int, team_name varchar(30))
-- Create table If Not Exists Matches (match_id int, host_team int, guest_team int, host_goals int, guest_goals int)
-- Truncate table Teams
-- insert into Teams (team_id, team_name) values ('10', 'Leetcode FC')
-- insert into Teams (team_id, team_name) values ('20', 'NewYork FC')
-- insert into Teams (team_id, team_name) values ('30', 'Atlanta FC')
-- insert into Teams (team_id, team_name) values ('40', 'Chicago FC')
-- insert into Teams (team_id, team_name) values ('50', 'Toronto FC')
-- Truncate table Matches
-- insert into Matches (match_id, host_team, guest_team, host_goals, guest_goals) values ('1', '10', '20', '3', '0')
-- insert into Matches (match_id, host_team, guest_team, host_goals, guest_goals) values ('2', '30', '10', '2', '2')
-- insert into Matches (match_id, host_team, guest_team, host_goals, guest_goals) values ('3', '10', '50', '5', '1')
-- insert into Matches (match_id, host_team, guest_team, host_goals, guest_goals) values ('4', '20', '30', '1', '0')
-- insert into Matches (match_id, host_team, guest_team, host_goals, guest_goals) values ('5', '50', '30', '1', '0')

-- Write your MySQL query statement below
SELECT
    t.team_id AS team_id,
    t.team_name AS team_name,
    IFNULL(SUM(s.score),0) AS num_points
FROM
    Teams AS t 
LEFT JOIN
    (
        (
            SELECT 
                host_team AS team_id,
                IF(
                    -- 如果双方打成平手(即，与对方得分相同)，则得 1 分
                    host_goals = guest_goals,1,
                    -- 如果球队赢了比赛(即比对手进更多的球)，就得 3 分。
                    -- 如果球队输掉了比赛(例如，比对手少进球)，就 不得分
                    IF(host_goals > guest_goals, 3, 0)
                ) AS score
            FROM
                Matches 
        )
        UNION ALL
        (
            SELECT 
                guest_team AS team_id,
                IF(
                    -- 如果双方打成平手(即，与对方得分相同)，则得 1 分
                    host_goals = guest_goals,1,
                    -- 如果球队赢了比赛(即比对手进更多的球)，就得 3 分。
                    -- 如果球队输掉了比赛(例如，比对手少进球)，就 不得分
                    IF(host_goals < guest_goals, 3, 0)
                ) AS score
            FROM
                Matches 
        )
    ) AS s
ON t.team_id = s.team_id
GROUP BY
    t.team_id
ORDER BY
    num_points DESC, -- 根据 num_points 降序排序
    team_id  ASC -- 如果有两队积分相同，那么这两队按 team_id  升序排序