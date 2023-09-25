package parsers

import (
	"RoommateScan/core/scan"
	"RoommateScan/core/structure"
	"fmt"
)

func Parse(host structure.Host) {
	prots := Port_parser(host.Ports)
	ips := Ip_parser(host.Hostname)
	fmt.Print(prots, ips)
	scan.Check_alive(ips)

}
