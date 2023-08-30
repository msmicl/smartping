package nettools

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/cihub/seelog"
)

func StartQperfAsServer() {
	cmd := exec.Command("qperf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(err.Error())
	}
	print(string(output))
	seelog.Info("qperf server started successfully.")
}

func installQperf() {
	cmd := exec.Command("sudo yum install -y qperf")
	_, err := cmd.CombinedOutput()
	if err != nil {
		seelog.Info("qperf install failed with %s\n", string(err.Error()))
		seelog.Info("Please install qperf manually.")
		os.Exit(1)
	}
}

func QperfPing(ipAddr string) (float64, error) {
	cmd := exec.Command("qperf", ipAddr, "tcp_lat")
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(string(err.Error()))
		return 0, err
	}
	result := string(output)
	index := strings.Index(result, "latency  =  ") + strings.Count("latency  =  ", "") - 1
	valueStr := result[index:]
	var latency float64 = 0
	var unit string = " us"
	if strings.Contains(valueStr, "ms") {
		unit = " ms"
	}
	temp := strings.TrimSpace(valueStr)
	index = strings.Index(temp, unit)
	if index == -1 {
		return 0, nil // connection failed.
	}
	temp = temp[:index]
	latency, err = strconv.ParseFloat(temp, 32)
	if err != nil {
		println(string(err.Error()))
		return 0, err
	}
	if unit == " us" {
		latency = latency / 1e3
	}
	seelog.Info(fmt.Sprintf("qperf ping %s latency %.2f ms", ipAddr, latency))
	return latency, nil
}
