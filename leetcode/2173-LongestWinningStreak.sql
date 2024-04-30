-- 2173. Longest Winning Streak
-- Table: Matches
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | player_id   | int  |
-- | match_day   | date |
-- | result      | enum |
-- +-------------+------+
-- (player_id, match_day) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the ID of a player, the day of the match they played, and the result of that match.
-- The result column is an ENUM (category) type of ('Win', 'Draw', 'Lose').
 
-- The winning streak of a player is the number of consecutive wins uninterrupted by draws or losses.
-- Write a solution to count the longest winning streak for each player.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Matches table:
-- +-----------+------------+--------+
-- | player_id | match_day  | result |
-- +-----------+------------+--------+
-- | 1         | 2022-01-17 | Win    |
-- | 1         | 2022-01-18 | Win    |
-- | 1         | 2022-01-25 | Win    |
-- | 1         | 2022-01-31 | Draw   |
-- | 1         | 2022-02-08 | Win    |
-- | 2         | 2022-02-06 | Lose   |
-- | 2         | 2022-02-08 | Lose   |
-- | 3         | 2022-03-30 | Win    |
-- +-----------+------------+--------+
-- Output: 
-- +-----------+----------------+
-- | player_id | longest_streak |
-- +-----------+----------------+
-- | 1         | 3              |
-- | 2         | 0              |
-- | 3         | 1              |
-- +-----------+----------------+
-- Explanation: 
-- Player 1:
-- From 2022-01-17 to 2022-01-25, player 1 won 3 consecutive matches.
-- On 2022-01-31, player 1 had a draw.
-- On 2022-02-08, player 1 won a match.
-- The longest winning streak was 3 matches.
-- Player 2:
-- From 2022-02-06 to 2022-02-08, player 2 lost 2 consecutive matches.
-- The longest winning streak was 0 matches.
-- Player 3:
-- On 2022-03-30, player 3 won a match.
-- The longest winning streak was 1 match.
 
-- Follow up: If we are interested in calculating the longest streak without losing (i.e., win or draw), how will your solution change?

-- Create table If Not Exists Matches (player_id int, match_day date, result ENUM('Win', 'Draw', 'Lose'))
-- Truncate table Matches
-- insert into Matches (player_id, match_day, result) values ('1', '2022-01-17', 'Win')
-- insert into Matches (player_id, match_day, result) values ('1', '2022-01-18', 'Win')
-- insert into Matches (player_id, match_day, result) values ('1', '2022-01-25', 'Win')
-- insert into Matches (player_id, match_day, result) values ('1', '2022-01-31', 'Draw')
-- insert into Matches (player_id, match_day, result) values ('1', '2022-02-08', 'Win')
-- insert into Matches (player_id, match_day, result) values ('2', '2022-02-06', 'Lose')
-- insert into Matches (player_id, match_day, result) values ('2', '2022-02-08', 'Lose')
-- insert into Matches (player_id, match_day, result) values ('3', '2022-03-30', 'Win')

-- Write your MySQL query statement below

-- use ROW_NUMBER
-- SELECT 
--     player_id, 
--     match_day, 
--     result,
--     (
--         ROW_NUMBER() OVER(PARTITION BY player_id ORDER BY match_day ASC) 
--         - 
--         ROW_NUMBER() OVER(PARTITION BY player_id, result ORDER BY match_day ASC) 
--     ) AS id
-- FROM 
--     Matches

-- | player_id | match_day  | result | id |
-- | --------- | ---------- | ------ | -- |
-- | 1         | 2022-01-17 | Win    | 0  |
-- | 1         | 2022-01-18 | Win    | 0  |
-- | 1         | 2022-01-25 | Win    | 0  |
-- | 1         | 2022-02-08 | Win    | 1  |
-- | 1         | 2022-01-31 | Draw   | 3  |
-- | 2         | 2022-02-06 | Lose   | 0  |
-- | 2         | 2022-02-08 | Lose   | 0  |
-- | 3         | 2022-03-30 | Win    | 0  |

SELECT 
    player_id, 
    MAX(cnt) AS longest_streak
FROM 
(
    SELECT player_id, 
        id, 
        SUM(if(result = 'Win',1, 0)) AS cnt
    FROM 
    (
        SELECT 
            player_id, 
            match_day, 
            result,
            (
                ROW_NUMBER() OVER(PARTITION BY player_id ORDER BY match_day ASC) 
                - 
                ROW_NUMBER() OVER(PARTITION BY player_id, result ORDER BY match_day ASC) 
            ) AS id
        FROM 
            Matches
    ) AS a
    GROUP BY 
        player_id, id
) AS b
GROUP BY 
    player_id


-- use variable
SELECT 
    player_id,
    CONVERT(MAX(temp),UNSIGNED INTEGER) AS longest_streak 
FROM
(
    SELECT 
        *,
        @p:=(
            CASE 
                WHEN @q=player_id AND result = 'Win' THEN @p+1 -- 赢一场 + 1
                WHEN result != 'Win' THEN 0 -- 平或负重新累加
                ELSE 1 -- 首次赢设置为1
            END
        ) AS temp, 
        @q:=player_id 
    FROM
    ( SELECT * FROM Matches ORDER BY player_id,match_day) AS t,
    (SELECT @p:=0, @q:=NULL) AS init
) AS t 
GROUP BY 
    player_id