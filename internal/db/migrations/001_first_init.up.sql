CREATE TABLE sessions
(
    user_id       INTEGER PRIMARY KEY,
    next_step     TEXT,
    previous_step TEXT,
    data          TEXT,
    scenario      TEXT
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
    version_id   INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id      INTEGER NOT NULL,
    group_id     TEXT NOT NULL, -- UUID в формате TEXT
    version      INTEGER NOT NULL,
    name         TEXT    NOT NULL,
    description  TEXT,
    color        TEXT,
    icon         TEXT,
    repeat_type  TEXT    NOT NULL DEFAULT 'daily',
    days_of_week TEXT,
    isDefault    INTEGER          DEFAULT 0,
    is_active    BOOLEAN NOT NULL DEFAULT 1,
    first_date   DATETIME,
    last_date    DATETIME,
    created_at   DATETIME         DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE habit_completions
(
    habit_version_id INTEGER NOT NULL,
    date             DATE    NOT NULL,
    completed        BOOLEAN NOT NULL DEFAULT 1,
    PRIMARY KEY (habit_version_id, date),
    FOREIGN KEY (habit_version_id) REFERENCES habits (version_id) ON DELETE CASCADE
);
