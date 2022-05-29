-- 1076. Project Employees II
-- Table: Project
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | project_id  | int     |
-- | employee_id | int     |
-- +-------------+---------+
-- (project_id, employee_id) is the primary key of this table.
-- employee_id is a foreign key to Employee table.
-- Each row of this table indicates that the employee with employee_id is working on the project with project_id.
--  
-- Table: Employee
--
-- +------------------+---------+
-- | Column Name      | Type    |
-- +------------------+---------+
-- | employee_id      | int     |
-- | name             | varchar |
-- | experience_years | int     |
-- +------------------+---------+
-- employee_id is the primary key of this table.
-- Each row of this table contains information about one employee.
--  
-- Write an SQL query that reports all the projects that have the most employees.
-- Return the result table in any order.
-- The query result format is in the following example.
--
-- Example 1:
--
-- Input:
-- Project table:
-- +-------------+-------------+
-- | project_id  | employee_id |
-- +-------------+-------------+
-- | 1           | 1           |
-- | 1           | 2           |
-- | 1           | 3           |
-- | 2           | 1           |
-- | 2           | 4           |
-- +-------------+-------------+
-- Employee table:
-- +-------------+--------+------------------+
-- | employee_id | name   | experience_years |
-- +-------------+--------+------------------+
-- | 1           | Khaled | 3                |
-- | 2           | Ali    | 2                |
-- | 3           | John   | 1                |
-- | 4           | Doe    | 2                |
-- +-------------+--------+------------------+
-- Output:
-- +-------------+
-- | project_id  |
-- +-------------+
-- | 1           |
-- +-------------+
-- Explanation: The first project has 3 employees while the second one has 2.
--
-- Write your MySQL query statement below
SELECT
    project_id
FROM
    (
        SELECT
            p.project_id,
            COUNT(*) AS num
        FROM
            Project  AS p,
            Employee AS e
        WHERE
            p.employee_id = e.employee_id
        GROUP BY
            p.project_id
    ) AS a
WHERE
    num = (
        SELECT
            MAX(b.num) AS num
        FROM
            (
                SELECT
                    p.project_id,
                    COUNT(*) AS num
                FROM
                    Project  AS p,
                    Employee AS e
                WHERE
                    p.employee_id = e.employee_id
                GROUP BY
                    p.project_id
            ) AS b
    )

--  use rank()
SELECT
    project_id
FROM
    (
        SELECT
            project_id,
            RANK() OVER(ORDER BY COUNT(employee_id) DESC) AS r
        FROM
            Project
        GROUP BY
            project_id
    ) AS a
WHERE
    a.r = 1
