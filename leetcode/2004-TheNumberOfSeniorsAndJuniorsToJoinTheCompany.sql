-- 2004. The Number of Seniors and Juniors to Join the Company
-- Table: Candidates
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | employee_id | int  |
-- | experience  | enum |
-- | salary      | int  |
-- +-------------+------+
-- employee_id is the column with unique values for this table.
-- experience is an ENUM (category) type of values ('Senior', 'Junior').
-- Each row of this table indicates the id of a candidate, their monthly salary, and their experience.
-- A company wants to hire new employees. The budget of the company for the salaries is $70000. The company's criteria for hiring are:
--     Hiring the largest number of seniors.
--     After hiring the maximum number of seniors, use the remaining budget to hire the largest number of juniors.

-- Write a solution to find the number of seniors and juniors hired under the mentioned criteria.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Candidates table:
-- +-------------+------------+--------+
-- | employee_id | experience | salary |
-- +-------------+------------+--------+
-- | 1           | Junior     | 10000  |
-- | 9           | Junior     | 10000  |
-- | 2           | Senior     | 20000  |
-- | 11          | Senior     | 20000  |
-- | 13          | Senior     | 50000  |
-- | 4           | Junior     | 40000  |
-- +-------------+------------+--------+
-- Output: 
-- +------------+---------------------+
-- | experience | accepted_candidates |
-- +------------+---------------------+
-- | Senior     | 2                   |
-- | Junior     | 2                   |
-- +------------+---------------------+
-- Explanation: 
-- We can hire 2 seniors with IDs (2, 11). Since the budget is $70000 and the sum of their salaries is $40000, we still have $30000 but they are not enough to hire the senior candidate with ID 13.
-- We can hire 2 juniors with IDs (1, 9). Since the remaining budget is $30000 and the sum of their salaries is $20000, we still have $10000 but they are not enough to hire the junior candidate with ID 4.

-- Example 2:
-- Input: 
-- Candidates table:
-- +-------------+------------+--------+
-- | employee_id | experience | salary |
-- +-------------+------------+--------+
-- | 1           | Junior     | 10000  |
-- | 9           | Junior     | 10000  |
-- | 2           | Senior     | 80000  |
-- | 11          | Senior     | 80000  |
-- | 13          | Senior     | 80000  |
-- | 4           | Junior     | 40000  |
-- +-------------+------------+--------+
-- Output: 
-- +------------+---------------------+
-- | experience | accepted_candidates |
-- +------------+---------------------+
-- | Senior     | 0                   |
-- | Junior     | 3                   |
-- +------------+---------------------+
-- Explanation: 
-- We cannot hire any seniors with the current budget as we need at least $80000 to hire one senior.
-- We can hire all three juniors with the remaining budget.

-- Create table If Not Exists Candidates (employee_id int, experience ENUM('Senior', 'Junior'), salary int)
-- Truncate table Candidates
-- insert into Candidates (employee_id, experience, salary) values ('1', 'Junior', '10000')
-- insert into Candidates (employee_id, experience, salary) values ('9', 'Junior', '10000')
-- insert into Candidates (employee_id, experience, salary) values ('2', 'Senior', '20000')
-- insert into Candidates (employee_id, experience, salary) values ('11', 'Senior', '20000')
-- insert into Candidates (employee_id, experience, salary) values ('13', 'Senior', '50000')
-- insert into Candidates (employee_id, experience, salary) values ('4', 'Junior', '40000')

WITH t AS (
    SELECT 
        employee_id,
        experience,
        salary,
        SUM(salary) OVER(PARTITION BY experience ORDER BY salary,employee_id) AS total_salary -- 不同级别累加出薪水
    FROM 
        Candidates
)
-- SELECT * FROM t
-- | employee_id | experience | salary | total_salary |
-- | ----------- | ---------- | ------ | ------------ |
-- | 2           | Senior     | 20000  | 20000        |
-- | 11          | Senior     | 20000  | 40000        |
-- | 13          | Senior     | 50000  | 90000        |
-- | 1           | Junior     | 10000  | 10000        |
-- | 9           | Junior     | 10000  | 20000        |
-- | 4           | Junior     | 40000  | 60000        |

(-- 先取出 70000总额 能找的所有高级员工
    SELECT 
        IFNULL(experience,'Senior') AS experience,
        COUNT(distinct employee_id) AS accepted_candidates 
    FROM 
        t 
    WHERE 
        experience='Senior' AND total_salary <= 70000
)
UNION ALL
(-- 70000 除去招聘高级人员的花费能招到的人
    SELECT 
        IFNULL(experience,'Junior') AS experience,
        COUNT(distinct employee_id) AS accepted_candidates 
    FROM 
        t
    WHERE 
        experience = 'Junior' AND 
        total_salary <= 70000 - ( --  70000 减去 招聘高级人员的花费 本题的解法关键
            SELECT 
                IFNULL(MAX(total_salary), 0) 
            FROM 
                t 
            WHERE  
                experience='Senior' AND total_salary <= 70000
        )
)

