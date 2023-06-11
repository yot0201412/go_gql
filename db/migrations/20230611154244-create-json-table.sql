
-- +migrate Up
create table json_table (
  id serial primary key,
  json_data jsonb not null
);

-- +migrate Down
drop table json_table;

