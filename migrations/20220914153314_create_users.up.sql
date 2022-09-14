create table if not exists users(
  id serial not null unique primary key,
  email varchar(255) not null unique,
  password varchar(60) not null,
  created_at timestamp without time zone default now(),
  updated_at timestamp without time zone default now()
);

