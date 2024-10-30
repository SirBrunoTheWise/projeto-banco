
INSERT INTO users (
  alias, email, password, first_name, last_name, cpf, phone_number
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;


SELECT * FROM users
WHERE user_ID = $1 LIMIT 1;


SELECT * FROM users
ORDER BY user_ID
LIMIT $1
OFFSET $2;


UPDATE users
  set first_name = $2
WHERE user_ID = $1
RETURNING *;


DELETE FROM users
WHERE user_ID = $1;