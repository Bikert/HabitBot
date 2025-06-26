CREATE TABLE sessions
(
    user_id       INTEGER PRIMARY KEY,
    next_step     TEXT,
    previous_step TEXT,
    data          TEXT
);

CREATE TABLE users
(
    id         INTEGER PRIMARY KEY,
    username   TEXT,
    first_name TEXT,
    last_name  TEXT,
    created_at DATETIME
);

CREATE TABLE habits
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id      INTEGER NOT NULL,
    group_id     TEXT    NOT NULL, -- UUID в формате TEXT
    version      INTEGER NOT NULL,
    name         TEXT    NOT NULL,
    description  TEXT,
    color        TEXT,
    icon         TEXT,
    is_active    BOOLEAN NOT NULL DEFAULT 1,
    repeat_type  TEXT    NOT NULL DEFAULT 'daily',
    days_of_week TEXT,
    isDefault    INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE habit_completions
(
    habit_id INTEGER NOT NULL,
    date     DATE    NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT 1,
    PRIMARY KEY (habit_id, date),
    FOREIGN KEY (habit_id) REFERENCES habits (id) ON DELETE CASCADE
);