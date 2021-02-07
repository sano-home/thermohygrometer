const express = require('express')
const next = require('next')
const { createProxyMiddleware } = require('http-proxy-middleware')

const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handle = app.getRequestHandler()

const host = '0.0.0.0';
const port = parseInt(process.env.PORT, 10) || 7000;

app.prepare().then(() => {
  const server = express()

  server.use(
    '/api',
    createProxyMiddleware({
      target: 'http://192.168.86.29:8000/', // FIXME
      changeOrigin: true,
      pathRewrite: {'^/api/' : '/'}
    })
  );

  server.all('*', (req, res) => {
    return handle(req, res)
  })

  server.listen(port, host, err => {
    if (err) throw err
    console.log(`> Ready on ${host}:${port}`)
  })
})
