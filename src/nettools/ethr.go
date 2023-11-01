package nettools

import (
	"time"

	"github.com/msmicl/smartping/src/ethr"
)

func StartEthrServer() {
	ethr.StartEthrService(8888)
}

func EthrPing(ip string, port int) (uint32, uint32, uint32, time.Duration, time.Duration, time.Duration) {
	return ethr.EthrPing(ip, port)
}

func EthrProb(port int) bool {
	return true
}
