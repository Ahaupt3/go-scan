package main

import (
	"net"
	"os"
	"strings"

	"github.com/ahaupt3/go-scan/scan"
	"github.com/fatih/color"
)

var ip = os.Args[1]
var scanType = strings.ToLower(os.Args[2])

func Usage() {
	color.Red("Usage: " + os.Args[0] + " <ip> <scan type: quick|full|udp>")
	os.Exit(1)
}

func CheckArgs() {
	if len(os.Args) < 3 {
		color.Red("Not enough arguments!")
		Usage()
	} else if len(os.Args) > 3 {
		color.Red("Too many arguments!")
		Usage()
	} else if net.ParseIP(ip) == nil {
		color.Red("Invalid IP!")
		Usage()
	} else if scanType != "quick" && scanType != "full" && scanType != "udp" {
		color.Red("Invalid Scan Type!")
		Usage()
	}
}

func main() {
	CheckArgs()

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
}
