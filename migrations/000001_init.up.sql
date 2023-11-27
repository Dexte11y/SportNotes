CREATE TABLE users (
  id INT PRIMARY KEY,
  login VARCHAR(255),
  name VARCHAR(255),
  surname VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255)
);

CREATE TABLE workouts (
  id INT PRIMARY KEY,
  id_user INT,
  created_at DATE
);

CREATE TABLE trainings (
  id INT PRIMARY KEY,
  id_workout INT,
  type VARCHAR(255),
  name VARCHAR(255),
  approaches INT,
  repetitions INT,
  weight INT
);