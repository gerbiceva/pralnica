-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE id = $1;

-- name: ConfirmUser :exec
UPDATE users
SET confirmed = TRUE
WHERE uuid = $1;

-- name: BanUser :exec
UPDATE users
SET disabled = TRUE
WHERE uuid = $1;

-- name: EditUser :exec
UPDATE users
SET phone = $2,
    name = $3,
    surname = $4,
    room = $5,
    email = $6,
    role = $7
WHERE uuid = $1;

-- name: GetUser :one
SELECT uuid, phone, name, surname, room, email, disabled, confirmed, role, created_at, updated_at
FROM users
WHERE uuid = $1;

-- name: GetUserReservations :many
SELECT * FROM get_user_reservations($1);

-- name: GetReservationsByMonth :many
SELECT * FROM get_reservations_by_month($1);

-- name: SearchUsers :many
SELECT * FROM search_users($1);