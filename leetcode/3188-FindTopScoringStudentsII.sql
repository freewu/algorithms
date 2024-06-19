-- 3188. Find Top Scoring Students II
-- Table: students
-- +-------------+----------+
-- | Column Name | Type     | 
-- +-------------+----------+
-- | student_id  | int      |
-- | name        | varchar  |
-- | major       | varchar  |
-- +-------------+----------+
-- student_id is the primary key for this table. 
-- Each row contains the student ID, student name, and their major.

-- Table: courses
-- +-------------+-------------------+
-- | Column Name | Type              |       
-- +-------------+-------------------+
-- | course_id   | int               |    
-- | name        | varchar           |      
-- | credits     | int               |           
-- | major       | varchar           |       
-- | mandatory   | enum              |      
-- +-------------+-------------------+
-- course_id is the primary key for this table. 
-- mandatory is an enum type of ('Yes', 'No').
-- Each row contains the course ID, course name, credits, major it belongs to, and whether the course is mandatory.

-- Table: enrollments
-- +-------------+----------+
-- | Column Name | Type     | 
-- +-------------+----------+
-- | student_id  | int      |
-- | course_id   | int      |
-- | semester    | varchar  |
-- | grade       | varchar  |
-- | GPA         | decimal  | 
-- +-------------+----------+
-- (student_id, course_id, semester) is the primary key (combination of columns with unique values) for this table.
-- Each row contains the student ID, course ID, semester, and grade received.
-- Write a solution to find the students who meet the following criteria:

-- Have taken all mandatory courses and at least two elective courses offered in their major.
-- Achieved a grade of A in all mandatory courses and at least B in elective courses.
-- Maintained an average GPA of at least 2.5 across all their courses (including those outside their major).
-- Return the result table ordered by student_id in ascending order.

-- Example:
-- Input:
-- students table:
-- +------------+------------------+------------------+
-- | student_id | name             | major            |
-- +------------+------------------+------------------+
-- | 1          | Alice            | Computer Science |
-- | 2          | Bob              | Computer Science |
-- | 3          | Charlie          | Mathematics      |
-- | 4          | David            | Mathematics      |
-- +------------+------------------+------------------+
-- courses table:
-- +-----------+-------------------+---------+------------------+----------+
-- | course_id | name              | credits | major            | mandatory|
-- +-----------+-------------------+---------+------------------+----------+
-- | 101       | Algorithms        | 3       | Computer Science | yes      |
-- | 102       | Data Structures   | 3       | Computer Science | yes      |
-- | 103       | Calculus          | 4       | Mathematics      | yes      |
-- | 104       | Linear Algebra    | 4       | Mathematics      | yes      |
-- | 105       | Machine Learning  | 3       | Computer Science | no       |
-- | 106       | Probability       | 3       | Mathematics      | no       |
-- | 107       | Operating Systems | 3       | Computer Science | no       |
-- | 108       | Statistics        | 3       | Mathematics      | no       |
-- +-----------+-------------------+---------+------------------+----------+
-- enrollments table:
-- +------------+-----------+-------------+-------+-----+
-- | student_id | course_id | semester    | grade | GPA |
-- +------------+-----------+-------------+-------+-----+
-- | 1          | 101       | Fall 2023   | A     | 4.0 |
-- | 1          | 102       | Spring 2023 | A     | 4.0 |
-- | 1          | 105       | Spring 2023 | A     | 4.0 |
-- | 1          | 107       | Fall 2023   | B     | 3.5 |
-- | 2          | 101       | Fall 2023   | A     | 4.0 |
-- | 2          | 102       | Spring 2023 | B     | 3.0 |
-- | 3          | 103       | Fall 2023   | A     | 4.0 |
-- | 3          | 104       | Spring 2023 | A     | 4.0 |
-- | 3          | 106       | Spring 2023 | A     | 4.0 |
-- | 3          | 108       | Fall 2023   | B     | 3.5 |
-- | 4          | 103       | Fall 2023   | B     | 3.0 |
-- | 4          | 104       | Spring 2023 | B     | 3.0 |
-- +------------+-----------+-------------+-------+-----+
 
-- Output:
-- +------------+
-- | student_id |
-- +------------+
-- | 1          |
-- | 3          |
-- +------------+
-- Explanation:
-- Alice (student_id 1) is a Computer Science major and has taken both Algorithms and Data Structures, receiving an A in both. She has also taken Machine Learning and Operating Systems as electives, receiving an A and B respectively.
-- Bob (student_id 2) is a Computer Science major but did not receive an A in all required courses.
-- Charlie (student_id 3) is a Mathematics major and has taken both Calculus and Linear Algebra, receiving an A in both. He has also taken Probability and Statistics as electives, receiving an A and B respectively.
-- David (student_id 4) is a Mathematics major but did not receive an A in all required courses.
-- Note: Output table is ordered by student_id in ascending order.

-- CREATE TABLE if not exists students (
--     student_id INT ,
--     name VARCHAR(255),
--     major VARCHAR(255)
-- )
-- CREATE TABLE if not exists courses (
--     course_id INT ,
--     name VARCHAR(255),
--     credits INT,
--     major VARCHAR(255),
--     mandatory ENUM('yes', 'no') DEFAULT 'no'
-- )
-- CREATE TABLE  if not exists enrollments (
--     student_id INT,
--     course_id INT,
--     semester VARCHAR(255),
--     grade VARCHAR(10),
-- GPA decimal(3,2)

-- );
-- Truncate table students
-- insert into students (student_id, name, major) values ('1', 'Alice', 'Computer Science')
-- insert into students (student_id, name, major) values ('2', 'Bob', 'Computer Science')
-- insert into students (student_id, name, major) values ('3', 'Charlie', 'Mathematics')
-- insert into students (student_id, name, major) values ('4', 'David', 'Mathematics')
-- Truncate table courses
-- insert into courses (course_id, name, credits, major, mandatory) values ('101', 'Algorithms', '3', 'Computer Science', 'Yes')
-- insert into courses (course_id, name, credits, major, mandatory) values ('102', 'Data Structures', '3', 'Computer Science', 'Yes')
-- insert into courses (course_id, name, credits, major, mandatory) values ('103', 'Calculus', '4', 'Mathematics', 'Yes')
-- insert into courses (course_id, name, credits, major, mandatory) values ('104', 'Linear Algebra', '4', 'Mathematics', 'Yes')
-- insert into courses (course_id, name, credits, major, mandatory) values ('105', 'Machine Learning', '3', 'Computer Science', 'No')
-- insert into courses (course_id, name, credits, major, mandatory) values ('106', 'Probability', '3', 'Mathematics', 'No')
-- insert into courses (course_id, name, credits, major, mandatory) values ('107', 'Operating Systems', '3', 'Computer Science', 'No')
-- insert into courses (course_id, name, credits, major, mandatory) values ('108', 'Statistics', '3', 'Mathematics', 'No')
-- Truncate table enrollments
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('1', '101', 'Fall 2023', 'A', '4.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('1', '102', 'Spring 2023', 'A', '4.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('1', '105', 'Spring 2023', 'A', '4.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('1', '107', 'Fall 2023', 'B', '3.5')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('2', '101', 'Fall 2023', 'A', '4.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('2', '102', 'Spring 2023', 'B', '3.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('3', '103', 'Fall 2023', 'A', '4.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('3', '104', 'Spring 2023', 'A', '4.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('3', '106', 'Spring 2023', 'A', '4.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('3', '108', 'Fall 2023', 'B', '3.5')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('4', '103', 'Fall 2023', 'B', '3.0')
-- insert into enrollments (student_id, course_id, semester, grade, GPA) values ('4', '104', 'Spring 2023', 'B', '3.0')

WITH t AS ( -- 筛选出平均 GPA 大于等于 2.5 的学生
        SELECT 
            student_id
        FROM 
            enrollments
        GROUP BY 
            student_id
        HAVING 
            AVG(GPA) >= 2.5
)
SELECT 
    student_id
FROM
    t
JOIN students USING (student_id)
JOIN courses USING (major)
LEFT JOIN enrollments USING (student_id, course_id)
GROUP BY 
    student_id
HAVING
    SUM(mandatory = 'yes' AND grade = 'A') = SUM(mandatory = 'yes') -- 已经 修完他们专业中所有的必修课程 
    AND SUM(mandatory = 'no' AND grade IS NOT NULL) = SUM(mandatory = 'no' AND grade IN ('A', 'B')) -- 在 所有必修课程 中取得等级 A 并且 选修课程 至少取得 B
    AND SUM(mandatory = 'no' AND grade IS NOT NULL) >= 2 -- 修完至少两个 选修课程
ORDER BY 
    student_id