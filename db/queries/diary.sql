
INSERT INTO diary (
  date_of, user_ID, exercise, meal, cards
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;


SELECT * FROM diary
WHERE date_of = $1 AND user_ID = $2 LIMIT 1;


SELECT * FROM diary
WHERE user_ID = $1
ORDER BY date_of DESC
LIMIT $2
OFFSET $3;


SELECT * FROM diary
WHERE user_ID = $1 
AND date_of BETWEEN $2 AND $3
ORDER BY date_of DESC;

UPDATE diary
SET 
  exercise = $3,
  meal = $4
WHERE date_of = $1 AND user_ID = $2
RETURNING *;


DELETE FROM diary
WHERE date_of = $1 AND user_ID = $2;

DELETE FROM diary
WHERE user_ID = $1;