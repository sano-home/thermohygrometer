CREATE TABLE temperature_and_humidity (id integer primary key, temperature real not null, humidity real not null, unixtimestamp integer not null);
CREATE INDEX idx_unixtimestamp on temperature_and_humidity (unixtimestamp);
