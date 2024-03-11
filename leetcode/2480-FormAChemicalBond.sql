-- 2480. Form a Chemical Bond
-- Table: Elements
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | symbol      | varchar |
-- | type        | enum    |
-- | electrons   | int     |
-- +-------------+---------+
-- symbol is the primary key (column with unique values) for this table.
-- Each row of this table contains information of one element.
-- type is an ENUM (category) of type ('Metal', 'Nonmetal', 'Noble')
--     - If type is Noble, electrons is 0.
--     - If type is Metal, electrons is the number of electrons that one atom of this element can give.
--     - If type is Nonmetal, electrons is the number of electrons that one atom of this element needs.
 
-- Two elements can form a bond if one of them is 'Metal' and the other is 'Nonmetal'.
-- Write a solution to find all the pairs of elements that can form a bond.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Elements table:
-- +--------+----------+-----------+
-- | symbol | type     | electrons |
-- +--------+----------+-----------+
-- | He     | Noble    | 0         |
-- | Na     | Metal    | 1         |
-- | Ca     | Metal    | 2         |
-- | La     | Metal    | 3         |
-- | Cl     | Nonmetal | 1         |
-- | O      | Nonmetal | 2         |
-- | N      | Nonmetal | 3         |
-- +--------+----------+-----------+
-- Output: 
-- +-------+----------+
-- | metal | nonmetal |
-- +-------+----------+
-- | La    | Cl       |
-- | Ca    | Cl       |
-- | Na    | Cl       |
-- | La    | O        |
-- | Ca    | O        |
-- | Na    | O        |
-- | La    | N        |
-- | Ca    | N        |
-- | Na    | N        |
-- +-------+----------+
-- Explanation: 
-- Metal elements are La, Ca, and Na.
-- Nonmeal elements are Cl, O, and N.
-- Each Metal element pairs with a Nonmetal element in the output table.

-- Create table If Not Exists Elements (symbol varchar(2), type ENUM('Metal','Nonmetal','Noble'), electrons int)
-- Truncate table Elements
-- insert into Elements (symbol, type, electrons) values ('He', 'Noble', '0')
-- insert into Elements (symbol, type, electrons) values ('Na', 'Metal', '1')
-- insert into Elements (symbol, type, electrons) values ('Ca', 'Metal', '2')
-- insert into Elements (symbol, type, electrons) values ('La', 'Metal', '3')
-- insert into Elements (symbol, type, electrons) values ('Cl', 'Nonmetal', '1')
-- insert into Elements (symbol, type, electrons) values ('O', 'Nonmetal', '2')
-- insert into Elements (symbol, type, electrons) values ('N', 'Nonmetal', '3')

-- Write your MySQL query statement below
SELECT
    a.symbol AS metal,
    b.symbol AS nonmetal 
FROM
    (
        SELECT
            symbol
        FROM
            Elements 
        WHERE
            type = 'Metal'
    ) AS a,
        (
        SELECT
            symbol
        FROM
            Elements 
        WHERE
            type = 'Nonmetal'
    ) AS b

-- best solution
SELECT 
    a.symbol AS 'Metal', 
    b.symbol AS 'Nonmetal' 
FROM 
    Elements AS a, 
    Elements AS b  
WHERE 
    a.type ='Metal' AND 
    b.type = 'Nonmetal'