DROP TABLE IF EXISTS body_metrics ;

CREATE TABLE physical_metrics
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    date    DATE    NOT NULL,
    height  REAL,
    weight  REAL,
    biceps  REAL,
    chest   REAL,
    waist   REAL,
    belly   REAL,
    hips    REAL,
    thigh1  REAL,
    thigh2  REAL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);