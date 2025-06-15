-- Создаём таблицу sessions
CREATE TABLE IF NOT EXISTS sessions
(
    user_id       INTEGER PRIMARY KEY,
    next_step     TEXT,
    previous_step TEXT,
    data          TEXT
);

-- Создаём таблицу users
CREATE TABLE IF NOT EXISTS users
(
    id         INTEGER PRIMARY KEY,
    username   TEXT,
    first_name TEXT,
    last_name  TEXT,
    created_at DATETIME
);

-- Создаём таблицу habits
CREATE TABLE IF NOT EXISTS habits
(
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    title     TEXT,
    isDefault INTEGER DEFAULT 0
);

-- Создаём таблицу user_habits с внешними ключами
CREATE TABLE IF NOT EXISTS user_habits
(
    user_id  INTEGER,
    habit_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (habit_id) REFERENCES habits (id)
);