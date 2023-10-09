package nettools

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/cihub/seelog"
)

func StartQperfAsServer() {
	cleanQPerfServer()
	cmd := exec.Command("qperf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(err.Error())
	}
	print(string(output))
	fmt.Println("qperf server started successfully.")
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
		println(string(err.Error()) + " " + string(output))
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

func cleanQPerfServer() {
	grep := exec.Command("grep", "qperf")
	ps := exec.Command("ps", "-a")

	// Get ps's stdout and attach it to grep's stdin.
	pipe, _ := ps.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe

	// Run ps first.
	ps.Start()

	// Run and get the output of grep.
	qperfres, err := grep.Output()
	if err != nil {
		fmt.Println("grep failed. " + string(err.Error()))
		return
	}
	fmt.Println(string(qperfres))
	lines := strings.Split(string(qperfres), "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		cmdtext := line[len(line)-5:]
		if cmdtext == "qperf" {
			pid := strings.Split(line, " ")[0]
			if len(pid) > 0 {
				killQperf(pid)
			}
		}
	}
}

func killQperf(pid string) {
	cmd := exec.Command("sudo", "kill", pid)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Kill " + pid + " error: " + string(err.Error()))
	} else {
		fmt.Println("Kill qperf succeeded.")
	}
}

func CheckQperfStatus() {

	for {
		time.Sleep(10 * time.Second)
		cmd := exec.Command("qperf", "127.0.0.1", "tcp_lat")
		output, _ := cmd.CombinedOutput()
		if strings.Contains(string(output), "failed to connect") {
			fmt.Println("qperf server is not running correctly, restart it now.")
			go StartQperfAsServer()
		}
	}
}
