create table if not exists tasks
(
    id         uuid primary key                                       default gen_random_uuid(),
    people_id  uuid references people (id) on delete cascade not null,
    name       varchar                                       not null,
    start_time timestamptz                                   not null default now(),
    end_time   timestamptz                                            default null
);

create or replace function manage_timestamps()
    returns trigger as $$
begin
    update tasks
    set end_time = now()
    where people_id = new.people_id and end_time is null;

    return new;
end;
$$ language plpgsql;

create trigger check_and_update_timestamps
    before insert on tasks
    for each row
execute function manage_timestamps();

