create table if not exists student(
    id serial primary key,
    first_name varchar,
    last_name varchar,
    username varchar,
    email varchar,
    phone_number varchar,
    created_at timestamp default current_timestamp
);