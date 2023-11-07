package nettools

import (
	"time"

	"github.com/msmicl/smartping/src/ethr"
	"github.com/smartping/smartping/src/g"
)

func StartEthrServer() {
	var ethrPort = g.Cfg.EthrPort
	ethr.StartEthrService(ethrPort)
}

func EthrPing(ip string, port int) (uint32, uint32, uint32, time.Duration, time.Duration, time.Duration) {
	return ethr.EthrPing(ip, port)
}
