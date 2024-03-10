-- 2339. All the Matches of the League
-- Table: Teams
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | team_name   | varchar |
-- +-------------+---------+
-- team_name is the column with unique values of this table.
-- Each row of this table shows the name of a team.
-- Write a solution to report all the possible matches of the league. Note that every two teams play two matches with each other, with one team being the home_team once and the other time being the away_team.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Teams table:
-- +-------------+
-- | team_name   |
-- +-------------+
-- | Leetcode FC |
-- | Ahly SC     |
-- | Real Madrid |
-- +-------------+
-- Output: 
-- +-------------+-------------+
-- | home_team   | away_team   |
-- +-------------+-------------+
-- | Real Madrid | Leetcode FC |
-- | Real Madrid | Ahly SC     |
-- | Leetcode FC | Real Madrid |
-- | Leetcode FC | Ahly SC     |
-- | Ahly SC     | Real Madrid |
-- | Ahly SC     | Leetcode FC |
-- +-------------+-------------+
-- Explanation: All the matches of the league are shown in the table.

-- Create table If Not Exists Teams (team_name varchar(50))
-- Truncate table Teams
-- insert into Teams (team_name) values ('Leetcode FC')
-- insert into Teams (team_name) values ('Ahly SC')
-- insert into Teams (team_name) values ('Real Madrid')

-- Write your MySQL query statement below
SELECT 
    a.team_name AS home_team,
    b.team_name AS away_team
FROM
    Teams AS a,
    Teams AS b 
WHERE
    a.team_name != b.team_name