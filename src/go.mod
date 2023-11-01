module smartping

go 1.21.0

replace (
	github.com/smartping/smartping v0.8.0 => ./
	github.com/msmicl/smartping/src/ethr => ./ethr
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
	github.com/smartping/smartping/src/nettools v0.0.0-00010101000000-000000000000
)

require (
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/lucas-clemente/quic-go v0.20.1 // indirect
	github.com/marten-seemann/qtls-go1-15 v0.1.4 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.3 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/msmicl/smartping/src/ethr v0.0.0-20231031005255-23546478a2f7 // indirect
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1 // indirect
	github.com/wcharczuk/go-chart v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/image v0.11.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)
