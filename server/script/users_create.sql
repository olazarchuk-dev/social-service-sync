create table users
(
    id serial not null constraint user_pkey primary key,
    device_name text,
    password text,
    email text,
    image text
);
