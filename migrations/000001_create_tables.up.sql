create table if not exists "groups" (
    id uuid primary key,
    name varchar(255) not null,
    group_id uuid not null,
    release_date DATE not null,
    lyrics text not null,
    link varchar(255) not null
);

create table if not exists songs (
    id uuid primary key,
    name varchar(255) not null,
    group_id uuid not null
);
