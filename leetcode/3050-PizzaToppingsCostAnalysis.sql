-- 3050. Pizza Toppings Cost Analysis
-- Table: Toppings
-- +--------------+---------+ 
-- | Column Name  | Type    | 
-- +--------------+---------+ 
-- | topping_name | varchar | 
-- | cost         | decimal |
-- +--------------+---------+
-- topping_name is the primary key for this table.
-- Each row of this table contains topping name and the cost of the topping. 
-- Write a solution to calculate the total cost of all possible 3-topping pizza combinations from a given list of toppings. The total cost of toppings must be rounded to 2 decimal places.

-- Note:
--      Do not include the pizzas where a topping is repeated. For example, ‘Pepperoni, Pepperoni, Onion Pizza’.
--      Toppings must be listed in alphabetical order. For example, 'Chicken, Onions, Sausage'. 'Onion, Sausage, Chicken' is not acceptable.

-- Return the result table ordered by total cost in descending order and combination of toppings in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Toppings table:
-- +--------------+------+
-- | topping_name | cost |
-- +--------------+------+
-- | Pepperoni    | 0.50 |
-- | Sausage      | 0.70 |
-- | Chicken      | 0.55 |
-- | Extra Cheese | 0.40 |
-- +--------------+------+
-- Output: 
-- +--------------------------------+------------+
-- | pizza                          | total_cost | 
-- +--------------------------------+------------+
-- | Chicken,Pepperoni,Sausage      | 1.75       |  
-- | Chicken,Extra Cheese,Sausage   | 1.65       |
-- | Extra Cheese,Pepperoni,Sausage | 1.60       |
-- | Chicken,Extra Cheese,Pepperoni | 1.45       | 
-- +--------------------------------+------------+
-- Explanation: 
-- There are only four different combinations possible with the three topings:
-- - Chicken, Pepperoni, Sausage: Total cost is $1.75 (Chicken $0.55, Pepperoni $0.50, Sausage $0.70).
-- - Chicken, Extra Cheese, Sausage: Total cost is $1.65 (Chicken $0.55, Extra Cheese $0.40, Sausage $0.70).
-- - Extra Cheese, Pepperoni, Sausage: Total cost is $1.60 (Extra Cheese $0.40, Pepperoni $0.50, Sausage $0.70).
-- - Chicken, Extra Cheese, Pepperoni: Total cost is $1.45 (Chicken $0.55, Extra Cheese $0.40, Pepperoni $0.50).
-- Output table is ordered by the total cost in descending order.

-- Create table if not exists Toppings(topping_name varchar(100), cost decimal(5,2))
-- Truncate table Toppings
-- insert into Toppings (topping_name, cost) values ('Pepperoni', '0.5')
-- insert into Toppings (topping_name, cost) values ('Sausage', '0.7')
-- insert into Toppings (topping_name, cost) values ('Chicken', '0.55')
-- insert into Toppings (topping_name, cost) values ('Extra Cheese', '0.4')

SELECT
    CONCAT(t1.topping_name,",",t2.topping_name, ",",t3.topping_name) AS pizza,
    (t1.cost  + t2.cost + t3.cost ) AS total_cost 
FROM 
    Toppings AS t1,
    Toppings AS t2,
    Toppings AS t3
WHERE
    t1.topping_name < t2.topping_name AND
    t2.topping_name < t3.topping_name
ORDER BY 
    total_cost DESC, -- 以总花费 降序 排序
    t1.topping_name,t2.topping_name,t3.topping_name  -- 配料的组合 升序 排序