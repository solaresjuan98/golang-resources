
## Testing in go

Before testing your code, you must init a module

```go
go mod init github.com/<username>/<module_name>
```

Then, type the following command to test your code:
```go
go test 
```

If you want to get the code coverage, write the following command:
```go
go test -cover
go test -coverprofile=coverage.out
```

If you want to get code coverage in a legible way, write the following command
```go
go tool cover -func=coverage.out
```

Get code coverage in HTML format
```go
go tool cover -html=coverage.out
```

