# thermohygrometer

# API

## GET /current

```
{
    "temperature":22,
    "humidity":15,
    "timestamp":"2021-01-10T13:25:48Z"
}
```

## GET /histories

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