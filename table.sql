CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "description" VARCHAR(255),
)

CREATE TABLE time_logs (
    id SERIAL PRIMARY KEY,
    "time" INTERVAL NOT NULL,
    "date" TIMESTAMP NOT NULL
);