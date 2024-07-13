-- name: StartTask :one
insert into tasks
    (people_id, name)
values ($1, $2)
returning *;

-- name: EndTask :one
update tasks
set end_time = now()
where people_id = $1
  and end_time is null
returning *;

-- name: GetTasksByPeople :many
select name,
       (extract(epoch from coalesce(end_time, now()) - start_time) / 3600)::integer   as hours,
       (extract(epoch from coalesce(end_time, now()) - start_time) / 60 % 60)::integer as minutes
from tasks
where people_id = $1
  and (
    (sqlc.narg('period_start')::timestamptz is null and sqlc.narg('period_end')::timestamptz is null)
        or tstzrange(
                   coalesce(cast(sqlc.narg('period_start') as timestamptz), '-infinity'::timestamptz),
                   coalesce(cast(sqlc.narg('period_end') as timestamptz), 'infinity'::timestamptz), '[]'
           ) @> tstzrange(start_time, coalesce(end_time, 'infinity'::timestamptz), '[]')
    )
order by (coalesce(end_time, now()) - start_time) desc;
