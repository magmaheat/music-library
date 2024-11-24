create table if not exists "groups" (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null
);

create table if not exists songs (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null,
    group_id uuid not null
);
