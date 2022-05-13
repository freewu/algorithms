-- 2500 · Highest scoring player in the team
-- Description
-- The players table contains information about all players and the teams table contains information about all teams.
-- Write an SQL statement to find the highest scoring player for each team.

-- Table Definition 1: players (players table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	player name
-- scores	int	player scores
-- team_id	int	player's team

-- Table Definition 2: teams (Teams table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	team name

-- Example
-- Sample 1:

-- Table Contents 1: players

-- id	name	scores	team_id
-- 1	zhangsan	15	2
-- 2	lisi	10	1
-- 3	wanger	23	4
-- 4	zhaoliu	5	1
-- 5	niuniu	17	2
-- 6	zhangfei	32	3
-- 7	guanyu	39	4
-- 8	liubei	16	3
-- 9	linqi	3	2
-- Table Contents 2: teams

-- id	name
-- 1	team 1
-- 2	team 2
-- 3	team 3
-- 4	team 4
-- After running your SQL statement, the table should return.

-- team	player	score
-- team 1	lisi	10
-- team 2	niuniu	17
-- team 3	zhangfei	32
-- team 4	guanyu	39
-- Sample 2:

-- Table Contents 1: players

-- id	name	scores	team_id
-- 1	zhangsan	15	2
-- 2	lisi	10	3
-- 3	wanger	23	4
-- 4	zhaoliu	5	5
-- Table Contents 2: teams

-- id	name
-- 1	team 1
-- After running your SQL statement, it should return.

-- team	player	score
-- There is no matching data in sample 2, so the output contains only table headers and no data.

-- Write your SQL Query here --
SELECT
    t.name AS team,
    p.name AS player,
    p.scores AS score
FROM
    teams AS t,
    players AS p,
    (-- 取每队最多得分球员
        SELECT
            MAX(scores) AS scores,
            team_id
        FROM
            players
        GROUP BY
            team_id
    ) AS b 
WHERE
    t.id = b.team_id AND
    p.team_id = b.team_id AND
    p.scores = b.scores