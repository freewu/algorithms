-- 1112. Highest Grade For Each Student
-- Table: Enrollments
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | student_id    | int     |
-- | course_id     | int     |
-- | grade         | int     |
-- +---------------+---------+
-- (student_id, course_id) is the primary key (combination of columns with unique values) of this table.
-- grade is never NULL.
 
-- Write a solution to find the highest grade with its corresponding course for each student. 
-- In case of a tie, you should find the course with the smallest course_id.
-- Return the result table ordered by student_id in ascending order.
-- The result format is in the following example.

-- Example 1:

-- Input: 
-- Enrollments table:
-- +------------+-------------------+
-- | student_id | course_id | grade |
-- +------------+-----------+-------+
-- | 2          | 2         | 95    |
-- | 2          | 3         | 95    |
-- | 1          | 1         | 90    |
-- | 1          | 2         | 99    |
-- | 3          | 1         | 80    |
-- | 3          | 2         | 75    |
-- | 3          | 3         | 82    |
-- +------------+-----------+-------+
-- Output: 
-- +------------+-------------------+
-- | student_id | course_id | grade |
-- +------------+-----------+-------+
-- | 1          | 2         | 99    |
-- | 2          | 2         | 95    |
-- | 3          | 3         | 82    |
-- +------------+-----------+-------+

SELECT 
    a.student_id AS student_id,
    MIN(a.course_id) AS course_id, -- 若科目成绩并列，取 course_id 最小的一门
    b.grade AS  grade
FROM 
    Enrollments AS a,
    (
        SELECT 
            student_id,  
            MAX(grade) AS  grade
        FROM
            Enrollments
        GROUP BY
            student_id
    ) AS b
WHERE
    a.student_id = b.student_id AND
    a.grade = b.grade
GROUP BY 
    a.student_id
ORDER BY 
    a.student_id ASC -- 查询结果需按 student_id 增序进行排序


-- best solution
SELECT 
    student_id, course_id, grade 
FROM 
    (
        SELECT 
            *, 
            --  partition by student_id 按 student_id 分区
            --  order by grade desc, course_id 按 成绩(desc) 课程ID(asc) 排序
            rank() OVER (
                PARTITION BY student_id 
                ORDER BY  grade DESC, course_id ASC
            ) 'rk'
        FROM 
            Enrollments 
    ) tb 
WHERE 
    rk = 1 
ORDER BY 
    student_id ASC