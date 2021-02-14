UI App for Thermohygrometer

<img width="512" alt="thermohygrometer" src="https://raw.githubusercontent.com/sano-home/thermohygrometer/main/thermohygrometer.png">

# Setup

```
npm install
```

# Build

```
npm run build
```

# Run

```
THERMOHYGROMETER_API_URL=<API Url> THERMOHYGROMETER_UI_PORT=<UI App Port> npm start
```

# Development

- Run UI app with mock data (`./pages/api/*`) in development mode

```
npm run mock
```

- Run UI app with mock data in production mode

```
npm run build
npm run start:mock
```
