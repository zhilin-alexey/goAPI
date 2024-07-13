create table if not exists people
(
    id              UUID primary key default gen_random_uuid(),
    surname         varchar(20),
    name            varchar(20),
    patronymic      varchar(20),
    address         varchar(200),
    passport_serie  varchar(4) not null,
    passport_number varchar(6) not null
);

create unique index if not exists unique_passport
    on people (passport_serie, passport_number);