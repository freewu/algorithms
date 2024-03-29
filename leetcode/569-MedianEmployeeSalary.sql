-- 569. Median Employee Salary
-- Table: Employee
--
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | id           | int     |
-- | company      | varchar |
-- | salary       | int     |
-- +--------------+---------+
-- id is the primary key column for this table.
-- Each row of this table indicates the company and the salary of one employee.
--  
-- Write an SQL query to find the median salary of each company.
-- Return the result table in any order.
-- The query result format is in the following example.
--
-- Example 1:
--
-- Input:
-- Employee table:
-- +----+---------+--------+
-- | id | company | salary |
-- +----+---------+--------+
-- | 1  | A       | 2341   |
-- | 2  | A       | 341    |
-- | 3  | A       | 15     |
-- | 4  | A       | 15314  |
-- | 5  | A       | 451    |
-- | 6  | A       | 513    |
-- | 7  | B       | 15     |
-- | 8  | B       | 13     |
-- | 9  | B       | 1154   |
-- | 10 | B       | 1345   |
-- | 11 | B       | 1221   |
-- | 12 | B       | 234    |
-- | 13 | C       | 2345   |
-- | 14 | C       | 2645   |
-- | 15 | C       | 2645   |
-- | 16 | C       | 2652   |
-- | 17 | C       | 65     |
-- +----+---------+--------+
-- Output:
-- +----+---------+--------+
-- | id | company | salary |
-- +----+---------+--------+
-- | 5  | A       | 451    |
-- | 6  | A       | 513    |
-- | 12 | B       | 234    |
-- | 9  | B       | 1154   |
-- | 14 | C       | 2645   |
-- +----+---------+--------+
--  
--
-- Follow up: Could you solve it without using any built-in or window functions?
--

--Write your MySQL query statement below
-- without using any built-in or window functions
SELECT 
    Id, 
    Company, 
    Salary
FROM 
    Employee
WHERE 
    Id IN (
        SELECT 
            e1.Id
        FROM 
            Employee e1, 
            Employee e2
        WHERE 
            e1.Company = e2.Company
        GROUP BY 
            e1.Id
        HAVING 
            SUM(CASE WHEN e1.Salary >= e2.Salary THEN 1 ELSE 0 END) >= COUNT(*)/2 AND 
            SUM(CASE WHEN e1.Salary <= e2.Salary THEN 1 ELSE 0 END) >= COUNT(*)/2)
GROUP BY 
    Company, Salary
ORDER BY 
    Company

-- row_number
SELECT 
    Id,
    Company,
    Salary 
FROM 
(
    SELECT 
        Id,
        Company,
        Salary,
        ROW_NUMBER() OVER(PARTITION BY Company ORDER BY Salary) AS rnk,
        COUNT(Salary) OVER(PARTITION BY Company) AS cnt 
    FROM 
        Employee 
) AS t 
WHERE 
    rnk IN ( cnt / 2,cnt / 2 + 1,cnt / 2+ 0.5)