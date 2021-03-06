package tcp_estats

import "net"

type Address struct {
	IsIPv4 bool
	addr   uint32
	addr6  uint64
}

func NewIPv4(ip uint32) *Address {
	a := new(Address)
	a.IsIPv4 = true
	a.addr = ip
	return a
}

func NewIPv6(ip uint64) *Address {
	a := new(Address)
	a.IsIPv4 = false
	a.addr6 = ip
	return a
}

func (a Address) GetIPv4() uint32 {
	if !a.IsIPv4 {
		panic("not IPv4")
	}
	return a.addr
}

func (a Address) GetIPv6() uint64 {
	if a.IsIPv4 {
		panic("not IPv6")
	}
	return a.addr6
}

// assumes ipv4
func intToIP(num uint32) net.IP {
	ip := make(net.IP, 4)
	native.PutUint32(ip, num)
	return ip
}

// assumes ipv4
func ipToInt(ip net.IP) uint32 {
	return native.Uint32(ip.To4())
}
