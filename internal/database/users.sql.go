// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(email, password_hash, spotify_id)
VALUES ($1, $2, $3)
RETURNING user_uuid, id, created_at, version
`

type CreateUserParams struct {
	Email        interface{} `json:"email"`
	PasswordHash []byte      `json:"password_hash"`
	SpotifyID    string      `json:"spotify_id"`
}

type CreateUserRow struct {
	UserUuid  uuid.UUID          `json:"user_uuid"`
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Version   int32              `json:"version"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.PasswordHash, arg.SpotifyID)
	var i CreateUserRow
	err := row.Scan(
		&i.UserUuid,
		&i.ID,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_uuid = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userUuid uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, userUuid)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, user_uuid, spotify_id, created_at, updated_at, name, email, password_hash, activated, version
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email interface{}) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.SpotifyID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.Activated,
		&i.Version,
	)
	return i, err
}

const getUserBySpotifyID = `-- name: GetUserBySpotifyID :one
SELECT user_uuid
FROM users 
WHERE spotify_id = $1
`

func (q *Queries) GetUserBySpotifyID(ctx context.Context, spotifyID string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getUserBySpotifyID, spotifyID)
	var user_uuid uuid.UUID
	err := row.Scan(&user_uuid)
	return user_uuid, err
}

const getUserByUUID = `-- name: GetUserByUUID :one
SELECT id, user_uuid, spotify_id, created_at, updated_at, name, email, password_hash, activated, version
FROM users
WHERE user_uuid = $1
`

func (q *Queries) GetUserByUUID(ctx context.Context, userUuid uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUUID, userUuid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.SpotifyID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.Activated,
		&i.Version,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
WHERE id=$5 AND version = $6
RETURNING id, user_uuid, spotify_id, created_at, updated_at, name, email, password_hash, activated, version
`

type UpdateUserParams struct {
	Name         *string     `json:"name"`
	Email        interface{} `json:"email"`
	PasswordHash []byte      `json:"password_hash"`
	Activated    bool        `json:"activated"`
	ID           int64       `json:"id"`
	Version      int32       `json:"version"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.PasswordHash,
		arg.Activated,
		arg.ID,
		arg.Version,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.SpotifyID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.Activated,
		&i.Version,
	)
	return i, err
}
