-- 2853. Highest Salaries Difference
-- Table: Salaries
-- +-------------+---------+ 
-- | Column Name | Type    | 
-- +-------------+---------+ 
-- | emp_name    | varchar | 
-- | department  | varchar | 
-- | salary      | int     |
-- +-------------+---------+
-- (emp_name, department) is the primary key (combination of unique values) for this table.
-- Each row of this table contains emp_name, department and salary. 
-- There will be at least one entry for the engineering and marketing departments.
-- Write a solution to calculate the difference between the highest salaries in the marketing and engineering department. 
-- Output the absolute difference in salaries.
-- Return the result table.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Salaries table:
-- +----------+-------------+--------+
-- | emp_name | department  | salary |
-- +----------+-------------+--------+
-- | Kathy    | Engineering | 50000  |
-- | Roy      | Marketing   | 30000  |
-- | Charles  | Engineering | 45000  |
-- | Jack     | Engineering | 85000  |
-- | Benjamin | Marketing   | 34000  |
-- | Anthony  | Marketing   | 42000  |
-- | Edward   | Engineering | 102000 |
-- | Terry    | Engineering | 44000  |
-- | Evelyn   | Marketing   | 53000  |
-- | Arthur   | Engineering | 32000  |
-- +----------+-------------+--------+
-- Output: 
-- +-------------------+
-- | salary_difference | 
-- +-------------------+
-- | 49000             | 
-- +-------------------+
-- Explanation: 
-- - The Engineering and Marketing departments have the highest salaries of 102,000 and 53,000, respectively. 
-- Resulting in an absolute difference of 49,000.

-- Create table if not exists Salaries(emp_name varchar(30), department varchar(30),salary int)
-- Truncate table Salaries
-- insert into Salaries (emp_name, department, salary) values ('Kathy', 'Engineering', '50000')
-- insert into Salaries (emp_name, department, salary) values ('Roy', 'Marketing', '30000')
-- insert into Salaries (emp_name, department, salary) values ('Charles', 'Engineering', '45000')
-- insert into Salaries (emp_name, department, salary) values ('Jack', 'Engineering', '85000')
-- insert into Salaries (emp_name, department, salary) values ('Benjamin', 'Marketing', '34000')
-- insert into Salaries (emp_name, department, salary) values ('Anthony', 'Marketing', '42000')
-- insert into Salaries (emp_name, department, salary) values ('Edward', 'Engineering', '102000')
-- insert into Salaries (emp_name, department, salary) values ('Terry', 'Engineering', '44000')
-- insert into Salaries (emp_name, department, salary) values ('Evelyn', 'Marketing', '53000')
-- insert into Salaries (emp_name, department, salary) values ('Arthur', 'Engineering', '32000')

SELECT
    ABS(m.salary - e.salary) AS salary_difference
FROM
    (
        SELECT
            MAX(salary) AS salary
        FROM    
            Salaries 
        WHERE     
            department = "Marketing"
    ) AS m,
    (
        SELECT
            MAX(salary) AS salary
        FROM    
            Salaries 
        WHERE     
            department = "Engineering"
    ) AS e

-- best solution
SELECT 
    ABS(
        MAX(CASE WHEN department = 'Engineering' THEN salary END) 
        -
        MAX(CASE WHEN department = 'Marketing' THEN salary END)
    ) AS salary_difference
FROM 
    Salaries

SELECT
    MAX(salary) - MIN(salary) AS salary_difference 
FROM (
    SELECT 
        MAX(salary) AS salary
    FROM 
        Salaries
    WHERE 
        department = 'Engineering' OR 
        department = 'Marketing'
    GROUP BY
        department
) AS t