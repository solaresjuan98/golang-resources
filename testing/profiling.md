## Profiling in go


```go
go test -cpuprofile=cpu.out
```

```go
go tool pprof cpu.out
```

## List of useful pprof commands

* Type command top to show the percentage of time used in the tested code.

* List command shows you a more detailed information about the duration of the execution in the code.

* Web command shows a code execution graph (dot must be installed).