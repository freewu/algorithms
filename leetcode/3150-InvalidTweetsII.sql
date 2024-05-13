-- 3150. Invalid Tweets II
-- Table: Tweets
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | tweet_id       | int     |
-- | content        | varchar |
-- +----------------+---------+
-- tweet_id is the primary key (column with unique values) for this table.
-- This table contains all the tweets in a social media app.

-- Write a solution to find invalid tweets. A tweet is considered invalid if it meets any of the following criteria:
--     It exceeds 140 characters in length.
--     It has more than 3 mentions.
--     It includes more than 3 hashtags.

-- Return the result table ordered by tweet_id in ascending order.
-- The result format is in the following example.

-- Example:
-- Input:
-- Tweets table:
-- +----------+-----------------------------------------------------------------------------------+
-- | tweet_id | content                                                                           |
-- +----------+-----------------------------------------------------------------------------------+
-- | 1        | Traveling, exploring, and living my best life @JaneSmith @SaraJohnson @LisaTaylor |
-- |          | @MikeBrown #Foodie #Fitness #Learning                                             | 
-- | 2        | Just had the best dinner with friends! #Foodie #Friends #Fun                      |
-- | 4        | Working hard on my new project #Work #Goals #Productivity #Fun                    |
-- +----------+-----------------------------------------------------------------------------------+
-- Output:
-- +----------+
-- | tweet_id |
-- +----------+
-- | 1        |
-- | 4        |
-- +----------+
-- Explanation:
-- tweet_id 1 contains 4 mentions.
-- tweet_id 4 contains 4 hashtags.
-- Output table is ordered by tweet_id in ascending order.

-- Create table If Not Exists Tweets(tweet_id int, content varchar(500))
-- Truncate table Tweets
-- insert into Tweets (tweet_id, content) values ('1', 'What an amazing meal @MaxPower @AlexJones @JohnDoe #Learning #Fitness #Love')
-- insert into Tweets (tweet_id, content) values ('2', 'Learning something new every day @AnnaWilson #Learning #Foodie')
-- insert into Tweets (tweet_id, content) values ('3', 'Never been happier about today's achievements @SaraJohnson @JohnDoe @AnnaWilson #Fashion')
-- insert into Tweets (tweet_id, content) values ('4', 'Traveling, exploring, and living my best life @JaneSmith @JohnDoe @ChrisAnderson @AlexJones #WorkLife #Travel')
-- insert into Tweets (tweet_id, content) values ('5', 'Work hard, play hard, and cherish every moment @AlexJones #Fashion #Foodie')
-- insert into Tweets (tweet_id, content) values ('6', 'Never been happier about today's achievements @ChrisAnderson #Fashion #WorkLife')
-- insert into Tweets (tweet_id, content) values ('7', 'So grateful for today's experiences @AnnaWilson @LisaTaylor @ChrisAnderson @MikeBrown #Fashion #HappyDay #WorkLife #Nature')
-- insert into Tweets (tweet_id, content) values ('8', 'What an amazing meal @EmilyClark @AlexJones @MikeBrown #Fitness')
-- insert into Tweets (tweet_id, content) values ('9', 'Learning something new every day @EmilyClark @AnnaWilson @MaxPower #Travel')
-- insert into Tweets (tweet_id, content) values ('10', 'So grateful for today's experiences @ChrisAnderson #Nature')
-- insert into Tweets (tweet_id, content) values ('11', 'So grateful for today's experiences @AlexJones #Art #WorkLife')
-- insert into Tweets (tweet_id, content) values ('12', 'Learning something new every day @JaneSmith @MikeBrown #Travel')
-- insert into Tweets (tweet_id, content) values ('13', 'What an amazing meal @EmilyClark @JohnDoe @LisaTaylor @MaxPower #Foodie #Fitness')
-- insert into Tweets (tweet_id, content) values ('14', 'Work hard, play hard, and cherish every moment @LisaTaylor @SaraJohnson @MaxPower @ChrisAnderson #TechLife #Nature #Music')
-- insert into Tweets (tweet_id, content) values ('15', 'What a beautiful day it is @EmilyClark @MaxPower @SaraJohnson #Fashion')
-- insert into Tweets (tweet_id, content) values ('16', 'What a beautiful day it is @AnnaWilson @JaneSmith #Fashion #Love #TechLife')

-- Write your MySQL query statement below
SELECT 
    tweet_id 
FROM
    Tweets 
WHERE
    LENGTH(content) > 140 OR -- It includes more than 3 hashtags.
    (LENGTH(content) - LENGTH(REPLACE(content, '#', ''))) > 3 OR -- It includes more than 3 hashtags.
    (LENGTH(content) - LENGTH(REPLACE(content, '@', ''))) > 3 -- It has more than 3 mentions. 
ORDER BY
    tweet_id 