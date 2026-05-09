# Todo Server
This is a API server for Todo app in go

## Dependencies
- `Gin` Web framework

## Running
```shell
$ go get
$ go build
$ ./todo-server
```

It runs in `8080` port.

## Tasks to do
- Implement repository for db (PostgreSQL - Docker)
- Seperate http request handler to `handler` package
- Authentication (JWT, OAuth2) & Authorization
- Proper validation of fields
- Custom error messages
- Consistent response json format
- Fix bugs in getOneTodo
    - Better error message
    - When integer overflows: better error
- Fix bugs in createTodo
    - Create ignores invalid fields. Need better error
- Fix bugs in updateTodo
    - Update ignores invalid fields. Need better error
- Fix bugs in deleteTodo
    - When specifid id as character: better error message
- Write tests in Go tests (net/http/httptest)
- Add multiple roles
- Add authentication