package db

import (
	"context"
)

const createCard = `-- name: CreateCard :one
INSERT INTO cards (
  card_type, card_number, card_progression, card_image
) VALUES (
  $1, $2, $3, $4
)
RETURNING card_id, card_type, card_number, card_progression, card_image
`

type CreateCardParams struct {
	CardType        int16       `json:"card_type"`
	CardNumber      int64       `json:"card_number"`
	CardProgression interface{} `json:"card_progression"`
	CardImage       []byte      `json:"card_image"`
}

func (q *Queries) CreateCard(ctx context.Context, arg CreateCardParams) (Card, error) {
	row := q.db.QueryRowContext(ctx, createCard,
		arg.CardType,
		arg.CardNumber,
		arg.CardProgression,
		arg.CardImage,
	)
	var i Card
	err := row.Scan(
		&i.CardID,
		&i.CardType,
		&i.CardNumber,
		&i.CardProgression,
		&i.CardImage,
	)
	return i, err
}

const deleteCard = `-- name: DeleteCard :exec
DELETE FROM cards
WHERE card_ID = $1
`

func (q *Queries) DeleteCard(ctx context.Context, cardID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCard, cardID)
	return err
}

const getCard = `-- name: GetCard :one
SELECT card_id, card_type, card_number, card_progression, card_image FROM cards
WHERE card_ID = $1 LIMIT 1
`

func (q *Queries) GetCard(ctx context.Context, cardID int64) (Card, error) {
	row := q.db.QueryRowContext(ctx, getCard, cardID)
	var i Card
	err := row.Scan(
		&i.CardID,
		&i.CardType,
		&i.CardNumber,
		&i.CardProgression,
		&i.CardImage,
	)
	return i, err
}

const listCards = `-- name: ListCards :many
SELECT card_id, card_type, card_number, card_progression, card_image FROM cards
ORDER BY card_ID
LIMIT $1
OFFSET $2
`

type ListCardsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCards(ctx context.Context, arg ListCardsParams) ([]Card, error) {
	rows, err := q.db.QueryContext(ctx, listCards, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.CardID,
			&i.CardType,
			&i.CardNumber,
			&i.CardProgression,
			&i.CardImage,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCard = `-- name: UpdateCard :one
UPDATE cards
SET 
  card_type = $2,
  card_number = $3,
  card_progression = $4,
  card_image = $5
WHERE card_ID = $1
RETURNING card_id, card_type, card_number, card_progression, card_image
`

type UpdateCardParams struct {
	CardID          int64       `json:"card_id"`
	CardType        int16       `json:"card_type"`
	CardNumber      int64       `json:"card_number"`
	CardProgression interface{} `json:"card_progression"`
	CardImage       []byte      `json:"card_image"`
}

func (q *Queries) UpdateCard(ctx context.Context, arg UpdateCardParams) (Card, error) {
	row := q.db.QueryRowContext(ctx, updateCard,
		arg.CardID,
		arg.CardType,
		arg.CardNumber,
		arg.CardProgression,
		arg.CardImage,
	)
	var i Card
	err := row.Scan(
		&i.CardID,
		&i.CardType,
		&i.CardNumber,
		&i.CardProgression,
		&i.CardImage,
	)
	return i, err
}
