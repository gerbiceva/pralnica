// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package pralnicaDb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addReservation = `-- name: AddReservation :one
INSERT INTO reservations (
    uuid, 
    date, 
    termin, 
    washer
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, uuid, date, termin, washer
`

type AddReservationParams struct {
	Uuid   pgtype.UUID
	Date   pgtype.Date
	Termin int32
	Washer int32
}

func (q *Queries) AddReservation(ctx context.Context, arg AddReservationParams) (Reservation, error) {
	row := q.db.QueryRow(ctx, addReservation,
		arg.Uuid,
		arg.Date,
		arg.Termin,
		arg.Washer,
	)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Date,
		&i.Termin,
		&i.Washer,
	)
	return i, err
}

const banUser = `-- name: BanUser :exec
UPDATE users
SET disabled = TRUE
WHERE uuid = $1
`

func (q *Queries) BanUser(ctx context.Context, uuid pgtype.UUID) error {
	_, err := q.db.Exec(ctx, banUser, uuid)
	return err
}

const confirmUser = `-- name: ConfirmUser :exec
UPDATE users
SET confirmed = TRUE
WHERE uuid = $1
`

func (q *Queries) ConfirmUser(ctx context.Context, uuid pgtype.UUID) error {
	_, err := q.db.Exec(ctx, confirmUser, uuid)
	return err
}

const deleteReservation = `-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE id = $1
`

func (q *Queries) DeleteReservation(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteReservation, id)
	return err
}

const editUser = `-- name: EditUser :exec
UPDATE users
SET phone = $2,
    name = $3,
    surname = $4,
    room = $5,
    email = $6,
    role = $7
WHERE uuid = $1
`

type EditUserParams struct {
	Uuid    pgtype.UUID
	Phone   string
	Name    string
	Surname string
	Room    int32
	Email   string
	Role    string
}

func (q *Queries) EditUser(ctx context.Context, arg EditUserParams) error {
	_, err := q.db.Exec(ctx, editUser,
		arg.Uuid,
		arg.Phone,
		arg.Name,
		arg.Surname,
		arg.Room,
		arg.Email,
		arg.Role,
	)
	return err
}

const getReservationsByMonth = `-- name: GetReservationsByMonth :many
SELECT
    r.id,
    r.uuid,
    r.date,
    r.termin,
    r.washer
FROM
    reservations r
WHERE
    r.date >= DATE_TRUNC('month', $1::date)
    AND r.date < DATE_TRUNC('month', $1::date) + INTERVAL '1 month'
`

func (q *Queries) GetReservationsByMonth(ctx context.Context, dollar_1 pgtype.Date) ([]Reservation, error) {
	rows, err := q.db.Query(ctx, getReservationsByMonth, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reservation
	for rows.Next() {
		var i Reservation
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Date,
			&i.Termin,
			&i.Washer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT uuid, phone, name, surname, room, email, disabled, confirmed, role, created_at, updated_at
FROM users
WHERE uuid = $1
`

func (q *Queries) GetUser(ctx context.Context, uuid pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, uuid)
	var i User
	err := row.Scan(
		&i.Uuid,
		&i.Phone,
		&i.Name,
		&i.Surname,
		&i.Room,
		&i.Email,
		&i.Disabled,
		&i.Confirmed,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserReservations = `-- name: GetUserReservations :many
SELECT
    r.uuid,
    r.id,
    r.date,
    r.termin,
    r.washer
FROM
    reservations r
WHERE
    r.uuid = $1
`

type GetUserReservationsRow struct {
	Uuid   pgtype.UUID
	ID     int32
	Date   pgtype.Date
	Termin int32
	Washer int32
}

func (q *Queries) GetUserReservations(ctx context.Context, uuid pgtype.UUID) ([]GetUserReservationsRow, error) {
	rows, err := q.db.Query(ctx, getUserReservations, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserReservationsRow
	for rows.Next() {
		var i GetUserReservationsRow
		if err := rows.Scan(
			&i.Uuid,
			&i.ID,
			&i.Date,
			&i.Termin,
			&i.Washer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsers = `-- name: GetUsers :many
SELECT uuid, phone, name, surname, room, email, disabled, confirmed, role, created_at, updated_at
FROM users
ORDER BY uuid
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Uuid,
			&i.Phone,
			&i.Name,
			&i.Surname,
			&i.Room,
			&i.Email,
			&i.Disabled,
			&i.Confirmed,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchUsers = `-- name: SearchUsers :many
SELECT
    u.uuid,
    u.phone,
    u.name,
    u.surname,
    u.email,
    u.role
FROM
    users u
WHERE
    LOWER(u.name) LIKE LOWER('%' || $1 || '%')
    OR LOWER(u.surname) LIKE LOWER('%' || $1 || '%')
    OR LOWER(u.phone) LIKE LOWER('%' || $1 || '%')
    OR LOWER(u.email) LIKE LOWER('%' || $1 || '%')
`

type SearchUsersRow struct {
	Uuid    pgtype.UUID
	Phone   string
	Name    string
	Surname string
	Email   string
	Role    string
}

func (q *Queries) SearchUsers(ctx context.Context, dollar_1 pgtype.Text) ([]SearchUsersRow, error) {
	rows, err := q.db.Query(ctx, searchUsers, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchUsersRow
	for rows.Next() {
		var i SearchUsersRow
		if err := rows.Scan(
			&i.Uuid,
			&i.Phone,
			&i.Name,
			&i.Surname,
			&i.Email,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}