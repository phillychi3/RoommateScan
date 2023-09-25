package scan

import (
	"fmt"

	probing "github.com/prometheus-community/pro-bing"
)

func Check_alive(ips []string) {
	for _, ip := range ips {
		pinger, err := probing.NewPinger(ip)
		pinger.SetPrivileged(true)
		pinger.Timeout = 2000
		if err != nil {
			fmt.Println("wowERROR:", err)
		}
		err = pinger.Run()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if pinger.Statistics().PacketsRecv > 0 {
				fmt.Println(ip, "is alive")
			} else {
				fmt.Println(ip, "is dead")
			}
		}

	}
}
