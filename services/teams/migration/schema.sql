create table teams
(
  id serial primary key,
  name_teams varchar(50) not null,
  user_id varchar(50) not null references users(id),
  descripcion varchar(255) not null,
  rekrutment boolean not null,
  cari_lawan boolean default false,
  kejuaraan integer default 0,
  created_at timestamp default now(),
  updated_at timestamp default now()
);

create table team_members (
  id serial primary key,
  user_id varchar(50) not null references users(id),
  team_id varchar(50) not null references teams(id),
  created_at timestamp default now(),
  updated_at timestamp default now()
);
