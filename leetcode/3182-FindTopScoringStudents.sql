-- 3182. Find Top Scoring Students
-- Table: students
-- +-------------+----------+
-- | Column Name | Type     | 
-- +-------------+----------+
-- | student_id  | int      |
-- | name        | varchar  |
-- | major       | varchar  |
-- +-------------+----------+
-- student_id is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the student ID, student name, and their major.

-- Table: courses
-- +-------------+----------+
-- | Column Name | Type     | 
-- +-------------+----------+
-- | course_id   | int      |
-- | name        | varchar  |
-- | credits     | int      |
-- | major       | varchar  |
-- +-------------+----------+
-- course_id is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the course ID, course name, the number of credits for the course, and the major it belongs to.

-- Table: enrollments
-- +-------------+----------+
-- | Column Name | Type     | 
-- +-------------+----------+
-- | student_id  | int      |
-- | course_id   | int      |
-- | semester    | varchar  |
-- | grade       | varchar  |
-- +-------------+----------+
-- (student_id, course_id, semester) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table contains the student ID, course ID, semester, and grade received.
-- Write a solution to find the students who have taken all courses offered in their major and have achieved a grade of A in all these courses.
-- Return the result table ordered by student_id in ascending order.
-- The result format is in the following example.

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
-- +-----------+-----------------+---------+------------------+
-- | course_id | name            | credits | major            |
-- +-----------+-----------------+---------+------------------+
-- | 101       | Algorithms      | 3       | Computer Science |
-- | 102       | Data Structures | 3       | Computer Science |
-- | 103       | Calculus        | 4       | Mathematics      |
-- | 104       | Linear Algebra  | 4       | Mathematics      |
-- +-----------+-----------------+---------+------------------+
-- enrollments table:
-- +------------+-----------+----------+-------+
-- | student_id | course_id | semester | grade |
-- +------------+-----------+----------+-------+
-- | 1          | 101       | Fall 2023| A     |
-- | 1          | 102       | Fall 2023| A     |
-- | 2          | 101       | Fall 2023| B     |
-- | 2          | 102       | Fall 2023| A     |
-- | 3          | 103       | Fall 2023| A     |
-- | 3          | 104       | Fall 2023| A     |
-- | 4          | 103       | Fall 2023| A     |
-- | 4          | 104       | Fall 2023| B     |
-- +------------+-----------+----------+-------+
-- Output:
-- +------------+
-- | student_id |
-- +------------+
-- | 1          |
-- | 3          |
-- +------------+
-- Explanation:
-- Alice (student_id 1) is a Computer Science major and has taken both "Algorithms" and "Data Structures", receiving an 'A' in both.
-- Bob (student_id 2) is a Computer Science major but did not receive an 'A' in all required courses.
-- Charlie (student_id 3) is a Mathematics major and has taken both "Calculus" and "Linear Algebra", receiving an 'A' in both.
-- David (student_id 4) is a Mathematics major but did not receive an 'A' in all required courses.
-- Note: Output table is ordered by student_id in ascending order.

-- CREATE TABLE if not exists students (
--     student_id INT ,
--     name VARCHAR(255),
--     major VARCHAR(255)
-- )
-- CREATE TABLE  if not exists courses (
--     course_id INT ,
--     name VARCHAR(255),
--     credits INT,
--     major VARCHAR(255)
-- )

-- CREATE TABLE  if not exists enrollments (
--     student_id INT,
--     course_id INT,
--     semester VARCHAR(255),
--     grade VARCHAR(10)

-- );
-- Truncate table students
-- insert into students (student_id, name, major) values ('1', 'Alice', 'Computer Science')
-- insert into students (student_id, name, major) values ('2', 'Bob', 'Computer Science')
-- insert into students (student_id, name, major) values ('3', 'Charlie', 'Mathematics')
-- insert into students (student_id, name, major) values ('4', 'David', 'Mathematics')
-- Truncate table courses
-- insert into courses (course_id, name, credits, major) values ('101', 'Algorithms', '3', 'Computer Science')
-- insert into courses (course_id, name, credits, major) values ('102', 'Data Structures', '3', 'Computer Science')
-- insert into courses (course_id, name, credits, major) values ('103', 'Calculus', '4', 'Mathematics')
-- insert into courses (course_id, name, credits, major) values ('104', 'Linear Algebra', '4', 'Mathematics')
-- Truncate table enrollments
-- insert into enrollments (student_id, course_id, semester, grade) values ('1', '101', 'Fall 2023', 'A')
-- insert into enrollments (student_id, course_id, semester, grade) values ('1', '102', 'Fall 2023', 'A')
-- insert into enrollments (student_id, course_id, semester, grade) values ('2', '101', 'Fall 2023', 'B')
-- insert into enrollments (student_id, course_id, semester, grade) values ('2', '102', 'Fall 2023', 'A')
-- insert into enrollments (student_id, course_id, semester, grade) values ('3', '103', 'Fall 2023', 'A')
-- insert into enrollments (student_id, course_id, semester, grade) values ('3', '104', 'Fall 2023', 'A')
-- insert into enrollments (student_id, course_id, semester, grade) values ('4', '103', 'Fall 2023', 'A')
-- insert into enrollments (student_id, course_id, semester, grade) values ('4', '104', 'Fall 2023', 'B')

# Write your MySQL query statement below
-- SELECT
--     s.student_id,
--     -- c.course_id,
--     -- e.grade,
--     count(*) AS cnt,
--     count(IF(e.grade = "A", 1, null)) AS cnta
-- FROM
--     students AS s
-- LEFT JOIN 
--     courses AS c 
-- ON 
--     c.major = s.major
-- LEFT JOIN
--     enrollments AS e
-- ON
--     e.student_id = s.student_id AND c.course_id = e.course_id
-- GROUP BY 
--     s.student_id

SELECT 
    student_id
FROM 
(
    SELECT
        s.student_id,
        -- c.course_id,
        -- e.grade,
        count(*) AS cnt,
        count(IF(e.grade = "A", 1, null)) AS cnta
    FROM
        students AS s
    LEFT JOIN 
        courses AS c 
    ON 
        c.major = s.major
    LEFT JOIN
        enrollments AS e
    ON
        e.student_id = s.student_id AND c.course_id = e.course_id
    GROUP BY 
        s.student_id
) AS t
WHERE
    cnt = cnta
ORDER BY 
    student_id 