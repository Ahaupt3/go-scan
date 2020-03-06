package scan

import (
	"net"
	"strconv"
	"time"
)

var open = 0

type PortResult struct {
	Port  int
	Protocol string
	State string
}

func ScanPort(protocol, ip string, port int) PortResult {
	result := PortResult{Port: port, Protocol: strings.ToUpper(protocol)}

	address := ip + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 3*time.Second)
	
	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()

	result.State = "Open"
	return result
}

func QuickScan(ip string) []ScanResult {
	var results := []PortResult

	for port := 1; port <= 1024; port++ {
		ScanPort("tcp", ip, port)
		results = append(results, ScanPort("tcp", ip, port))
	}

	return results
}

// func FullScan() {

// }

// func UDPScan() {

// }

// func Results() {

// }