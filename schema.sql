CREATE TABLE IF NOT EXISTS Users (
  id SERIAL PRIMARY KEY,
  username VARCHAR ( 30 ) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  priviledges VARCHAR ( 30 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS Groups (
  id SERIAL PRIMARY KEY,
  owner_id INTEGER REFERENCES Users,
  name VARCHAR ( 30 ) NOT NULL,
  description TEXT,
  CHECK (name <> '')
);

CREATE TABLE IF NOT EXISTS Words (
  id SERIAL PRIMARY KEY,
  owner_id INTEGER REFERENCES Users,
  word VARCHAR ( 30 ) NOT NULL,
  description TEXT,
  CHECK (word <> '')
);

CREATE TABLE IF NOT EXISTS WordItems (
  id SERIAL PRIMARY KEY,
  word_id INTEGER REFERENCES Words,
  word VARCHAR ( 30 ) NOT NULL,
  description TEXT,
  CHECK (word <> '')
);

CREATE TABLE IF NOT EXISTS GroupLinks (
  group_id INTEGER REFERENCES Groups,
  word_id INTEGER REFERENCES Words,
  UNIQUE (group_id, word_id)
);

