package address

import (
	"errors"
	"log"
	"net"
)

// ResolveHostIPV4 ...
func ResolveHostIPV4() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		errHandle(err)
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIP, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {

			ip := networkIP.IP.String()
			return ip
		}
	}
	errHandle(errors.New("IP Not Found"))
	return ""
}

func errHandle(err error) {
	log.Fatal("Resolved Host IP Error :", err)
}
