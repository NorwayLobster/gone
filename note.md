<!--
 * @Date: 2022-02-10 11:18:45
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-10 21:14:38
 * @FilePath: /gone/note.md
-->
go test -mod=vendor ./... -cover -coverprofile out.cover && 
go tool cover -html=out.cover

find . -type f \
    -name '*.go' \
    -exec sed -i -e 's,github.com/my/project,github.com/my/project/v2,g' {} \;

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
/usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkDisabledWithoutFields$ go.uber.org/zap/benchmarks
go tool pprof -svg cpu.prof
go help testflag
go help test
go tool pprof -h
gofmt -l -w .
go doc gofmt