package main

import (
	"RoommateScan/core/parsers"
	"RoommateScan/core/structure"
	"flag"
)

func init() {
	flag.StringVar(&structure.Hostname, "h", "", "Hostname to scan, example: 192.168.0.1-255\n192.168.0.1/24\n192.168.0.1\n192.168.0.1-192.168.1.255")
	flag.StringVar(&structure.Port, "p", "", "what port to scan")
	flag.StringVar(&structure.Proxy, "proxy", "", "proxy url")
}

func main() {
	flag.Parse()
	var host structure.Host
	host.Hostname = structure.Hostname
	host.Ports = structure.Port
	parsers.Parse(host)

}
