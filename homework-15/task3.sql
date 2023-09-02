-- 1. выборка фильмов с названием студии;
SELECT m.title, s.name AS studio_name
FROM movies m
JOIN studios s ON m.studio_id = s.id;

-- 2. выборка фильмов для некоторого актёра;
SELECT m.title
FROM movies m
JOIN movie_actors ma ON m.id = ma.movie_id
JOIN actors a ON ma.actor_id = a.id
WHERE a.name = 'Emma Watson';

-- 3. подсчёт фильмов для некоторого режиссёра;
SELECT COUNT(m.id)
FROM movies m
JOIN movie_directors md ON m.id = md.movie_id
JOIN directors d ON md.director_id = d.id
WHERE d.name = 'Steven Spielberg';

-- 4. выборка фильмов для нескольких режиссёров из списка (подзапрос);
SELECT m.title
FROM movies m
WHERE m.id IN (
    SELECT md.movie_id
    FROM movie_directors md
    WHERE md.director_id IN (
        SELECT id FROM directors WHERE name IN ('Steven Spielberg', 'Martin Scorsese')
    )
);

-- 5. подсчёт количества фильмов для актёра;
SELECT COUNT(m.id)
FROM movies m
JOIN movie_actors ma ON m.id = ma.movie_id
JOIN actors a ON ma.actor_id = a.id
WHERE a.name = 'Robert Downey Jr.';

-- 6. выборка актёров и режиссёров, участвовавших более чем в 2 фильмах;
-- Актеры:
SELECT a.name
FROM actors a
JOIN movie_actors ma ON a.id = ma.actor_id
GROUP BY a.name
HAVING COUNT(ma.movie_id) > 2;

-- Режиссеры:
SELECT d.name
FROM directors d
JOIN movie_directors md ON d.id = md.director_id
GROUP BY d.name
HAVING COUNT(md.movie_id) > 2;

-- 7. подсчёт количества фильмов со сборами больше 1000;
SELECT COUNT(id)
FROM movies
WHERE box_office > 1000;

-- 8. подсчитать количество режиссёров, фильмы которых собрали больше 1000;
SELECT COUNT(DISTINCT md.director_id)
FROM movies m
JOIN movie_directors md ON m.id = md.movie_id
WHERE m.box_office > 1000;

-- 9. выборка различных фамилий актёров;
SELECT DISTINCT SPLIT_PART(name, ' ', 2) AS last_name
FROM Actor;

-- 9. выборка различных фамилий актёров;
SELECT DISTINCT SPLIT_PART(name, ' ', 2) AS last_name
FROM actors;

-- 10. подсчёт количества фильмов, имеющих дубли по названию.
SELECT COUNT(title) FROM movies GROUP BY title HAVING COUNT(title) > 1;
