package parsers

import (
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type Ipaddress struct {
	ip1  int
	ip2  int
	ip3  int
	ip4  int
	all  string
	isip bool
}

// 輸入IP，輸出IP列表
// 例如
// 192.168.0.1/24,
// 192.168.0.1-255
// 192,168.0.1,192,168.0.2
func Ip_parser(ip string) []string {
	var iplist []string
	if strings.Contains(ip, ",") {
		iplist = strings.Split(ip, ",")
	} else if strings.Contains(ip, "-") {
		start := strings.Split(ip, "-")[0]
		end := strings.Split(ip, "-")[1]
		if isip(start) {
			iplist = ipcount(iptoipaddress(start), iptoipaddress(end))
		} else {
			panic("ip format error")
		}

	}
	return iplist

}

func ipcount(ips Ipaddress, ipe Ipaddress) []string {
	// example 10.0.0.1-10.0.0.255
	var iplist []string
	//沒想法現在
	if !ipe.isip {
		ipe.ip1 = ips.ip1
		ipe.ip2 = ips.ip2
		ipe.ip3 = ips.ip3
		ipe.all = strconv.Itoa(ips.ip1) + "." + strconv.Itoa(ips.ip2) + "." + strconv.Itoa(ips.ip3) + "." + strconv.Itoa(ipe.ip4)
	}
	fmt.Print(ips, ipe)

	start := ip2int(net.ParseIP(ips.all))
	end := ip2int(net.ParseIP(ipe.all))
	for i := start; i <= end; i++ {
		iplist = append(iplist, int2ip(i).String())
	}
	return iplist

}

func iptoipaddress(ip string) Ipaddress {
	var ipaddress Ipaddress
	if isip(ip) {
		ipaddress.ip1, _ = strconv.Atoi(strings.Split(ip, ".")[0])
		ipaddress.ip2, _ = strconv.Atoi(strings.Split(ip, ".")[1])
		ipaddress.ip3, _ = strconv.Atoi(strings.Split(ip, ".")[2])
		ipaddress.ip4, _ = strconv.Atoi(strings.Split(ip, ".")[3])
		ipaddress.isip = true
	} else {
		ipaddress.ip1 = 0
		ipaddress.ip2 = 0
		ipaddress.ip3 = 0
		ipaddress.ip4, _ = strconv.Atoi(ip)
		ipaddress.isip = false
	}
	ipaddress.all = ip
	return ipaddress
}

func isip(ip string) bool {
	if match, _ := regexp.MatchString(`(\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`, ip); match {
		return true
	} else {
		return false
	}
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
