# microURL
The one URL shortener to rule them all

## Usage
1. Run the migrations using [goose](https://github.com/pressly/goose)

SQLite:
```shell
goose -dir migrations/sqlite/ sqlite3 ./microURL.sqlite3 up
```
2. Start the backend
```shell
go run ./cmd/app/main.go
```
By default, it runs on port 8080

3. Host the frontend using [caddy](https://caddyserver.com/):
```shell
caddy run -c ./config/Caddyfile
```
You can access the frontend at [http://127.0.0.1:8086]()
