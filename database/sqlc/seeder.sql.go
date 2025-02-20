// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: seeder.sql

package goAPI

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const checkSeedExecuted = `-- name: CheckSeedExecuted :one
select exists(select 1
              from seeds
              where name = $1)
`

func (q *Queries) CheckSeedExecuted(ctx context.Context, name *string) (bool, error) {
	row := q.db.QueryRow(ctx, checkSeedExecuted, name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

type InsertPeopleParams struct {
	Surname        *string `example:"Иванов" json:"surname"`
	Name           *string `example:"Иван" json:"name"`
	Patronymic     *string `example:"Иванович" json:"patronymic"`
	Address        *string `example:"3-й Автозаводский проезд, вл13, Москва, 115280" json:"address"`
	PassportSerie  string  `example:"1234" json:"passportSerie" maxLength:"4" minLength:"4"`
	PassportNumber string  `example:"123456" json:"passportNumber" maxLength:"6" minLength:"6"`
}

type InsertTasksParams struct {
	PeopleID  uuid.UUID  `json:"peopleId"`
	Name      string     `example:"Помыть посуду" json:"name"`
	StartTime time.Time  `example:"2022-01-01T00:00:00Z" format:"dateTime" json:"startTime"`
	EndTime   *time.Time `example:"2022-01-01T00:00:00Z" format:"dateTime" json:"endTime"`
}

const markSeedRan = `-- name: MarkSeedRan :exec
insert into seeds
values ($1)
`

func (q *Queries) MarkSeedRan(ctx context.Context, dollar_1 *string) error {
	_, err := q.db.Exec(ctx, markSeedRan, dollar_1)
	return err
}

const randomPeopleId = `-- name: RandomPeopleId :one
select id
from people
order by random()
limit 1
`

func (q *Queries) RandomPeopleId(ctx context.Context) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, randomPeopleId)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
