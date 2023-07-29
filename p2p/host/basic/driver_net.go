package basichost

import (
	"net"
)

type NativeNetDriver interface {
	InterfaceAddrs() ([]net.Addr, error)
	Interfaces() ([]net.Interface, error)
}
