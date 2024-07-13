-- name: InsertPeople :copyfrom
insert into people
    (surname, name, patronymic, address, passport_serie, passport_number)
values ($1, $2, $3, $4, $5, $6);

-- name: InsertTasks :copyfrom
insert into tasks
    (people_id, name, start_time, end_time)
values ($1, $2, $3, $4);

-- name: RandomPeopleId :one
select id
from people
order by random()
limit 1;

-- name: CheckSeedExecuted :one
select exists(select 1
              from seeds
              where name = $1);

-- name: MarkSeedRan :exec
insert into seeds
values ($1);
