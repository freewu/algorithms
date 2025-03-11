-- 3482. Analyze Organization Hierarchy
-- Table: Employees
-- +----------------+---------+
-- | Column Name    | Type    | 
-- +----------------+---------+
-- | employee_id    | int     |
-- | employee_name  | varchar |
-- | manager_id     | int     |
-- | salary         | int     |
-- | department     | varchar |
-- +----------------+----------+
-- employee_id is the unique key for this table.
-- Each row contains information about an employee, including their ID, name, their manager's ID, salary, and department.
-- manager_id is null for the top-level manager (CEO).
-- Write a solution to analyze the organizational hierarchy and answer the following:
--     1. Hierarchy Levels: For each employee, determine their level in the organization (CEO is level 1, employees reporting directly to the CEO are level 2, and so on).
--     2. Team Size: For each employee who is a manager, count the total number of employees under them (direct and indirect reports).
--     3. Salary Budget: For each manager, calculate the total salary budget they control (sum of salaries of all employees under them, including indirect reports, plus their own salary).

-- Return the result table ordered by the result ordered by level in ascending order, then by budget in descending order, and finally by employee_name in ascending order.

-- The result format is in the following example.

-- Example:
-- Input:
-- Employees table:
-- +-------------+---------------+------------+--------+-------------+
-- | employee_id | employee_name | manager_id | salary | department  |
-- +-------------+---------------+------------+--------+-------------+
-- | 1           | Alice         | null       | 12000  | Executive   |
-- | 2           | Bob           | 1          | 10000  | Sales       |
-- | 3           | Charlie       | 1          | 10000  | Engineering |
-- | 4           | David         | 2          | 7500   | Sales       |
-- | 5           | Eva           | 2          | 7500   | Sales       |
-- | 6           | Frank         | 3          | 9000   | Engineering |
-- | 7           | Grace         | 3          | 8500   | Engineering |
-- | 8           | Hank          | 4          | 6000   | Sales       |
-- | 9           | Ivy           | 6          | 7000   | Engineering |
-- | 10          | Judy          | 6          | 7000   | Engineering |
-- +-------------+---------------+------------+--------+-------------+
-- Output:
-- +-------------+---------------+-------+-----------+--------+
-- | employee_id | employee_name | level | team_size | budget |
-- +-------------+---------------+-------+-----------+--------+
-- | 1           | Alice         | 1     | 9         | 84500  |
-- | 3           | Charlie       | 2     | 4         | 41500  |
-- | 2           | Bob           | 2     | 3         | 31000  |
-- | 6           | Frank         | 3     | 2         | 23000  |
-- | 4           | David         | 3     | 1         | 13500  |
-- | 7           | Grace         | 3     | 0         | 8500   |
-- | 5           | Eva           | 3     | 0         | 7500   |
-- | 9           | Ivy           | 4     | 0         | 7000   |
-- | 10          | Judy          | 4     | 0         | 7000   |
-- | 8           | Hank          | 4     | 0         | 6000   |
-- +-------------+---------------+-------+-----------+--------+
-- Explanation:
-- Organization Structure:
-- Alice (ID: 1) is the CEO (level 1) with no manager
-- Bob (ID: 2) and Charlie (ID: 3) report directly to Alice (level 2)
-- David (ID: 4), Eva (ID: 5) report to Bob, while Frank (ID: 6) and Grace (ID: 7) report to Charlie (level 3)
-- Hank (ID: 8) reports to David, and Ivy (ID: 9) and Judy (ID: 10) report to Frank (level 4)
-- Level Calculation:
-- The CEO (Alice) is at level 1
-- Each subsequent level of management adds 1 to the level
-- Team Size Calculation:
-- Alice has 9 employees under her (the entire company except herself)
-- Bob has 3 employees (David, Eva, and Hank)
-- Charlie has 4 employees (Frank, Grace, Ivy, and Judy)
-- David has 1 employee (Hank)
-- Frank has 2 employees (Ivy and Judy)
-- Eva, Grace, Hank, Ivy, and Judy have no direct reports (team_size = 0)
-- Budget Calculation:
-- Alice's budget: Her salary (12000) + all employees' salaries (72500) = 84500
-- Charlie's budget: His salary (10000) + Frank's budget (23000) + Grace's salary (8500) = 41500
-- Bob's budget: His salary (10000) + David's budget (13500) + Eva's salary (7500) = 31000
-- Frank's budget: His salary (9000) + Ivy's salary (7000) + Judy's salary (7000) = 23000
-- David's budget: His salary (7500) + Hank's salary (6000) = 13500
-- Employees with no direct reports have budgets equal to their own salary

-- Note:
--     The result is ordered first by level in ascending order
--     Within the same level, employees are ordered by budget in descending order then by name in ascending order  

-- CREATE TABLE if not exists Employees (
--     employee_id INT,
--     employee_name VARCHAR(100),
--     manager_id INT,
--     salary INT,
--     department VARCHAR(50)
-- )
-- Truncate table Employees
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('1', 'Alice', NULL, '12000', 'Executive')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('2', 'Bob', '1', '10000', 'Sales')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('3', 'Charlie', '1', '10000', 'Engineering')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('4', 'David', '2', '7500', 'Sales')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('5', 'Eva', '2', '7500', 'Sales')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('6', 'Frank', '3', '9000', 'Engineering')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('7', 'Grace', '3', '8500', 'Engineering')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('8', 'Hank', '4', '6000', 'Sales')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('9', 'Ivy', '6', '7000', 'Engineering')
-- insert into Employees (employee_id, employee_name, manager_id, salary, department) values ('10', 'Judy', '6', '7000', 'Engineering')

-- Write your MySQL query statement below
WITH RECURSIVE org_hierarchy(orig_employee_id, orig_employee_name, employee_id, employee_name, manager_id, salary, org_level) AS
(
    SELECT 
        employee_id AS orig_employee_id,
        employee_name AS orig_employee_name,
        employee_id,
        employee_name,
        manager_id,
        salary,
        1 AS org_level
    FROM 
        Employees
    UNION ALL
    SELECT 
        o.orig_employee_id,
        o.orig_employee_name,
        e.employee_id,
        e.employee_name,
        e.manager_id,
        e.salary,
        o.org_level + 1
    FROM 
        org_hierarchy AS o, Employees AS e
    WHERE 
        e.manager_id = o.employee_id
),
CEO_hierarchy AS (
    SELECT 
        o.employee_id as SUB_employee_id,
        o.employee_name,
        o.org_level as sub_level
    FROM 
        org_hierarchy AS o, Employees AS e
    WHERE
        o.orig_employee_id = e.employee_id AND
        e.manager_id is null
)
SELECT
    o.ORIG_EMPLOYEE_ID as employee_id,
    o.ORIG_EMPLOYEE_name as employee_name,
    m.sub_level as "level",
    COUNT(*) - 1 as team_size,
    SUM(o.salary) as budget
FROM 
    org_hierarchy AS o, CEO_hierarchy AS m
WHERE 
    o.ORIG_EMPLOYEE_ID = m.SUB_employee_id
GROUP BY 
    o.ORIG_EMPLOYEE_ID, o.ORIG_EMPLOYEE_name, m.sub_level
ORDER BY 
    3 asc, 5 desc, 2