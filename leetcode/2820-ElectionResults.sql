-- 2820. Election Results
-- Table: Votes
-- +-------------+---------+ 
-- | Column Name | Type    | 
-- +-------------+---------+ 
-- | voter       | varchar | 
-- | candidate   | varchar |
-- +-------------+---------+
-- (voter, candidate) is the primary key (combination of unique values) for this table.
-- Each row of this table contains name of the voter and their candidate. 
-- The election is conducted in a city where everyone can vote for one or more candidates or choose not to vote. Each person has 1 vote so if they vote for multiple candidates, their vote gets equally split across them. For example, if a person votes for 2 candidates, these candidates receive an equivalent of 0.5 votes each.

-- Write a solution to find candidate who got the most votes and won the election. Output the name of the candidate or If multiple candidates have an equal number of votes, display the names of all of them.
-- Return the result table ordered by candidate in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Votes table:
-- +----------+-----------+
-- | voter    | candidate |
-- +----------+-----------+
-- | Kathy    | null      |
-- | Charles  | Ryan      |
-- | Charles  | Christine |
-- | Charles  | Kathy     |
-- | Benjamin | Christine |
-- | Anthony  | Ryan      |
-- | Edward   | Ryan      |
-- | Terry    | null      |
-- | Evelyn   | Kathy     |
-- | Arthur   | Christine |
-- +----------+-----------+
-- Output: 
-- +-----------+
-- | candidate | 
-- +-----------+
-- | Christine |  
-- | Ryan      |  
-- +-----------+
-- Explanation: 
-- - Kathy and Terry opted not to participate in voting, resulting in their votes being recorded as 0. Charles distributed his vote among three candidates, equating to 0.33 for each candidate. On the other hand, Benjamin, Arthur, Anthony, Edward, and Evelyn each cast their votes for a single candidate.
-- - Collectively, Candidate Ryan and Christine amassed a total of 2.33 votes, while Kathy received a combined total of 1.33 votes.
-- Since Ryan and Christine received an equal number of votes, we will display their names in ascending order.

-- Create table if not exists Votes(voter varchar(30), candidate varchar(30))
-- Truncate table Votes
-- insert into Votes (voter, candidate) values ('Kathy', 'None')
-- insert into Votes (voter, candidate) values ('Charles', 'Ryan')
-- insert into Votes (voter, candidate) values ('Charles', 'Christine')
-- insert into Votes (voter, candidate) values ('Charles', 'Kathy')
-- insert into Votes (voter, candidate) values ('Benjamin', 'Christine')
-- insert into Votes (voter, candidate) values ('Anthony', 'Ryan')
-- insert into Votes (voter, candidate) values ('Edward', 'Ryan')
-- insert into Votes (voter, candidate) values ('Terry', 'None')
-- insert into Votes (voter, candidate) values ('Evelyn', 'Kathy')
-- insert into Votes (voter, candidate) values ('Arthur', 'Christine')

-- Write your MySQL query statement below
-- -- 计算每个投票者的投票值 投的人越多越小
-- SELECT 
--     voter,
--     COUNT(*) AS cnt,
--     1 / COUNT(*) AS  vote_val
-- FROM
--     Votes
-- GROUP BY 
--     voter

-- -- 统计评分并排名
-- SELECT 
--     v.candidate,
--     SUM(t.vote_val) AS val,
--     RANK() OVER(ORDER BY SUM(t.vote_val) DESC) AS rk
-- FROM 
--     Votes AS v 
-- LEFT JOIN
--     (
--         -- 计算每个投票者的投票值 投的人越多越小
--         SELECT 
--             voter,
--             COUNT(*) AS cnt,
--             1 / COUNT(*) AS  vote_val
--         FROM
--             Votes
--         GROUP BY 
--             voter
--     ) AS t 
-- ON 
--     t.voter = v.voter
-- WHERE
--     v.candidate IS NOT NULL
-- GROUP BY
--     v.candidate 

SELECT 
    candidate
FROM 
    (
        -- 统计评分并排名
        SELECT 
            v.candidate,
            SUM(t.vote_val) AS val,
            RANK() OVER(ORDER BY SUM(t.vote_val) DESC) AS rk
        FROM 
            Votes AS v 
        LEFT JOIN
            (
                -- 计算每个投票者的投票值 投的人越多越小
                SELECT 
                    voter,
                    COUNT(*) AS cnt,
                    1 / COUNT(*) AS  vote_val
                FROM
                    Votes
                GROUP BY 
                    voter
            ) AS t 
        ON 
            t.voter = v.voter
        WHERE
            v.candidate IS NOT NULL
        GROUP BY
            v.candidate 
    ) AS tt
WHERE
    rk = 1
ORDER BY 
    candidate -- 返回按 candidate 升序排序 的结果表