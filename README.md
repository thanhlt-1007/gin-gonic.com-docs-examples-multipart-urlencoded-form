# gin-gonic.com-docs-examples-multipart-urlencoded-form

- Multipart/Urlencoded form

- Reference: https://gin-gonic.com/docs/examples/multipart-urlencoded-form/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

- Default message

```sh
curl --location 'localhost:8080/message' \
--form 'message="Hello"'
```

- Message to admin

```sh
curl --location 'localhost:8080/message' \
--form 'message="Hello"' \
--form 'name="admin"'
```
