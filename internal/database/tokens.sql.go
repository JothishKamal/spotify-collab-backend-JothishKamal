// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tokens.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getOAuthToken = `-- name: GetOAuthToken :one
SELECT refresh, access, expiry, spotify_id
FROM tokens
INNER JOIN users on users.user_uuid = tokens.user_uuid
WHERE tokens.user_uuid=$1
`

type GetOAuthTokenRow struct {
	Refresh   []byte             `json:"refresh"`
	Access    []byte             `json:"access"`
	Expiry    pgtype.Timestamptz `json:"expiry"`
	SpotifyID string             `json:"spotify_id"`
}

func (q *Queries) GetOAuthToken(ctx context.Context, userUuid uuid.UUID) (GetOAuthTokenRow, error) {
	row := q.db.QueryRow(ctx, getOAuthToken, userUuid)
	var i GetOAuthTokenRow
	err := row.Scan(
		&i.Refresh,
		&i.Access,
		&i.Expiry,
		&i.SpotifyID,
	)
	return i, err
}

const newToken = `-- name: NewToken :one


INSERT INTO tokens(refresh, access, user_uuid, expiry)
VALUES ($1, $2, $3, $4)
RETURNING user_uuid, refresh, access, expiry
`

type NewTokenParams struct {
	Refresh  []byte             `json:"refresh"`
	Access   []byte             `json:"access"`
	UserUuid uuid.UUID          `json:"user_uuid"`
	Expiry   pgtype.Timestamptz `json:"expiry"`
}

// -- name: CreateToken :exec
// INSERT INTO tokens (hash, user_uuid, expiry, scope)
// VALUES ($1, $2, $3, $4);
// -- name: DeleteTokensForUser :exec
// DELETE FROM tokens
// WHERE scope=$1 AND user_uuid=$2;
func (q *Queries) NewToken(ctx context.Context, arg NewTokenParams) (Token, error) {
	row := q.db.QueryRow(ctx, newToken,
		arg.Refresh,
		arg.Access,
		arg.UserUuid,
		arg.Expiry,
	)
	var i Token
	err := row.Scan(
		&i.UserUuid,
		&i.Refresh,
		&i.Access,
		&i.Expiry,
	)
	return i, err
}

const updateToken = `-- name: UpdateToken :one
UPDATE tokens
SET refresh=$1, access=$2
WHERE user_uuid=$3
RETURNING user_uuid, refresh, access, expiry
`

type UpdateTokenParams struct {
	Refresh  []byte    `json:"refresh"`
	Access   []byte    `json:"access"`
	UserUuid uuid.UUID `json:"user_uuid"`
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) (Token, error) {
	row := q.db.QueryRow(ctx, updateToken, arg.Refresh, arg.Access, arg.UserUuid)
	var i Token
	err := row.Scan(
		&i.UserUuid,
		&i.Refresh,
		&i.Access,
		&i.Expiry,
	)
	return i, err
}
