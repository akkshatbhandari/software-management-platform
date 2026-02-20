create table refresh_tokens(
    id serial primary key,
    user_id integer not null,
    token text not null,
    created_at timestamp default current_timestamp,
    expires_at timestamp not null,

    constraint fk_refresh_user
    foreign key(user_id) references users(id) on delete cascade
);