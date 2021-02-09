# thermohygrometer

Read temperature and humidity from DHT11 sensor and save its history to SQLite3 database on Raspberry Pi. The thermohygrometer is three main things below.

- Collector: Read data from DHT11 via GPIO and insert it into SQLite3 database.
- Server: API Server to retrieve temperature and humidity history.
- SQLite3: As a database.

# Installation

Clone and build on Raspberry Pi that you wanna run it on.

## SQLite3 database

Prepare SQLite3 database before running collector and server. Run `sqlite3` command then create a table and an index to initialize database.

```shell
sqlite3 /path/to/sensor.db
```

```sql
CREATE TABLE temperature_and_humidity (id integer primary key, temperature real not null, humidity real not null, unixtimestamp integer not null);
CREATE INDEX idx_unixtimestamp on temperature_and_humidity (unixtimestamp);
```

## Collector

NOTE: Run as root to access GPIO.

```shell
cd thermohygrometer/cmd/collector
go build
sudo ./collector -gpio 14 -db /path/to/sensor.db -interval 10s
```

- -gpio: The GPIO input-pin number.
- -db: Path to the SQLite3 database.
- -interval: Data collection intervals.

## Server

```shell
cd thermohygrometer/cmd/server
go build
server -host 0.0.0.0 -port 8000 -db /path/to/sensor.db
```

- -host: Listen IP address.
- -port: Listen port.
- -db: Path to the SQLite3 database.

# Server API

## GET /current

Get the latest temperature and humidity from database.

### Example

```
{
    "temperature":22,
    "humidity":15,
    "timestamp":"2021-01-10T13:25:48Z"
}
```

## GET /histories

Get temperature and humidity history in the range of given parameters.

- parameters
  - before: RFC3339 format (ex: 2021-01-10T13:28:00.337Z)
  - interval: millisecond
  - count: The number of data

### Example

The followings shows the response of `/histories?before=2021-02-09T03:54:13Z&interval=3600000&count=12`. It gets last `12` temperature and humidity data before `2021-02-09 03:54:13 UTC` at intervals of `3600000 millisecond (1 hour)`.

```
{
    "pages": {
        "total": 12
    },
    "data": [
        {
            "temperature": 23,
            "humidity": 12,
            "timestamp": "2021-02-09T15:54:13Z"
        },
        {
            "temperature": 22,
            "humidity": 12,
            "timestamp": "2021-02-09T14:54:06Z"
        },
        {
            "temperature": 25,
            "humidity": 6,
            "timestamp": "2021-02-09T13:52:13Z"
        },
        {
            "temperature": 24,
            "humidity": 13,
            "timestamp": "2021-02-09T12:54:11Z"
        },
        {
            "temperature": 21,
            "humidity": 12,
            "timestamp": "2021-02-09T11:54:13Z"
        },
        {
            "temperature": 23,
            "humidity": 15,
            "timestamp": "2021-02-09T10:54:13Z"
        },
        {
            "temperature": 23,
            "humidity": 12,
            "timestamp": "2021-02-09T09:53:53Z"
        },
        {
            "temperature": 24,
            "humidity": 3,
            "timestamp": "2021-02-09T08:53:35Z"
        },
        {
            "temperature": 23,
            "humidity": 5,
            "timestamp": "2021-02-09T07:54:06Z"
        },
        {
            "temperature": 24,
            "humidity": 5,
            "timestamp": "2021-02-09T06:53:51Z"
        },
        {
            "temperature": 26,
            "humidity": 4,
            "timestamp": "2021-02-09T05:54:16Z"
        },
        {
            "temperature": 24,
            "humidity": 5,
            "timestamp": "2021-02-09T04:54:16Z"
        }
    ]
}
```

# License

MIT