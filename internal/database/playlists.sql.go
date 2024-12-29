// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: playlists.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createPlaylist = `-- name: CreatePlaylist :one
INSERT INTO playlists (playlist_id, user_uuid, name, playlist_code)
VALUES ($1, $2, $3, $4)
RETURNING user_uuid, playlist_uuid, playlist_id, name, playlist_code, created_at, updated_at
`

type CreatePlaylistParams struct {
	PlaylistID   string    `json:"playlist_id"`
	UserUuid     uuid.UUID `json:"user_uuid"`
	Name         string    `json:"name"`
	PlaylistCode string    `json:"playlist_code"`
}

func (q *Queries) CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (Playlist, error) {
	row := q.db.QueryRow(ctx, createPlaylist,
		arg.PlaylistID,
		arg.UserUuid,
		arg.Name,
		arg.PlaylistCode,
	)
	var i Playlist
	err := row.Scan(
		&i.UserUuid,
		&i.PlaylistUuid,
		&i.PlaylistID,
		&i.Name,
		&i.PlaylistCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePlaylist = `-- name: DeletePlaylist :execrows
DELETE FROM playlists
WHERE playlist_uuid = $1
`

func (q *Queries) DeletePlaylist(ctx context.Context, playlistUuid uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deletePlaylist, playlistUuid)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const getPlaylist = `-- name: GetPlaylist :one
SELECT user_uuid, playlist_uuid, playlist_id, name, playlist_code, created_at, updated_at
FROM playlists
WHERE playlist_uuid = $1
`

func (q *Queries) GetPlaylist(ctx context.Context, playlistUuid uuid.UUID) (Playlist, error) {
	row := q.db.QueryRow(ctx, getPlaylist, playlistUuid)
	var i Playlist
	err := row.Scan(
		&i.UserUuid,
		&i.PlaylistUuid,
		&i.PlaylistID,
		&i.Name,
		&i.PlaylistCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlaylistIDByUUID = `-- name: GetPlaylistIDByUUID :one
SELECT playlist_id
FROM playlists
WHERE playlist_uuid = $1
`

func (q *Queries) GetPlaylistIDByUUID(ctx context.Context, playlistUuid uuid.UUID) (string, error) {
	row := q.db.QueryRow(ctx, getPlaylistIDByUUID, playlistUuid)
	var playlist_id string
	err := row.Scan(&playlist_id)
	return playlist_id, err
}

const getPlaylistUUIDByCode = `-- name: GetPlaylistUUIDByCode :one
SELECT playlist_uuid
FROM playlists
WHERE playlist_code = $1
`

func (q *Queries) GetPlaylistUUIDByCode(ctx context.Context, playlistCode string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getPlaylistUUIDByCode, playlistCode)
	var playlist_uuid uuid.UUID
	err := row.Scan(&playlist_uuid)
	return playlist_uuid, err
}

const getPlaylistUUIDByName = `-- name: GetPlaylistUUIDByName :one
SELECT playlist_uuid
FROM playlists
WHERE user_uuid = $1 AND name = $2
`

type GetPlaylistUUIDByNameParams struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	Name     string    `json:"name"`
}

func (q *Queries) GetPlaylistUUIDByName(ctx context.Context, arg GetPlaylistUUIDByNameParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getPlaylistUUIDByName, arg.UserUuid, arg.Name)
	var playlist_uuid uuid.UUID
	err := row.Scan(&playlist_uuid)
	return playlist_uuid, err
}

const listPlaylists = `-- name: ListPlaylists :many
SELECT user_uuid, playlist_uuid, playlist_id, name, playlist_code, created_at, updated_at
FROM playlists
WHERE user_uuid = $1
`

func (q *Queries) ListPlaylists(ctx context.Context, userUuid uuid.UUID) ([]Playlist, error) {
	rows, err := q.db.Query(ctx, listPlaylists, userUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Playlist
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.UserUuid,
			&i.PlaylistUuid,
			&i.PlaylistID,
			&i.Name,
			&i.PlaylistCode,
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

const updatePlaylistName = `-- name: UpdatePlaylistName :one
UPDATE playlists
SET name = $1
WHERE playlist_uuid = $2
RETURNING user_uuid, playlist_uuid, playlist_id, name, playlist_code, created_at, updated_at
`

type UpdatePlaylistNameParams struct {
	Name         string    `json:"name"`
	PlaylistUuid uuid.UUID `json:"playlist_uuid"`
}

func (q *Queries) UpdatePlaylistName(ctx context.Context, arg UpdatePlaylistNameParams) (Playlist, error) {
	row := q.db.QueryRow(ctx, updatePlaylistName, arg.Name, arg.PlaylistUuid)
	var i Playlist
	err := row.Scan(
		&i.UserUuid,
		&i.PlaylistUuid,
		&i.PlaylistID,
		&i.Name,
		&i.PlaylistCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
