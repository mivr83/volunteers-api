create database volunteers_db;
create table volunteers
(
  id       serial PRIMARY KEY,
  email    varchar(128) unique,
  name     varchar(64),
  password varchar(64)
);
create table teams
(
  id            serial PRIMARY KEY,
  name          varchar(32) unique,
  created_by_id serial,
  team_motto    varchar(256)
);
create table team_members
(
  id           serial PRIMARY KEY,
  team_id      serial references teams (id) ON DELETE CASCADE,
  volunteer_id serial references volunteers (id) ON DELETE CASCADE,
  CONSTRAINT unq_team_id_volunteer_id UNIQUE (team_id, volunteer_id)
);