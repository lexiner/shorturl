package util

import (
	"fmt"
	"net"

	"github.com/astaxie/beego"
)

type Helper struct {
	beego.Controller
}

func (h *Helper) GetConfig(key string) string {
	return beego.AppConfig.String(key)
}

// ip to long
func (h *Helper) GetIpToLong() (uint32, error) {
	ip := h.Ctx.Input.IP()
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return 0, fmt.Errorf("invalid IP address format")
	}
	ipv4 := parsedIP.To4()
	if ipv4 == nil {
		return 0, fmt.Errorf("not an IPv4 address")
	}
	return uint32(ipv4[0])<<24 | uint32(ipv4[1])<<16 | uint32(ipv4[2])<<8 | uint32(ipv4[3]), nil
}
