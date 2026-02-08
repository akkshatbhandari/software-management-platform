create table users(
    id serial primary key,
    email text not null unique,
    password_hash text not null,
    created_at timestamp default current_timestamp
);