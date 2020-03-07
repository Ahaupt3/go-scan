# go-scan
A simple port scanner built in Go.

## Install
```bash
git clone https://www.github.com/ahaupt3/go-scan/goroutines
cd go-scan
go install go-scan.go
```

## Usage
``` bash
go-scan <ip> <scan type: quick|full|udp>
```

## Scan Types
### Quick
Scans first 1024 TCP ports.

### Full
Scans all 65,535 TCP ports (issues w/ goroutines overwriting results).

### UDP
Scans all 65,535 UDP ports (issues w/ goroutines overwriting results).
