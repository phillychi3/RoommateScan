package structure

type Host struct {
	Hostname string // Hostname 網址或者IP
	IP       string // 轉換後的IP
	Ports    string
}

var (
	Hostname string
	Port     string
	Proxy    string
)
