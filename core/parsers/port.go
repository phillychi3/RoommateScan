package parsers

import (
	"strconv"
	"strings"
)

func Port_parser(port string) []string {
	var portlist []string
	if port == "" {
		for i := 1; i <= 65535; i++ {
			portlist = append(portlist, strconv.Itoa(i))
		}
		return portlist
	}
	if strings.Contains(port, ",") {
		portlist = strings.Split(port, ",")
	} else if strings.Contains(port, "-") {
		stepport := strings.Split(port, "-")
		start, err := strconv.Atoi(stepport[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(stepport[1])
		if err != nil {
			panic(err)
		}
		for i := start; i <= end; i++ {
			portlist = append(portlist, strconv.Itoa(i))
		}

	} else {
		portlist = append(portlist, port)
	}
	return portlist
}
