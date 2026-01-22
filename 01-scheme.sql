-- Номер завдання 2, пункт 1, створення схеми

CREATE TYPE sensor_type AS ENUM ('V', 'R');

CREATE TABLE rooms (
                       id      BIGSERIAL PRIMARY KEY,
                       name    TEXT NOT NULL UNIQUE
);

CREATE TABLE sensors (
                         id        BIGSERIAL PRIMARY KEY,
                         room_id  BIGINT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE
                         name      TEXT NOT NULL UNIQUE,
                         type      sensor_type NOT NULL,
);

CREATE INDEX idx_sensors_room ON sensors(room_id);
CREATE INDEX idx_sensors_type ON sensors(type);

CREATE TABLE measurements (
                              id         BIGSERIAL PRIMARY KEY,
                              sensor_id  BIGINT NOT NULL REFERENCES sensors(id) ON DELETE CASCADE,
                              value      DOUBLE PRECISION NOT NULL,
                              ts         TIMESTAMP(6) NOT NULL
);

CREATE INDEX idx_measurements_sensor_ts
    ON measurements(sensor_id, ts);