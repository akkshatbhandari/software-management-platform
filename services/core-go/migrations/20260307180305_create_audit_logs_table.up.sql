create table audit_logs(
    id serial primary key,
    user_id integer,
    action text not null,
    resource text not null,
    resource_id integer,
    metadata jsonb,
    created_at timestamp default current_timestamp
);