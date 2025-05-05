BEGIN;
INSERT INTO persons (name, surname, patronymic, gender, age, nationality, updated_at)
VALUES
    ('Иван', 'Иванов', 'Иванович', 'male', 30, 'RU', NOW()),
    ('Мария', 'Петрова', 'Сергеевна', 'female', 25, 'RU', NOW()),
    ('John', 'Doe', NULL, 'male', 40, 'US', NOW()),
    ('Анна', 'Смирнова', NULL, 'female', 35, 'RU', NOW()),
    ('Hans', 'Schmidt', NULL, 'male', 45, 'DE', NOW());

COMMIT;