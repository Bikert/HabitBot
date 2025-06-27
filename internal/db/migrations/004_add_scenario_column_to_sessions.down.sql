PRAGMA foreign_keys=off;

CREATE TABLE sessions_new (
                              user_id       INTEGER PRIMARY KEY,
                              next_step     TEXT,
                              previous_step TEXT,
                              data          TEXT
);

INSERT INTO sessions_new (user_id, next_step, previous_step, data)
SELECT user_id, next_step, previous_step, data FROM sessions;

DROP TABLE sessions;

ALTER TABLE sessions_new RENAME TO sessions;

PRAGMA foreign_keys=on;