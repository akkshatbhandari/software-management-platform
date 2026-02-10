alter table projects 
add column user_id integer not null;

alter table projects
add constraint fk_projects_user
foreign key (user_id) references users(id)
on delete cascade;
