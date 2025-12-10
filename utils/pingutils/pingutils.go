package pingutils

import (
	"os"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func Ping(addr string) (bool, error) {
	pinger, err := probing.NewPinger(addr)

	if err != nil {
		return false, err
	}

	pinger.SetPrivileged(os.Geteuid() == 0)
	pinger.Count = 3
	pinger.Timeout = 5 * time.Second
	
	err = pinger.Run()
	if err != nil {
		return false, err
	}

	stats := pinger.Statistics()

	if stats.PacketsRecv > 0 {
		return true, nil
	} else {
		return false, nil
	}
}