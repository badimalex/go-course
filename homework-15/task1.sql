CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birthdate DATE NOT NULL
);

CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birthdate DATE NOT NULL
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    release_year INT CHECK (release_year >= 1800),
    studio_id INT REFERENCES studios(id),
    box_office FLOAT,
    rating VARCHAR(5) CHECK (rating IN ('PG-10', 'PG-13', 'PG-18')),
    UNIQUE(title, release_year)
);

CREATE TABLE movie_actors (
    movie_id INT REFERENCES movies(id),
    actor_id INT REFERENCES actors(id),
    PRIMARY KEY(movie_id, actor_id)
);

CREATE TABLE movie_directors (
    movie_id INT REFERENCES movies(id),
    director_id INT REFERENCES directors(id),
    PRIMARY KEY(movie_id, director_id)
);
