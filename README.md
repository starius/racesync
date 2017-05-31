```
$ go version
go version go1.8.3 linux/amd64
$ go get -race github.com/starius/racesync
# sync
.goroot/src/sync/cond.go:8: can't find import: "sync/atomic"
```

```
$ go get -t github.com/starius/racesync
$ go test -race github.com/starius/racesync/...
# net
can't create $WORK/net.a: open $WORK/net.a: no such file or directory
```

The issue: https://github.com/golang/go/issues/20512
