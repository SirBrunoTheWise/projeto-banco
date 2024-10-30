package db

import (
	"context"
	"time"
)

const createDiaryEntry = `-- name: CreateDiaryEntry :one
INSERT INTO diary (
  date_of, user_ID, exercise, meal, cards
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING date_of, user_id, exercise, meal, cards
`

type CreateDiaryEntryParams struct {
	DateOf   time.Time `json:"date_of"`
	UserID   int64     `json:"user_id"`
	Exercise int64     `json:"exercise"`
	Meal     int64     `json:"meal"`
	Cards    int64     `json:"cards"`
}

func (q *Queries) CreateDiaryEntry(ctx context.Context, arg CreateDiaryEntryParams) (Diary, error) {
	row := q.db.QueryRowContext(ctx, createDiaryEntry,
		arg.DateOf,
		arg.UserID,
		arg.Exercise,
		arg.Meal,
		arg.Cards,
	)
	var i Diary
	err := row.Scan(
		&i.DateOf,
		&i.UserID,
		&i.Exercise,
		&i.Meal,
		&i.Cards,
	)
	return i, err
}

const deleteDiaryEntry = `-- name: DeleteDiaryEntry :exec
DELETE FROM diary
WHERE date_of = $1 AND user_ID = $2
`

type DeleteDiaryEntryParams struct {
	DateOf time.Time `json:"date_of"`
	UserID int64     `json:"user_id"`
}

func (q *Queries) DeleteDiaryEntry(ctx context.Context, arg DeleteDiaryEntryParams) error {
	_, err := q.db.ExecContext(ctx, deleteDiaryEntry, arg.DateOf, arg.UserID)
	return err
}

const deleteUserDiaryEntries = `-- name: DeleteUserDiaryEntries :exec
DELETE FROM diary
WHERE user_ID = $1
`

func (q *Queries) DeleteUserDiaryEntries(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserDiaryEntries, userID)
	return err
}

const getDiaryEntry = `-- name: GetDiaryEntry :one
SELECT date_of, user_id, exercise, meal, cards FROM diary
WHERE date_of = $1 AND user_ID = $2 LIMIT 1
`

type GetDiaryEntryParams struct {
	DateOf time.Time `json:"date_of"`
	UserID int64     `json:"user_id"`
}

func (q *Queries) GetDiaryEntry(ctx context.Context, arg GetDiaryEntryParams) (Diary, error) {
	row := q.db.QueryRowContext(ctx, getDiaryEntry, arg.DateOf, arg.UserID)
	var i Diary
	err := row.Scan(
		&i.DateOf,
		&i.UserID,
		&i.Exercise,
		&i.Meal,
		&i.Cards,
	)
	return i, err
}

const listDiaryEntries = `-- name: ListDiaryEntries :many
SELECT date_of, user_id, exercise, meal, cards FROM diary
WHERE user_ID = $1
ORDER BY date_of DESC
LIMIT $2
OFFSET $3
`

type ListDiaryEntriesParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDiaryEntries(ctx context.Context, arg ListDiaryEntriesParams) ([]Diary, error) {
	rows, err := q.db.QueryContext(ctx, listDiaryEntries, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Diary
	for rows.Next() {
		var i Diary
		if err := rows.Scan(
			&i.DateOf,
			&i.UserID,
			&i.Exercise,
			&i.Meal,
			&i.Cards,
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

const listDiaryEntriesByDateRange = `-- name: ListDiaryEntriesByDateRange :many
SELECT date_of, user_id, exercise, meal, cards FROM diary
WHERE user_ID = $1 
AND date_of BETWEEN $2 AND $3
ORDER BY date_of DESC
`

type ListDiaryEntriesByDateRangeParams struct {
	UserID   int64     `json:"user_id"`
	DateOf   time.Time `json:"date_of"`
	DateOf_2 time.Time `json:"date_of_2"`
}

func (q *Queries) ListDiaryEntriesByDateRange(ctx context.Context, arg ListDiaryEntriesByDateRangeParams) ([]Diary, error) {
	rows, err := q.db.QueryContext(ctx, listDiaryEntriesByDateRange, arg.UserID, arg.DateOf, arg.DateOf_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Diary
	for rows.Next() {
		var i Diary
		if err := rows.Scan(
			&i.DateOf,
			&i.UserID,
			&i.Exercise,
			&i.Meal,
			&i.Cards,
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

const updateDiaryEntry = `-- name: UpdateDiaryEntry :one
UPDATE diary
SET 
  exercise = $3,
  meal = $4
WHERE date_of = $1 AND user_ID = $2
RETURNING date_of, user_id, exercise, meal, cards
`

type UpdateDiaryEntryParams struct {
	DateOf   time.Time `json:"date_of"`
	UserID   int64     `json:"user_id"`
	Exercise int64     `json:"exercise"`
	Meal     int64     `json:"meal"`
}

func (q *Queries) UpdateDiaryEntry(ctx context.Context, arg UpdateDiaryEntryParams) (Diary, error) {
	row := q.db.QueryRowContext(ctx, updateDiaryEntry,
		arg.DateOf,
		arg.UserID,
		arg.Exercise,
		arg.Meal,
	)
	var i Diary
	err := row.Scan(
		&i.DateOf,
		&i.UserID,
		&i.Exercise,
		&i.Meal,
		&i.Cards,
	)
	return i, err
}
