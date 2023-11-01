package funcs

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/smartping/smartping/src/g"
	"github.com/smartping/smartping/src/nettools"
)

func Ping() {
	var wg sync.WaitGroup
	for _, target := range g.SelfCfg.Ping {
		wg.Add(1)
		go PingTask(g.Cfg.Network[target], &wg)
	}
	wg.Wait()
	go StartAlert()
}

// ping main function
func PingTask(t g.NetworkMember, wg *sync.WaitGroup) {
	seelog.Info("Start Ping " + t.Addr + "..")
	stat := g.PingSt{}
	stat.MinDelay = -1
	// lossPK := 0
	ipaddr, err := net.ResolveIPAddr("ip", t.Addr)
	// var delay float64 = 0.0
	if err == nil {
		sent, rcvd, lost, avg, min, max := nettools.EthrPing(ipaddr.IP.String(), 8888)
		stat.SendPk = int(sent)
		stat.RevcPk = int(rcvd)
		stat.LossPk = int(lost)
		stat.AvgDelay = float64(avg.Nanoseconds()) / 1e6
		stat.MinDelay = float64(min.Nanoseconds()) / 1e6
		stat.MaxDelay = float64(max.Nanoseconds()) / 1e6
	} else {
		stat.AvgDelay = 0.00
		stat.MinDelay = 0.00
		stat.MaxDelay = 0.00
		stat.SendPk = 0
		stat.RevcPk = 0
		stat.LossPk = 100
		seelog.Debug("[func:qperf ping] Finish Addr:", t.Addr, " Unable to resolve destination host")
	}
	PingStorage(stat, t.Addr)
	wg.Done()
	seelog.Info("Finish Ping " + t.Addr + "..")
}

// storage ping data
func PingStorage(pingres g.PingSt, Addr string) {
	logtime := time.Now().Format("2006-01-02 15:04")
	seelog.Info("[func:StartPing] ", "(", logtime, ")Starting PingStorage ", Addr)
	sql := "INSERT INTO [pinglog] (logtime, target, maxdelay, mindelay, avgdelay, sendpk, revcpk, losspk) values('" + logtime + "','" + Addr + "','" + strconv.FormatFloat(pingres.MaxDelay, 'f', 2, 64) + "','" + strconv.FormatFloat(pingres.MinDelay, 'f', 2, 64) + "','" + strconv.FormatFloat(pingres.AvgDelay, 'f', 2, 64) + "','" + strconv.Itoa(pingres.SendPk) + "','" + strconv.Itoa(pingres.RevcPk) + "','" + strconv.Itoa(pingres.LossPk) + "')"
	seelog.Debug("[func:StartPing] ", sql)
	g.DLock.Lock()
	_, err := g.Db.Exec(sql)
	if err != nil {
		seelog.Error("[func:StartPing] Sql Error ", err)
	}
	g.DLock.Unlock()
	seelog.Info("[func:StartPing] ", "(", logtime, ") Finish PingStorage  ", Addr)
}
