create table if not exists songs (
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    group_id uuid not null,
    release_date DATE not null,
    lyrics text not null,
    link varchar(255) not null
    FOREIGN KEY (group_id) REFERENCES "groups"(id)
);

create table if not exists "groups" (
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
);
