# port-scan-go
Port Scanner built in Go

## Install
```bash
git clone https://www.github.com/ahaupt3/port-scan-go/master
cd port-scan-go
go install port-scan-go.go
```

## Usage
``` bash
port-scan-go <ip> <scan type: ping|quick|full>
```

## Scan Types
### Quick
Scans first 1024 TCP ports.

### Full
Scans all 65,535 TCP ports.
### UDP
Scans first 65,535 UDP ports.
