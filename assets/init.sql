
CREATE TYPE gender AS ENUM  ('male', 'female');
-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id  uuid primary key ,
    gender gender,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    created_time TIMESTAMP NOT NULL ,
    updated_time TIMESTAMP NOT NULL
);

-- таблица актеров
CREATE TABLE IF NOT EXISTS actors (
    id uuid primary key ,
    name text NOT NULL,
    gender gender NOT NULL,
    birthday TIMESTAMP,
    description text,
    created_time TIMESTAMP NOT NULL ,
    updated_time TIMESTAMP NOT NULL
);

--  таблица фильмов
CREATE TABLE IF NOT EXISTS films (
    id uuid primary key,
    name TEXT not null,
    release_date timestamp not null,
    link text,
    description text,
    avg_rating NUMERIC(3,1) CHECK (avg_rating >= 0 AND avg_rating <= 10),
    created_time TIMESTAMP NOT NULL,
    updated_time TIMESTAMP NOT NULL
);

-- таблица с отзывами на фильмы
CREATE TABLE IF NOT EXISTS review (
    id uuid primary key ,
    rating int8,
    description text,
    author_id uuid,
    film_id uuid,
    created_time TIMESTAMP NOT NULL,
    updated_time TIMESTAMP NOT NULL,

    foreign key (author_id) references users(id),
    foreign key (film_id) references films(id)
);

-- таблица для связи фильмы <-> актеры
CREATE TABLE IF NOT EXISTS films_actors (
    id uuid primary key ,
    film_id uuid,
    actor_id uuid,

    foreign key (film_id) references films(id),
    foreign key (actor_id) references actors(id)
);