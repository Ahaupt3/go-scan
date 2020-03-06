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

type portResult struct {
	Port     int
	Protocol string
	State    string
}

func scanPort(protocol, ip string, port int) portResult {
	result := portResult{Port: port, Protocol: strings.ToUpper(protocol)}

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

func setScan(protocol, ip string, maxPort int) []portResult {
	var results []portResult

	for port := 1; port <= maxPort; port++ {
		results = append(results, scanPort(protocol, ip, port))
	}
	return results
}

// QuickScan - TCP scan of first 1024 ports
func QuickScan(ip string) {
	protocol := "tcp"
	maxPort := 1024

	results(setScan(protocol, ip, maxPort))
}

// FullScan - TCP scan of all 65,535 ports
func FullScan(ip string) {
	protocol := "tcp"
	maxPort := 65535

	results(setScan(protocol, ip, maxPort))
}

// UDPScan - UDP scan of all 65,535 ports
func UDPScan(ip string) {
	protocol := "udp"
	maxPort := 65535

	results(setScan(protocol, ip, maxPort))
}

func results(scanresults []portResult) {
	cleaned := sortResults(cleanResults(scanresults))

	for i := 0; i < len(cleaned); i++ {
		color.Green("Port: " + strconv.Itoa(cleaned[i].Port) + "/" + cleaned[i].Protocol + " - " + cleaned[i].State)
	}

	color.Yellow("Scan Complete!")
	os.Exit(0)
}

func cleanResults(dirty []portResult) []portResult {
	var cleaned []portResult

	for i := 0; i < len(dirty); i++ {
		if dirty[i].State == "Open" {
			cleaned = append(cleaned, dirty[i])
			open++
		}
	}

	return cleaned
}

func sortResults(clean []portResult) []portResult {
	sort.Slice(clean, func(i, j int) bool {
		return clean[i].Port < clean[j].Port
	})

	return clean
}
