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

# ROUTES

To list albums, do `GET /albums`:

```sh
$ curl localhost:8080/albums
```
Example of response:
```json
[
    {
        "id": 1,
        "title": "The Works",
        "artist": "Queen",
        "price": 59.99
    },
    {
        "id": 2,
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    },
    {
        "id": 3,
        "title": "The Beatles",
        "artist": "Beatles",
        "price": 29.99
    }
]  
```

To get paginated data, provide `limit` and `offset` URL params:

```sh
$ curl localhost:8080/albums?limit=2&offset=0
```
Example of response:
```json
[
    {
        "id": 1,
        "title": "The Works",
        "artist": "Queen",
        "price": 59.99
    },
    {
        "id": 2,
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    }
]  
```

To get single item, use `GET /albums/:id` pattern:

```sh
$ curl localhost:8080/albums/2
```
Example of response:
```json
{
    "id": 2,
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

To create new item, use `POST /albums` and provide `application/json` body:

```sh
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "A Kind of Magic", "artist": "Queen","price": 74.99}'
```
Example of response:
```http
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Mon, 14 Oct 2024 14:52:56 GMT
Content-Length: 115

{
    "id": 4,
    "title": "A Kind of Magic",
    "artist": "Queen",
    "price": 74.99
}
```
