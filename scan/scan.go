package scan

import (
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var open = 0

type PortResult struct {
	Port     int
	Protocol string
	State    string
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

func QuickScan(ip string) {
	var results []PortResult

	for port := 1; port <= 1024; port++ {
		results = append(results, ScanPort("tcp", ip, port))
	}

	Results(results)
}

func FullScan(ip string) {
	var results []PortResult

	for port := 1; port <= 65535; port++ {
		results = append(results, ScanPort("tcp", ip, port))
	}

	Results(results)
}

func UDPScan(ip string) {
	var results []PortResult

	for port := 1; port <= 65535; port++ {
		results = append(results, ScanPort("udp", ip, port))
	}

	Results(results)
}

func Results(scanresults []PortResult) {
	cleaned := SortResults(CleanResults(scanresults))

	for i := 0; i < len(cleaned); i++ {
		color.Green("Port: " + strconv.Itoa(cleaned[i].Port) + "/" + cleaned[i].Protocol + " - " + cleaned[i].State)
	}

	color.Yellow("Scan Complete!")
	os.Exit(0)
}

func CleanResults(dirty []PortResult) []PortResult {
	var cleaned []PortResult

	for i := 0; i < len(dirty); i++ {
		if dirty[i].State == "Open" {
			cleaned = append(cleaned, dirty[i])
			open++
		}
	}

	return cleaned
}

func SortResults(clean []PortResult) []PortResult {
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Port < clean[j].Port
	})

	return clean
}
