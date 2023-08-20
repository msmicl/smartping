module smartping

go 1.21.0

replace (
	github.com/smartping/smartping v0.8.0 => ./
	github.com/smartping/smartping/src/funcs => ./funcs
	github.com/smartping/smartping/src/g => ./g
	github.com/smartping/smartping/src/http => ./http
	github.com/smartping/smartping/src/logs => ./logs
	github.com/smartping/smartping/src/nettools => ./nettools
)

require (
	github.com/jakecoffman/cron v0.0.0-20190106200828-7e2009c226a5
	github.com/smartping/smartping/src/funcs v0.0.0-00010101000000-000000000000
	github.com/smartping/smartping/src/g v0.0.0-00010101000000-000000000000
	github.com/smartping/smartping/src/http v0.0.0-00010101000000-000000000000
)

require (
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/smartping/smartping/src/nettools v0.0.0-00010101000000-000000000000 // indirect
	github.com/wcharczuk/go-chart v2.0.1+incompatible // indirect
	golang.org/x/image v0.11.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
)
