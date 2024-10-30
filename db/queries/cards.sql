
INSERT INTO cards (
  card_type, card_number, card_progression, card_image
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

SELECT * FROM cards
WHERE card_ID = $1 LIMIT 1;


SELECT * FROM cards
ORDER BY card_ID
LIMIT $1
OFFSET $2;


UPDATE cards
SET 
  card_type = $2,
  card_number = $3,
  card_progression = $4,
  card_image = $5
WHERE card_ID = $1
RETURNING *;


DELETE FROM cards
WHERE card_ID = $1;