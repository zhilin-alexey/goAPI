-- name: GetByPassport :one
select surname, name, patronymic, address
from people
where passport_serie = $1
  and passport_number = $2
limit 1;

-- name: GetMultiple :many
select *
from people
where id = coalesce(cast(sqlc.narg('id') as uuid), id)
  and name = coalesce($1, name)
  and surname = coalesce($2, surname)
  and patronymic = coalesce($3, patronymic)
  and address = coalesce($4, address)
  and passport_serie = coalesce(cast(sqlc.narg('passport_serie') as varchar), passport_serie)
  and passport_number = coalesce(cast(sqlc.narg('passport_number') as varchar), passport_number)
offset coalesce(sqlc.narg('offset'), 0::bigint) limit cast(sqlc.narg('limit') as bigint);

-- name: Delete :exec
delete
from people
where (id = coalesce(cast(sqlc.narg('id') as uuid), id)
    and surname = coalesce($2, surname)
    and patronymic = coalesce($3, patronymic)
    and address = coalesce($4, address)
    and passport_serie = coalesce(cast(sqlc.narg('passport_serie') as varchar), passport_serie)
    and passport_number = coalesce(cast(sqlc.narg('passport_number') as varchar), passport_number))
   or false
returning *;

-- name: Create :one
insert into people
    (passport_serie, passport_number)
values ($1, $2)
returning *;

-- name: Edit :one
update people
set name            = coalesce($2, name),
    surname         = coalesce($3, surname),
    patronymic      = coalesce($4, patronymic),
    address         = coalesce($5, address),
    passport_serie  = coalesce(cast(sqlc.narg('passport_serie') as varchar), passport_serie),
    passport_number = coalesce(cast(sqlc.narg('passport_number') as varchar), passport_number)
where id = $1
returning *;
