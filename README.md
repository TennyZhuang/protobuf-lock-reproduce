Related to

* https://github.com/golang/protobuf/pull/1004
* https://github.com/golang/protobuf/issues/888

Environment

* protoc: libprotoc 3.11.2
* protoc-gen-go: 1.3.2

Generate test.pb.go

```
protoc --go_out=paths=source_relative:. test.proto
```

Run benchmark

```
go test -bench=. -benchtime=20s
```

A test result

```
goos: darwin
goarch: amd64
pkg: test_proto
BenchmarkGolangProto-4   	   19132	   1226288 ns/op
--- BENCH: BenchmarkGolangProto-4
    proto_test.go:43: max latency 2.115446ms
    proto_test.go:43: max latency 107.731387ms
    proto_test.go:43: max latency 12.523745369s
    proto_test.go:43: max latency 23.421683369s
BenchmarkGoGoProto-4     	   21175	   1370062 ns/op
--- BENCH: BenchmarkGoGoProto-4
    proto_test.go:78: max latency 2.871029ms
    proto_test.go:78: max latency 108.636449ms
    proto_test.go:78: max latency 11.31063355s
    proto_test.go:78: max latency 29.000715051s
PASS
ok  	test_proto	76.914s
```

Please focus on max latency output.
