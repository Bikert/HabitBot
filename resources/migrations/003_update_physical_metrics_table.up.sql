DROP TABLE IF EXISTS physical_metrics;

CREATE TABLE body_metrics
(
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id         INTEGER NOT NULL,
    date            DATE    NOT NULL,
    weight          REAL,
    biceps_left     REAL,
    biceps_right    REAL,
    chest           REAL,
    waist           REAL,
    belly           REAL,
    hips            REAL,
    thigh_max_left  REAL,
    thigh_max_right REAL,
    thigh_low_left  REAL,
    thigh_low_right REAL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (user_id, date)
);