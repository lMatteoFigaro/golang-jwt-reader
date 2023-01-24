# golang-jwt-reader
Golang jwt reader

Poc to validate jwt token in go with RSA signing

Start with 

```bash
go run main.go
```

It starts on localhost:8090 and validate token given in header. 

Example request 

```bash
curl --location --request GET 'http://localhost:8090/authentication' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.jFC5vYe50WBFsCyx5syiVHWbS1cqmt1PPW8NQuYJTR9VVxKxeT8hZ7lDWDgjyvKTL3h3ZL9w8xLi6onfea64uZOz5Z53GlP6aS2Atevu-AMbtL5Jw17V_kcPjBqBqCg6G2qQfEm3astUbKMXfrbakNMUuHj0P5z8WjXBP7pp8VQin9ixH5zP7wZLxO_4E1bgXWY8ggeGIn5coaiQh8KJ0SwAN8cnh4PfOw8O5McIEYou6UH5w1UjJzKzA_TuZJReJ8RvcrKp4b6MQqAP8ob2INXn5rZ_qZmgI7OAlUuaDIvvLQtOWEzEfAU12dun81qyNlA8RdAkWQnG4vr0iFv3LA'
```


