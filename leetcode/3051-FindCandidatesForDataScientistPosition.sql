-- 3051. Find Candidates for Data Scientist Position
-- Table: Candidates
-- +--------------+---------+ 
-- | Column Name  | Type    | 
-- +--------------+---------+ 
-- | candidate_id | int     | 
-- | skill        | varchar |
-- +--------------+---------+
-- (candidate_id, skill) is the primary key (columns with unique values) for this table.
-- Each row includes candidate_id and skill.
-- Write a query to find the candidates best suited for a Data Scientist position. The candidate must be proficient in Python, Tableau, and PostgreSQL.
-- Return the result table ordered by candidate_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Candidates table:
-- +---------------+--------------+
-- | candidate_id  | skill        | 
-- +---------------+--------------+
-- | 123           | Python       |
-- | 234           | R            | 
-- | 123           | Tableau      | 
-- | 123           | PostgreSQL   | 
-- | 234           | PowerBI      | 
-- | 234           | SQL Server   | 
-- | 147           | Python       | 
-- | 147           | Tableau      | 
-- | 147           | Java         |
-- | 147           | PostgreSQL   |
-- | 256           | Tableau      |
-- | 102           | DataAnalysis |
-- +---------------+--------------+
-- Output: 
-- +--------------+
-- | candidate_id |  
-- +--------------+
-- | 123          |  
-- | 147          | 
-- +--------------+
-- Explanation: 
-- - Candidates 123 and 147 possess the necessary skills in Python, Tableau, and PostgreSQL for the data scientist position.
-- - Candidates 234 and 102 do not possess any of the required skills for this position.
-- - Candidate 256 has proficiency in Tableau but is missing skills in Python and PostgreSQL.
-- The output table is sorted by candidate_id in ascending order.

-- Create table If Not Exists Candidates (candidate_id int, skill varchar(30))
-- Truncate table Candidates
-- insert into Candidates (candidate_id, skill) values ('123', 'Python')
-- insert into Candidates (candidate_id, skill) values ('234', 'R')
-- insert into Candidates (candidate_id, skill) values ('123', 'Tableau')
-- insert into Candidates (candidate_id, skill) values ('123', 'PostgreSQL')
-- insert into Candidates (candidate_id, skill) values ('234', 'PowerBI')
-- insert into Candidates (candidate_id, skill) values ('234', 'SQL Server')
-- insert into Candidates (candidate_id, skill) values ('147', 'Python')
-- insert into Candidates (candidate_id, skill) values ('147', 'Tableau')
-- insert into Candidates (candidate_id, skill) values ('147', 'Java')
-- insert into Candidates (candidate_id, skill) values ('147', 'PostgreSQL')
-- insert into Candidates (candidate_id, skill) values ('256', 'Tableau')
-- insert into Candidates (candidate_id, skill) values ('102', 'DataAnalysis')

-- Write your MySQL query statement below
SELECT
    candidate_id
FROM
    Candidates 
WHERE
    skill IN ('Python','Tableau','PostgreSQL')
GROUP BY
    candidate_id
HAVING
    COUNT(*) = 3 --  The candidate must be proficient in Python, Tableau, and PostgreSQL.
ORDER BY 
    candidate_id -- the result table ordered by candidate_id in ascending order.
