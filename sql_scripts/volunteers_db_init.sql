drop database if exists volunteers;
create database volunteers;
drop table if exists volunteers cascade;
drop table if exists teams cascade;
drop table if exists team_members cascade;
create table volunteers
(
  id       serial PRIMARY KEY,
  email    varchar(128),
  name     varchar(64),
  password varchar(64)
);
create table teams
(
  id          serial PRIMARY KEY,
  name        varchar(32),
  createdById serial,
  team_motto  varchar(256)
);
create table team_members
(
  id          serial PRIMARY KEY,
  teamId      serial references teams (id) ON DELETE CASCADE,
  volunteerId serial references volunteers (id) ON DELETE CASCADE
);