Example golang api-webapp structure for web project
===

Origin example

Prerequisits:
===

1. Linux + Golang 1.1X.X
2. Free port on 127.0.0.1:3000

Run webapp:
===

0. In terminal: go mod tidy
1. Setup entry point of your webservice in file cmd/server/main.go:entryPoint
1. Setup sqlite3 libs for using the "apidb.sqlite3"
2. Import db_init_sqlite.sql if you database is empty
3. Run in terminal:

```bash
chmod +x run.sh
./run.sh
```
