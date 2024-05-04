-- 3057. Employees Project Allocation
-- Table: Project
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | project_id  | int     |
-- | employee_id | int     |
-- | workload    | int     |
-- +-------------+---------+
-- employee_id is the primary key (column with unique values) of this table.
-- employee_id is a foreign key (reference column) to Employee table.
-- Each row of this table indicates that the employee with employee_id is working on the project with project_id and the workload of the project.

-- Table: Employees
-- +------------------+---------+
-- | Column Name      | Type    |
-- +------------------+---------+
-- | employee_id      | int     |
-- | name             | varchar |
-- | team             | varchar |
-- +------------------+---------+
-- employee_id is the primary key (column with unique values) of this table.
-- Each row of this table contains information about one employee.
-- Write a solution to find the employees who are allocated to projects with a workload that exceeds the average workload of all employees for their respective teams

-- Return the result table ordered by employee_id, project_id in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Project table:
-- +-------------+-------------+----------+
-- | project_id  | employee_id | workload |
-- +-------------+-------------+----------+
-- | 1           | 1           |  45      |
-- | 1           | 2           |  90      | 
-- | 2           | 3           |  12      |
-- | 2           | 4           |  68      |
-- +-------------+-------------+----------+
-- Employees table:
-- +-------------+--------+------+
-- | employee_id | name   | team |
-- +-------------+--------+------+
-- | 1           | Khaled | A    |
-- | 2           | Ali    | B    |
-- | 3           | John   | B    |
-- | 4           | Doe    | A    |
-- +-------------+--------+------+
-- Output: 
-- +-------------+------------+---------------+------------------+
-- | employee_id | project_id | employee_name | project_workload |
-- +-------------+------------+---------------+------------------+  
-- | 2           | 1          | Ali           | 90               | 
-- | 4           | 2          | Doe           | 68               | 
-- +-------------+------------+---------------+------------------+
-- Explanation: 
-- - Employee with ID 1 has a project workload of 45 and belongs to Team A, where the average workload is 56.50. Since his project workload does not exceed the team's average workload, he will be excluded.
-- - Employee with ID 2 has a project workload of 90 and belongs to Team B, where the average workload is 51.00. Since his project workload does exceed the team's average workload, he will be included.
-- - Employee with ID 3 has a project workload of 12 and belongs to Team B, where the average workload is 51.00. Since his project workload does not exceed the team's average workload, he will be excluded.
-- - Employee with ID 4 has a project workload of 68 and belongs to Team A, where the average workload is 56.50. Since his project workload does exceed the team's average workload, he will be included.
-- Result table orderd by employee_id, project_id in ascending order.

-- Create table If Not Exists Project (project_id int, employee_id int, workload int)
-- Create table If Not Exists Employees (employee_id int, name varchar(20), team varchar(20))
-- Truncate table Project
-- insert into Project (project_id, employee_id, workload) values ('1', '1', '45')
-- insert into Project (project_id, employee_id, workload) values ('1', '2', '90')
-- insert into Project (project_id, employee_id, workload) values ('2', '3', '12')
-- insert into Project (project_id, employee_id, workload) values ('2', '4', '68')
-- Truncate table Employees
-- insert into Employees (employee_id, name, team) values ('1', 'Khaled', 'A')
-- insert into Employees (employee_id, name, team) values ('2', 'Ali', 'B')
-- insert into Employees (employee_id, name, team) values ('3', 'John', 'B')
-- insert into Employees (employee_id, name, team) values ('4', 'Doe', 'A')

WITH t AS ( -- 汇总信息
    SELECT 
        p.project_id,
        p.workload,
        e.*
    FROM 
        Employees AS e
    LEFT JOIN
        Project AS p 
    ON 
        e.employee_id  = p.employee_id 
),
a AS (
    SELECT AVG(workload) AS avg_workload,team FROM t GROUP BY team
)
-- SELECT * FROM t
-- | project_id | workload | employee_id | name   | team |
-- | ---------- | -------- | ----------- | ------ | ---- |
-- | 1          | 45       | 1           | Khaled | A    |
-- | 1          | 90       | 2           | Ali    | B    |
-- | 2          | 12       | 3           | John   | B    |
-- | 2          | 68       | 4           | Doe    | A    |
-- 每个团队的平均工作量
-- SELECT AVG(workload) AS avg_workload FROM t GROUP BY team
-- | avg_workload |
-- | ------------- |
-- | 56.5          |
-- | 51            |

SELECT 
    t.employee_id,
    t.project_id,
    t.name AS employee_name,
    t.workload AS project_workload 
FROM 
    t
JOIN 
    a 
ON 
    t.team = a.team AND t.workload > a.avg_workload -- 分配给项目的工作量 超过各自团队 所有员工 平均工作量 的 员工
ORDER BY 
    t.employee_id,t.project_id -- 以 employee_id，project_id 升序 排序