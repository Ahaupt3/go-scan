package scan

import (
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	open         = 0
	finalResults []portResult
	wg           sync.WaitGroup
	mutex        sync.Mutex
)

type portResult struct {
	Port     int
	Protocol string
	State    string
}

func scanPort(protocol, ip string, port int) {
	defer wg.Done()
	result := portResult{Port: port, Protocol: strings.ToUpper(protocol)}

	address := ip + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 3*time.Second)

	if err != nil {
		return
	}
	defer conn.Close()

	result.State = "Open"
	finalResults = append(finalResults, result)
	return
}

func setScan(protocol, ip string, maxPort int) []portResult {
	wg.Add(maxPort)

	for port := 1; port <= maxPort; port++ {
		go scanPort(protocol, ip, port)
	}

	wg.Wait()
	return finalResults
}

// QuickScan - TCP scan of first 1024 ports
func QuickScan(ip string) {
	protocol := "tcp"
	maxPort := 1024

	resulting(setScan(protocol, ip, maxPort))
}

// FullScan - TCP scan of all 65,535 ports
func FullScan(ip string) {
	protocol := "tcp"
	maxPort := 65535

	resulting(setScan(protocol, ip, maxPort))
}

// UDPScan - UDP scan of all 65,535 ports
func UDPScan(ip string) {
	protocol := "udp"
	maxPort := 65535

	resulting(setScan(protocol, ip, maxPort))
}

func resulting(scanResults []portResult) {
	cleaned := sortResults(cleanResults(scanResults))

	for i := 0; i < len(cleaned); i++ {
		color.Green("Port: " + strconv.Itoa(cleaned[i].Port) + "/" + cleaned[i].Protocol + " - " + cleaned[i].State)
	}

	color.Yellow("Scan Complete!")
	// os.Exit(0)
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
