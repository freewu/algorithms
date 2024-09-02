-- 3278. Find Candidates for Data Scientist Position II
-- Table: Candidates
-- +--------------+---------+ 
-- | Column Name  | Type    | 
-- +--------------+---------+ 
-- | candidate_id | int     | 
-- | skill        | varchar |
-- | proficiency  | int     |
-- +--------------+---------+
-- (candidate_id, skill) is the unique key for this table.
-- Each row includes candidate_id, skill, and proficiency level (1-5).
-- Table: Projects
-- +--------------+---------+ 
-- | Column Name  | Type    | 
-- +--------------+---------+ 
-- | project_id   | int     | 
-- | skill        | varchar |
-- | importance   | int     |
-- +--------------+---------+
-- (project_id, skill) is the primary key for this table.
-- Each row includes project_id, required skill, and its importance (1-5) for the project.
-- Leetcode is staffing for multiple data science projects. 
-- Write a solution to find the best candidate for each project based on the following criteria:
--     1. Candidates must have all the skills required for a project.
--     2. Calculate a score for each candidate-project pair as follows:
--         Start with 100 points
--         Add 10 points for each skill where proficiency > importance
--         Subtract 5 points for each skill where proficiency < importance

-- Include only the top candidate (highest score) for each project. 
-- If there’s a tie, choose the candidate with the lower candidate_id. 
-- If there is no suitable candidate for a project, do not return that project.

-- Return a result table ordered by project_id in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- Candidates table:
-- +--------------+-----------+-------------+
-- | candidate_id | skill     | proficiency |
-- +--------------+-----------+-------------+
-- | 101          | Python    | 5           |
-- | 101          | Tableau   | 3           |
-- | 101          | PostgreSQL| 4           |
-- | 101          | TensorFlow| 2           |
-- | 102          | Python    | 4           |
-- | 102          | Tableau   | 5           |
-- | 102          | PostgreSQL| 4           |
-- | 102          | R         | 4           |
-- | 103          | Python    | 3           |
-- | 103          | Tableau   | 5           |
-- | 103          | PostgreSQL| 5           |
-- | 103          | Spark     | 4           |
-- +--------------+-----------+-------------+
-- Projects table:
-- +-------------+-----------+------------+
-- | project_id  | skill     | importance |
-- +-------------+-----------+------------+
-- | 501         | Python    | 4          |
-- | 501         | Tableau   | 3          |
-- | 501         | PostgreSQL| 5          |
-- | 502         | Python    | 3          |
-- | 502         | Tableau   | 4          |
-- | 502         | R         | 2          |
-- +-------------+-----------+------------+
-- Output:
-- +-------------+--------------+-------+
-- | project_id  | candidate_id | score |
-- +-------------+--------------+-------+
-- | 501         | 101          | 105   |
-- | 502         | 102          | 130   |
-- +-------------+--------------+-------+
-- Explanation:
-- For Project 501, Candidate 101 has the highest score of 105. All other candidates have the same score but Candidate 101 has the lowest candidate_id among them.
-- For Project 502, Candidate 102 has the highest score of 130.
-- The output table is ordered by project_id in ascending order.

-- Create table if not exists Candidates(candidate_id int, skill varchar(50), proficiency int)
-- Create table if not exists Projects( project_id int, skill varchar(50), importance int)
-- Truncate table Candidates
-- insert into Candidates (candidate_id, skill, proficiency) values ('101', 'Python', '5')
-- insert into Candidates (candidate_id, skill, proficiency) values ('101', 'Tableau', '3')
-- insert into Candidates (candidate_id, skill, proficiency) values ('101', 'PostgreSQL', '4')
-- insert into Candidates (candidate_id, skill, proficiency) values ('101', 'TensorFlow', '2')
-- insert into Candidates (candidate_id, skill, proficiency) values ('102', 'Python', '4')
-- insert into Candidates (candidate_id, skill, proficiency) values ('102', 'Tableau', '5')
-- insert into Candidates (candidate_id, skill, proficiency) values ('102', 'PostgreSQL', '4')
-- insert into Candidates (candidate_id, skill, proficiency) values ('102', 'R', '4')
-- insert into Candidates (candidate_id, skill, proficiency) values ('103', 'Python', '3')
-- insert into Candidates (candidate_id, skill, proficiency) values ('103', 'Tableau', '5')
-- insert into Candidates (candidate_id, skill, proficiency) values ('103', 'PostgreSQL', '5')
-- insert into Candidates (candidate_id, skill, proficiency) values ('103', 'Spark', '4')
-- Truncate table Projects
-- insert into Projects (project_id, skill, importance) values ('501', 'Python', '4')
-- insert into Projects (project_id, skill, importance) values ('501', 'Tableau', '3')
-- insert into Projects (project_id, skill, importance) values ('501', 'PostgreSQL', '5')
-- insert into Projects (project_id, skill, importance) values ('502', 'Python', '3')
-- insert into Projects (project_id, skill, importance) values ('502', 'Tableau', '4')
-- insert into Projects (project_id, skill, importance) values ('502', 'R', '2')

--  Write your MySQL query statement below
WITH t AS ( -- 得到每个用户每个技能在在每个项目中的评分
    SELECT
        c.candidate_id, 
        c.skill,
        p.project_id,
        IF(c.proficiency > p.importance, 10, IF(c.proficiency = p.importance, 0, -5)) AS score
    FROM
        Candidates AS c,
        Projects AS p
    WHERE
        c.skill = p.skill
),
t1 AS ( -- 筛选出所有符合条件的候选人
    SELECT 
        a.*
    FROM 
        (
            SELECT
                candidate_id,
                project_id,
                count(*) AS count,
                SUM(score) AS score
            FROM 
                t 
            GROUP BY 
                candidate_id, project_id
        ) AS a,
        (
            SELECT 
                project_id, 
                count(*) AS count
            FROM
                Projects 
            GROUP BY 
                project_id
        ) AS b
    WHERE
        a.project_id = b.project_id AND a.count = b.count -- 候选人必须拥有项目所需的 所有 技能
),
t2 AS ( -- 给候选人排名
        SELECT
        t1.*,
        RANK() OVER ( PARTITION BY project_id ORDER BY score desc,candidate_id ) AS rk -- 每个项目的最佳候选人（最高分）。如果 相同，选择有 更小 candidate_id 的候选人
    FROM
        t1
)

-- select * from t
-- | candidate_id | skill      | project_id | score |
-- | ------------ | ---------- | ---------- | ----- |
-- | 101          | Python     | 502        | 10    |
-- | 101          | Python     | 501        | 10    |
-- | 101          | Tableau    | 502        | -5    |
-- | 101          | Tableau    | 501        | 10    |
-- | 101          | PostgreSQL | 501        | -5    |
-- | 102          | Python     | 502        | 10    |
-- | 102          | Python     | 501        | 10    |
-- | 102          | Tableau    | 502        | 10    |
-- | 102          | Tableau    | 501        | 10    |
-- | 102          | PostgreSQL | 501        | -5    |
-- | 102          | R          | 502        | 10    |
-- | 103          | Python     | 502        | 10    |
-- | 103          | Python     | 501        | -5    |
-- | 103          | Tableau    | 502        | 10    |
-- | 103          | Tableau    | 501        | 10    |
-- | 103          | PostgreSQL | 501        | 10    |

-- select * from t1
-- | candidate_id | project_id | count | score |
-- | ------------ | ---------- | ----- | ----- |
-- | 101          | 501        | 3     | 5     |
-- | 102          | 501        | 3     | 5     |
-- | 103          | 501        | 3     | 5     |
-- | 102          | 502        | 3     | 30    |

-- select * from t2
-- | candidate_id | project_id | count | score | rk |
-- | ------------ | ---------- | ----- | ----- | -- |
-- | 101          | 501        | 3     | 5     | 1  |
-- | 102          | 501        | 3     | 5     | 2  |
-- | 103          | 501        | 3     | 5     | 3  |
-- | 102          | 502        | 3     | 30    | 1  |

SELECT
    project_id,
    candidate_id,
    (score + 100) AS score -- 从 100 分 开始
FROM
    t2 
WHERE
    t2.rk = 1
ORDER BY 
    project_id -- 结果表以 project_id 升序排序