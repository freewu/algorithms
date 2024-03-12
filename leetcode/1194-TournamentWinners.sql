-- 1194. Tournament Winners
-- Table: Players
-- +-------------+-------+
-- | Column Name | Type  |
-- +-------------+-------+
-- | player_id   | int   |
-- | group_id    | int   |
-- +-------------+-------+
-- player_id is the primary key (column with unique values) of this table.
-- Each row of this table indicates the group of each player.
-- Table: Matches
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | match_id      | int     |
-- | first_player  | int     |
-- | second_player | int     | 
-- | first_score   | int     |
-- | second_score  | int     |
-- +---------------+---------+
-- match_id is the primary key (column with unique values) of this table.
-- Each row is a record of a match, first_player and second_player contain the player_id of each match.
-- first_score and second_score contain the number of points of the first_player and second_player respectively.
-- You may assume that, in each match, players belong to the same group.

-- The winner in each group is the player who scored the maximum total points within the group. 
-- In the case of a tie, the lowest player_id wins.
-- Write a solution to find the winner in each group.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Players table:
-- +-----------+------------+
-- | player_id | group_id   |
-- +-----------+------------+
-- | 15        | 1          |
-- | 25        | 1          |
-- | 30        | 1          |
-- | 45        | 1          |
-- | 10        | 2          |
-- | 35        | 2          |
-- | 50        | 2          |
-- | 20        | 3          |
-- | 40        | 3          |
-- +-----------+------------+
-- Matches table:
-- +------------+--------------+---------------+-------------+--------------+
-- | match_id   | first_player | second_player | first_score | second_score |
-- +------------+--------------+---------------+-------------+--------------+
-- | 1          | 15           | 45            | 3           | 0            |
-- | 2          | 30           | 25            | 1           | 2            |
-- | 3          | 30           | 15            | 2           | 0            |
-- | 4          | 40           | 20            | 5           | 2            |
-- | 5          | 35           | 50            | 1           | 1            |
-- +------------+--------------+---------------+-------------+--------------+
-- Output: 
-- +-----------+------------+
-- | group_id  | player_id  |
-- +-----------+------------+ 
-- | 1         | 15         |
-- | 2         | 35         |
-- | 3         | 40         |
-- +-----------+------------+

-- Create table If Not Exists Players (player_id int, group_id int)
-- Create table If Not Exists Matches (match_id int, first_player int, second_player int, first_score int, second_score int)
-- Truncate table Players
-- insert into Players (player_id, group_id) values ('10', '2')
-- insert into Players (player_id, group_id) values ('15', '1')
-- insert into Players (player_id, group_id) values ('20', '3')
-- insert into Players (player_id, group_id) values ('25', '1')
-- insert into Players (player_id, group_id) values ('30', '1')
-- insert into Players (player_id, group_id) values ('35', '2')
-- insert into Players (player_id, group_id) values ('40', '3')
-- insert into Players (player_id, group_id) values ('45', '1')
-- insert into Players (player_id, group_id) values ('50', '2')
-- Truncate table Matches
-- insert into Matches (match_id, first_player, second_player, first_score, second_score) values ('1', '15', '45', '3', '0')
-- insert into Matches (match_id, first_player, second_player, first_score, second_score) values ('2', '30', '25', '1', '2')
-- insert into Matches (match_id, first_player, second_player, first_score, second_score) values ('3', '30', '15', '2', '0')
-- insert into Matches (match_id, first_player, second_player, first_score, second_score) values ('4', '40', '20', '5', '2')
-- insert into Matches (match_id, first_player, second_player, first_score, second_score) values ('5', '35', '50', '1', '1')

-- UNION ALL
SELECT 
    group_id, 
    player_id
FROM 
(
    SELECT 
        group_id, 
        player_id, 
        SUM(score) AS score
    FROM 
    (
        ( -- 每个用户总的 first_score
            SELECT 
                p.group_id, 
                p.player_id, 
                SUM(m.first_score) AS score
            FROM 
                Players AS p 
            JOIN 
                Matches AS m
            ON p.player_id = m.first_player
            GROUP BY 
                p.player_id
        ) 
        UNION ALL
        ( -- 每个用户总的 second_score
            SELECT 
                p.group_id, 
                p.player_id, 
                SUM(m.second_score) AS score
            FROM 
                Players AS p 
            JOIN 
                Matches AS m
            ON p.player_id = m.second_player
            GROUP BY 
                p.player_id
        ) 
    ) AS s
    GROUP BY 
        player_id
    ORDER BY 
        score DESC, player_id
) AS result
GROUP BY 
    group_id


-- rank
SELECT 
    group_id, 
    player_id 
FROM 
(
	SELECT
		group_id,
		t2.player_id,
		RANK() OVER(PARTITION BY group_id ORDER BY score DESC, t2.player_id) rk -- 如果平局，player_id 最小 的选手获胜
	FROM 
    (
		SELECT 
            player_id, 
            sum(score) AS score 
        FROM 
        (
			(
                SELECT first_player AS player_id, first_score AS score FROM matches
            )
			UNION ALL
            (
                SELECT second_player AS player_id, second_score AS score FROM matches
            )
		) t1 group by player_id
	) AS t2 
    LEFT JOIN players AS p
    ON t2.player_id = p.player_id
) AS t3 
WHERE rk = 1;
