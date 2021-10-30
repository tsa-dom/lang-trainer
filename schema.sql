CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR ( 30 ) UNIQUE,
  passwordHash TEXT,
  priviledges VARCHAR ( 20 )
);