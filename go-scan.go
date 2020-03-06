package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/ahaupt3/go-scan/scan"
	"github.com/fatih/color"
)

var ip = os.Args[1]
var scanType = strings.ToLower(os.Args[2])

func usage() {
	color.Red("Usage: " + os.Args[0] + " <ip> <scan type: quick|full|udp>")
	os.Exit(1)
}

func checkArgs() {
	if len(os.Args) < 3 {
		color.Red("Not enough arguments!")
		usage()
	} else if len(os.Args) > 3 {
		color.Red("Too many arguments!")
		usage()
	} else if net.ParseIP(ip) == nil {
		color.Red("Invalid IP!")
		usage()
	} else if scanType != "quick" && scanType != "full" && scanType != "udp" {
		color.Red("Invalid Scan Type!")
		usage()
	}
}

func main() {
	start := time.Now()
	checkArgs()

	color.Yellow("Scanning Ports...")

	if scanType == "quick" {
		scan.QuickScan(ip)
	} else if scanType == "full" {
		scan.FullScan(ip)
	} else if scanType == "udp" {
		scan.UDPScan(ip)
	} else {
		color.Red("Invalid Scan Type!")
		os.Exit(1)
	}
	elapsed := time.Since(start)
	log.Printf("go-scan took %s", elapsed)
}
