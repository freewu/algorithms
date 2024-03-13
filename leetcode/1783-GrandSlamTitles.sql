-- 1783. Grand Slam Titles
-- Table: Players
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | player_id      | int     |
-- | player_name    | varchar |
-- +----------------+---------+
-- player_id is the primary key (column with unique values) for this table.
-- Each row in this table contains the name and the ID of a tennis player.

-- Table: Championships
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | year          | int     |
-- | Wimbledon     | int     |
-- | Fr_open       | int     |
-- | US_open       | int     |
-- | Au_open       | int     |
-- +---------------+---------+
-- year is the primary key (column with unique values) for this table.
-- Each row of this table contains the IDs of the players who won one each tennis tournament of the grand slam.
 
-- Write a solution to report the number of grand slam tournaments won by each player. 
-- Do not include the players who did not win any tournament.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Players table:
-- +-----------+-------------+
-- | player_id | player_name |
-- +-----------+-------------+
-- | 1         | Nadal       |
-- | 2         | Federer     |
-- | 3         | Novak       |
-- +-----------+-------------+
-- Championships table:
-- +------+-----------+---------+---------+---------+
-- | year | Wimbledon | Fr_open | US_open | Au_open |
-- +------+-----------+---------+---------+---------+
-- | 2018 | 1         | 1       | 1       | 1       |
-- | 2019 | 1         | 1       | 2       | 2       |
-- | 2020 | 2         | 1       | 2       | 2       |
-- +------+-----------+---------+---------+---------+
-- Output: 
-- +-----------+-------------+-------------------+
-- | player_id | player_name | grand_slams_count |
-- +-----------+-------------+-------------------+
-- | 2         | Federer     | 5                 |
-- | 1         | Nadal       | 7                 |
-- +-----------+-------------+-------------------+
-- Explanation: 
-- Player 1 (Nadal) won 7 titles: Wimbledon (2018, 2019), Fr_open (2018, 2019, 2020), US_open (2018), and Au_open (2018).
-- Player 2 (Federer) won 5 titles: Wimbledon (2020), US_open (2019, 2020), and Au_open (2019, 2020).
-- Player 3 (Novak) did not win anything, we did not include them in the result table.

-- Create table If Not Exists Players (player_id int, player_name varchar(20))
-- Create table If Not Exists Championships (year int, Wimbledon int, Fr_open int, US_open int, Au_open int)
-- Truncate table Players
-- insert into Players (player_id, player_name) values ('1', 'Nadal')
-- insert into Players (player_id, player_name) values ('2', 'Federer')
-- insert into Players (player_id, player_name) values ('3', 'Novak')
-- Truncate table Championships
-- insert into Championships (year, Wimbledon, Fr_open, US_open, Au_open) values ('2018', '1', '1', '1', '1')
-- insert into Championships (year, Wimbledon, Fr_open, US_open, Au_open) values ('2019', '1', '1', '2', '2')
-- insert into Championships (year, Wimbledon, Fr_open, US_open, Au_open) values ('2020', '2', '1', '2', '2')

SELECT
    p.player_id AS player_id,
    p.player_name AS player_name,
    COUNT(*) AS grand_slams_count 
FROM
    Players AS p,
    (
        ( -- 温网
            SELECT
                Wimbledon AS player_id
            FROM
                Championships 
        ) 
        UNION ALL 
        ( -- 法国公开赛
            SELECT
                Fr_open AS player_id
            FROM
                Championships 
        )
        UNION ALL 
        ( -- 美国公开赛
            SELECT
                US_open AS player_id
            FROM
                Championships 
        ) 
        UNION ALL 
        ( -- 澳网公开赛
            SELECT
                Au_open AS player_id
            FROM
                Championships 
        ) 
    ) AS r
WHERE
    p.player_id  = r.player_id 
GROUP BY
    p.player_id
