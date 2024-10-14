# Run

```sh
$ go get .
$ go run .
```

Default database used to store data is `db.sqlite3`. To specify different file, use `--db` flag:

```sh
$ go run . --db db2.sqlite3
```

Default port is 8080. To specify a different one, use `--port` flag:

```sh
$ go run . --port 8081
```

Default host is `localhost`. To specify a different one, use `--host` flag:

```sh
$ go run . --host 0.0.0.0
```
