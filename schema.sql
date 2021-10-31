CREATE TABLE IF NOT EXISTS Users (
  id SERIAL PRIMARY KEY,
  username VARCHAR ( 30 ) UNIQUE NOT NULL,
  passwordHash TEXT NOT NULL,
  priviledges VARCHAR ( 30 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS WordItems (
  id SERIAL PRIMARY KEY,
  word VARCHAR ( 30 ) NOT NULL,
  description TEXT,
  CHECK (word <> '')
);

CREATE TABLE IF NOT EXISTS Words (
  id SERIAL PRIMARY KEY,
  ownerId INTEGER NOT NULL,
  wordItemId INTEGER NOT NULL,
  targetItemId INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS WordGroupLinks (
  id SERIAL PRIMARY KEY,
  wordGroupId INTEGER NOT NULL,
  wordId INTEGER NOT NULL,
  UNIQUE (wordGroupId, wordId)
);

CREATE TABLE IF NOT EXISTS WordGroups (
  id SERIAL PRIMARY KEY,
  ownerId INTEGER NOT NULL,
  name VARCHAR ( 30 ) NOT NULL,
  description TEXT,
  CHECK (name <> '')
);