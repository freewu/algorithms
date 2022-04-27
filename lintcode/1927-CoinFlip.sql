-- 1927 · Coin Flip
-- # Description
-- Given a coins table, the side attribute represents the front and back of the coin, p represents the front side (positive), and n represents the back side (negative)
-- Swap all p and n values (for example, change all p values to n and vice versa).
-- It is required to use only one UPDATE statement, and no intermediate temporary tables.

-- Table definition: coins

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- side	char	Sides of the coin
-- Tip:

-- You can solve this problem by learning UPDATE statement and conditional judgment related knowledge in SQL statements.
-- Only one UPDATE statement can be written, please do not write any SELECT statement.
-- We will individually verify whether the data in the database has been modified to the following information.
-- Example
-- Example 1:

-- Table content: coins

-- id	side
-- 1	p
-- 2	n
-- 3	p
-- 4	n
-- After running the SQL statement you have written, you will get the following table:

-- id	side
-- 1	n
-- 2	p
-- 3	n
-- 4	p
-- Example 2:

-- Table content: coins

-- id	side
-- 1	p
-- 2	n
-- After running the SQL statement you have written, you will get the following table:

-- id	side
-- 1	n
-- 2	p

-- IF(expr1,expr2,expr3)，如果expr1的值为true，则返回expr2的值，如果expr1的值为false，则返回expr3的值

UPDATE 
	coins
SET
	side = IF(side = 'n','p','n')