package utility

import (
	"errors"
	"net"

	log "github.com/sirupsen/logrus"
)

func ResolveHostIpV4() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		errHandle(err)
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIp, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {

			ip := networkIp.IP.String()
			return ip
		}
	}
	errHandle(errors.New("IP Not Found"))
	return ""
}

func errHandle(err error) {
	log.Fatal("Resolved Host IP Error :", err)
}
