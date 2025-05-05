BEGIN;
CREATE TABLE IF NOT EXISTS persons (
                                       id SERIAL PRIMARY KEY,
                                       name TEXT NOT NULL,
                                       surname TEXT NOT NULL,
                                       patronymic TEXT,
                                       gender TEXT,
                                       age INT,
                                       nationality TEXT,
                                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       updated_at TIMESTAMP
);
COMMIT;