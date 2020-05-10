CREATE SCHEMA IF NOT EXISTS helpee;

CREATE TABLE IF NOT EXISTS helpee.users (
    id SERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    date_of_birth BIGINT,
    address VARCHAR,
    email VARCHAR NOT NULL UNIQUE,
    username VARCHAR NOT NULL UNIQUE,
    password VARCHAR
);


INSERT INTO helpee.users (id, first_name, last_name, date_of_birth, address, email, username, password) VALUES
(1, 'Mladen', 'Mladjenovic', '522547200', 'Cerevicka 42/2', 'mladen@example.org', 'mladen', 'mladen1'),
(2, 'Milutin', 'Mladjenovic', '1430524800', 'Cerevicka 42/2','milutin@example.org', 'milutin', 'milutin1'),
(3, 'Momcilo', 'Mladjenovic', '1482451200', 'Cerevicka 42/2', 'momcilo@example.org', 'momcilo', 'momcilo1');