# microURL
The one URL shortener to rule them all

## Usage
1. Host the frontend using [caddy](https://caddyserver.com/):
```shell
caddy run -c ./config/Caddyfile
```
You can access the frontend on [http:127.0.0.1:8086]()
2. Start the backend
```shell
go run microURL
```
By default, it runs on port 8080