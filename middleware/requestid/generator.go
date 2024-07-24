package requestid

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	format  = formatRawIp(getLocalIP()) + "%03x-%x-" + fmt.Sprintf("%x", os.Getpid())
	counter uint64
	mu      sync.Mutex
)

// RequestIdGenerator generates a request id,
// The request id is composed of 4 parts
//
// all values are encode in hex: "[ip 8bit][counter 3bit]-[timestamp (ms)]-[pid]"
func RequestIdGenerator() string {
	mu.Lock()
	defer mu.Unlock()

	counter = (counter + 1) % 1000
	timestamp := time.Now().UnixMilli()

	return fmt.Sprintf(format, counter, timestamp)
}

func formatRawIp(ip string) string {
	rawIp := ""
	for _, p := range strings.Split(ip, ".") {
		if n, err := strconv.Atoi(p); err != nil {
			panic(err)
		} else {
			rawIp += fmt.Sprintf("%02x", n)
		}
	}

	return rawIp
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, address := range addrs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
