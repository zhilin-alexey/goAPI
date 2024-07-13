// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: people.sql

package goAPI

import (
	"context"

	"github.com/google/uuid"
)

const create = `-- name: Create :one
insert into people
    (passport_serie, passport_number)
values ($1, $2)
returning id, surname, name, patronymic, address, passport_serie, passport_number
`

type CreateParams struct {
	PassportSerie  string `example:"1234" json:"passportSerie" maxLength:"4" minLength:"4"`
	PassportNumber string `example:"123456" json:"passportNumber" maxLength:"6" minLength:"6"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Person, error) {
	row := q.db.QueryRow(ctx, create, arg.PassportSerie, arg.PassportNumber)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Surname,
		&i.Name,
		&i.Patronymic,
		&i.Address,
		&i.PassportSerie,
		&i.PassportNumber,
	)
	return i, err
}

const delete = `-- name: Delete :exec
delete
from people
where (id = coalesce(cast($1 as uuid), id)
    and surname = coalesce($2, surname)
    and patronymic = coalesce($3, patronymic)
    and address = coalesce($4, address)
    and passport_serie = coalesce(cast($5 as varchar), passport_serie)
    and passport_number = coalesce(cast($6 as varchar), passport_number))
   or false
returning id, surname, name, patronymic, address, passport_serie, passport_number
`

type DeleteParams struct {
	ID             *uuid.UUID `json:"id"`
	Surname        *string    `example:"Иванов" json:"surname"`
	Patronymic     *string    `example:"Иванович" json:"patronymic"`
	Address        *string    `example:"3-й Автозаводский проезд, вл13, Москва, 115280" json:"address"`
	PassportSerie  *string    `json:"passportSerie"`
	PassportNumber *string    `json:"passportNumber"`
}

func (q *Queries) Delete(ctx context.Context, arg DeleteParams) error {
	_, err := q.db.Exec(ctx, delete,
		arg.ID,
		arg.Surname,
		arg.Patronymic,
		arg.Address,
		arg.PassportSerie,
		arg.PassportNumber,
	)
	return err
}

const edit = `-- name: Edit :one
update people
set name            = coalesce($2, name),
    surname         = coalesce($3, surname),
    patronymic      = coalesce($4, patronymic),
    address         = coalesce($5, address),
    passport_serie  = coalesce(cast($6 as varchar), passport_serie),
    passport_number = coalesce(cast($7 as varchar), passport_number)
where id = $1
returning id, surname, name, patronymic, address, passport_serie, passport_number
`

type EditParams struct {
	ID             uuid.UUID `example:"00000000-0000-0000-0000-000000000000" format:"uuid" json:"id"`
	Name           *string   `example:"Иван" json:"name"`
	Surname        *string   `example:"Иванов" json:"surname"`
	Patronymic     *string   `example:"Иванович" json:"patronymic"`
	Address        *string   `example:"3-й Автозаводский проезд, вл13, Москва, 115280" json:"address"`
	PassportSerie  *string   `json:"passportSerie"`
	PassportNumber *string   `json:"passportNumber"`
}

func (q *Queries) Edit(ctx context.Context, arg EditParams) (Person, error) {
	row := q.db.QueryRow(ctx, edit,
		arg.ID,
		arg.Name,
		arg.Surname,
		arg.Patronymic,
		arg.Address,
		arg.PassportSerie,
		arg.PassportNumber,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Surname,
		&i.Name,
		&i.Patronymic,
		&i.Address,
		&i.PassportSerie,
		&i.PassportNumber,
	)
	return i, err
}

const getByPassport = `-- name: GetByPassport :one
select surname, name, patronymic, address
from people
where passport_serie = $1
  and passport_number = $2
limit 1
`

type GetByPassportParams struct {
	PassportSerie  string `example:"1234" json:"passportSerie" maxLength:"4" minLength:"4"`
	PassportNumber string `example:"123456" json:"passportNumber" maxLength:"6" minLength:"6"`
}

type GetByPassportRow struct {
	Surname    *string `example:"Иванов" json:"surname"`
	Name       *string `example:"Иван" json:"name"`
	Patronymic *string `example:"Иванович" json:"patronymic"`
	Address    *string `example:"3-й Автозаводский проезд, вл13, Москва, 115280" json:"address"`
}

func (q *Queries) GetByPassport(ctx context.Context, arg GetByPassportParams) (GetByPassportRow, error) {
	row := q.db.QueryRow(ctx, getByPassport, arg.PassportSerie, arg.PassportNumber)
	var i GetByPassportRow
	err := row.Scan(
		&i.Surname,
		&i.Name,
		&i.Patronymic,
		&i.Address,
	)
	return i, err
}

const getMultiple = `-- name: GetMultiple :many
select id, surname, name, patronymic, address, passport_serie, passport_number
from people
where id = coalesce(cast($5 as uuid), id)
  and name = coalesce($1, name)
  and surname = coalesce($2, surname)
  and patronymic = coalesce($3, patronymic)
  and address = coalesce($4, address)
  and passport_serie = coalesce(cast($6 as varchar), passport_serie)
  and passport_number = coalesce(cast($7 as varchar), passport_number)
offset coalesce($8, 0::bigint) limit cast($9 as bigint)
`

type GetMultipleParams struct {
	Name           *string    `example:"Иван" json:"name"`
	Surname        *string    `example:"Иванов" json:"surname"`
	Patronymic     *string    `example:"Иванович" json:"patronymic"`
	Address        *string    `example:"3-й Автозаводский проезд, вл13, Москва, 115280" json:"address"`
	ID             *uuid.UUID `json:"id"`
	PassportSerie  *string    `json:"passportSerie"`
	PassportNumber *string    `json:"passportNumber"`
	Offset         *int64     `json:"offset"`
	Limit          *int64     `json:"limit"`
}

func (q *Queries) GetMultiple(ctx context.Context, arg GetMultipleParams) ([]Person, error) {
	rows, err := q.db.Query(ctx, getMultiple,
		arg.Name,
		arg.Surname,
		arg.Patronymic,
		arg.Address,
		arg.ID,
		arg.PassportSerie,
		arg.PassportNumber,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.Surname,
			&i.Name,
			&i.Patronymic,
			&i.Address,
			&i.PassportSerie,
			&i.PassportNumber,
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
