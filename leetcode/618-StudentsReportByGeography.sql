-- 618. Students Report By Geography
-- Table: Student
--
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | name        | varchar |
-- | continent   | varchar |
-- +-------------+---------+
-- There is no primary key for this table. It may contain duplicate rows.
-- Each row of this table indicates the name of a student and the continent they came from.
--  
-- A school has students from Asia, Europe, and America.
--
-- Write an SQL query to pivot the continent column in the Student table so that each name is sorted alphabetically and displayed underneath its corresponding continent.
-- The output headers should be America, Asia, and Europe, respectively.
-- The test cases are generated so that the student number from America is not less than either Asia or Europe.
-- The query result format is in the following example.
--
-- Example 1:
--
-- Input:
-- Student table:
-- +--------+-----------+
-- | name   | continent |
-- +--------+-----------+
-- | Jane   | America   |
-- | Pascal | Europe    |
-- | Xi     | Asia      |
-- | Jack   | America   |
-- +--------+-----------+
-- Output:
-- +---------+------+--------+
-- | America | Asia | Europe |
-- +---------+------+--------+
-- | Jack    | Xi   | Pascal |
-- | Jane    | null | null   |
-- +---------+------+--------+
--  
-- Follow up: If it is unknown which continent has the most students, could you write a query to generate the student report?
--
-- Write your MySQL query statement below
SELECT
    a.America,
    b.Asia,
    c.Europe
FROM
    (
        SELECT
            name AS America,
            ROW_NUMBER() OVER(ORDER BY name ASC) AS r
        FROM
            Student
        WHERE
            continent  = 'America'
    ) AS a
LEFT JOIN
    (
        SELECT
            name AS Asia,
            ROW_NUMBER() OVER(ORDER BY name ASC) AS r
        FROM
            Student
        WHERE
            continent  = 'Asia'
    ) AS b
ON
    a.r = b.r
LEFT JOIN
    (
        SELECT
            name AS Europe,
            ROW_NUMBER() OVER(ORDER BY name ASC) AS r
        FROM
            Student
        WHERE
            continent  = 'Europe'
    ) AS c
ON
    a.r = c.r