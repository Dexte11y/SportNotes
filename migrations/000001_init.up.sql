CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  login VARCHAR(20),
  name VARCHAR(20),
  surname VARCHAR(20),
  email VARCHAR(20),
  password VARCHAR(20)
);

CREATE TABLE workouts (
  id SERIAL PRIMARY KEY,
  id_user INT references users(id) on delete cascade not null,
  type VARCHAR(50),
  created_at DATE
);

CREATE TABLE trainings (
  id SERIAL PRIMARY KEY,
  id_workout INT references workouts(id) on delete cascade not null,
  name VARCHAR(50),
  approaches INT,
  repetitions INT,
  weight INT
);

CREATE TABLE nutritions (
  id SERIAL PRIMARY KEY,
  id_user INT references users(id) on delete cascade not null,
  type VARCHAR(50),
  created_at DATE
);

CREATE TABLE foods (
  id SERIAL PRIMARY KEY,
  id_nutrition INT references nutritions(id) on delete cascade not null,
  name VARCHAR(50)
);