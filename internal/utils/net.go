package net

import (
	"net"
	"net/netip"

	"github.com/free5gc/udm/internal/logger"
)

func BindingLookup(bindIP string) string {
	ips, err := net.LookupIP(bindIP)
	if err != nil {
		logger.InitLog.Errorf("Resolve BindingIP hostname %s failed: %+v", bindIP, err)
	}
	return ips[0].String()
}

func RegisterAddr(registerIP string) netip.Addr {
	ips, err := net.LookupIP(registerIP)
	if err != nil {
		logger.InitLog.Errorf("Resolve RegisterIP hostname %s failed: %+v", registerIP, err)
	}
	ip, err := netip.ParseAddr(ips[0].String())
	if err != nil {
		logger.InitLog.Errorf("Parse RegisterIP hostname %s failed: %+v", registerIP, err)
	}
	return ip
}
