create table users (
  id serial primary key,
  name text not null,
  created_at timestamp with time zone not null default now()
);