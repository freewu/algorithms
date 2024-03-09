-- 3053. Classifying Triangles by Lengths
-- Table: Triangles
-- +-------------+------+ 
-- | Column Name | Type | 
-- +-------------+------+ 
-- | A           | int  | 
-- | B           | int  |
-- | C           | int  |
-- +-------------+------+
-- (A, B, C) is the primary key for this table.
-- Each row include the lengths of each of a triangle's three sides.
-- Write a query to find the type of triangle. Output one of the following for each row:
-- Equilateral: It's a triangle with 3 sides of equal length.
-- Isosceles: It's a triangle with 2 sides of equal length.
-- Scalene: It's a triangle with 3 sides of differing lengths.
-- Not A Triangle: The given values of A, B, and C don't form a triangle.
-- Return the result table in any order.
-- The result format is in the following example.
 
-- Example 1:
-- Input: 
-- Triangles table:
-- +----+----+----+
-- | A  | B  | C  |
-- +----+----+----+
-- | 20 | 20 | 23 |
-- | 20 | 20 | 20 |
-- | 20 | 21 | 22 |
-- | 13 | 14 | 30 |
-- +----+----+----+
-- Output: 
-- +----------------+
-- | triangle_type  | 
-- +----------------+
-- | Isosceles      | 
-- | Equilateral    |
-- | Scalene        |
-- | Not A Triangle |
-- +----------------+
-- Explanation: 
-- - Values in the first row from an Isosceles triangle, because A = B.
-- - Values in the second row from an Equilateral triangle, because A = B = C.
-- - Values in the third row from an Scalene triangle, because A != B != C.
-- - Values in the fourth row cannot form a triangle, because the combined value of sides A and B is not larger than that of side C.

-- Create table If Not Exists Triangles (A int, B int, C int)
-- Truncate table Triangles
-- insert into Triangles (A, B, C) values ('20', '20', '23')
-- insert into Triangles (A, B, C) values ('20', '20', '20')
-- insert into Triangles (A, B, C) values ('20', '21', '22')
-- insert into Triangles (A, B, C) values ('13', '14', '30')


-- use case when
SELECT 
    CASE 
        WHEN A + B <= C OR B + C <= A OR C + A <= B  THEN "Not A Triangle" -- 如果两个边相加小于最大的边 不能组成三角形
        WHEN A = B AND B = C THEN "Equilateral" --  三边相等为 等边三角形
        WHEN A = B OR B = C OR A = C THEN "Isosceles" -- 两边相等为 等腰三角形
        ELSE "Scalene" -- 不等边三角形
    END AS triangle_type
FROM 
    Triangles 

-- use IF
SELECT 
    IF( A + B > C AND B + C > A AND C + A > B, -- 两边相加大于第三边才能组成三角形
        IF(A = B AND B = C, 'Equilateral', -- 三边相等为 等边三角形
                IF(A = B OR B = C OR C = A, 'Isosceles', 'Scalene') -- 两边相等为 等腰三角形
        ), 'Not A Triangle'
    ) AS triangle_type
FROM 
    Triangles
