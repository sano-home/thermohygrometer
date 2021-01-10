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

```
{
    "temperature":22,
    "humidity":15,
    "timestamp":"2021-01-10T13:25:48Z"
}
```

## GET /histories?since=YYYY-MM-DDThh:mm:ss.SSSZ&before=YYYY-MM-DDThh:mm:ss.SSSZ

Get temperature and humidity history between `since` and `before`.

- parameters
  - since: RFC3339 format (ex: 2021-01-10T13:27:00.337Z)
  - before: RFC3339 format (ex: 2021-01-10T13:28:00.337Z)

```
{
  "pages": {
    "total": 6
  },
  "data": [
    {
      "temperature": 22,
      "humidity": 14,
      "timestamp": "2021-01-10T13:27:10Z"
    },
    {
      "temperature": 22,
      "humidity": 14,
      "timestamp": "2021-01-10T13:27:20Z"
    },
    {
      "temperature": 22,
      "humidity": 14,
      "timestamp": "2021-01-10T13:27:30Z"
    },
    {
      "temperature": 22,
      "humidity": 14,
      "timestamp": "2021-01-10T13:27:40Z"
    },
    {
      "temperature": 22,
      "humidity": 13,
      "timestamp": "2021-01-10T13:27:50Z"
    },
    {
      "temperature": 22,
      "humidity": 15,
      "timestamp": "2021-01-10T13:28:00Z"
    }
  ]
}
```

# License

MIT