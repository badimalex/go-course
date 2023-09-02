-- Добавляем студии
INSERT INTO studios (name) VALUES ('Warner Bros'), ('Marvel Studios'), ('Pixar'), ('Universal Pictures');

-- Добавляем актёров
INSERT INTO actors (name, birthdate) VALUES ('Robert Downey Jr.', '1965-04-04'), ('Scarlett Johansson', '1984-11-22'), ('Chris Hemsworth', '1983-08-11'), ('Emma Watson', '1990-04-15');

-- Добавляем режиссёров
INSERT INTO directors (name, birthdate) VALUES ('Christopher Nolan', '1970-07-30'), ('Joss Whedon', '1964-06-23'), ('Steven Spielberg', '1946-12-18'), ('Greta Gerwig', '1983-08-04');

-- Добавляем фильмы
INSERT INTO movies (title, release_year, studio_id, box_office, rating)
VALUES
('Inception', 2010, 1, 829.9, 'PG-13'),
('Avengers: Endgame', 2019, 2, 2797.8, 'PG-13'),
('Jurassic Park', 1993, 4, 912.7, 'PG-13'),
('Little Women', 2019, 4, 206.6, 'PG-13'),
('Toy Story', 1995, 3, 373.6, 'PG-10');

-- Связываем фильмы с актёрами
INSERT INTO movie_actors (movie_id, actor_id)
VALUES
(1, 1),
(2, 1),
(2, 2),
(2, 3),
(4, 4);

-- Связываем фильмы с режиссёрами
INSERT INTO movie_directors (movie_id, director_id)
VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 3);
