CREATE TABLE measurements (
    id SERIAL PRIMARY KEY,
    temperature DOUBLE PRECISION NOT NULL,
    humidity DOUBLE PRECISION NOT NULL,
    timestamp BIGINT NOT NULL
);