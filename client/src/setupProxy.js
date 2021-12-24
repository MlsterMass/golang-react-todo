const {createProxyMiddleware} = require("http-proxy-middleware");
module.exports = app => {
    app.use(
        createProxyMiddleware("/api/task",{
            target: "http://localhost:8888",
            changeOrigin:  true
        })
    )
}