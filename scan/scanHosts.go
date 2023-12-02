package scan

import (
	"fmt"
	"net"
	"time"
)

type Result struct {
	Host       string
	NotFound   bool
	PortStates []PortState
}

type PortState struct {
	Port int
	Open state
}

type state bool

func (s state) String() string {
	if s {
		return "open"
	}

	return "closed"
}

// scanPort performs a port scan on a single TCP port
func scanPort(host string, port int) PortState {
	p := PortState{Port: port}
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	scanConn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return p
	}

	if err := scanConn.Close(); err != nil {
		return p
	}

	p.Open = true

	return p
}

// Run peroforms a port scan on the hosts list
func Run(hl *HostsList, ports []int) []Result {
	res := make([]Result, 0, len(hl.Hosts))

	for _, h := range hl.Hosts {
		r := Result{
			Host: h,
		}
		if _, err := net.LookupHost(h); err != nil {
			r.NotFound = true
			res = append(res, r)

			continue
		}

		for _, p := range ports {
			r.PortStates = append(r.PortStates, scanPort(h, p))
		}

		res = append(res, r)
	}

	return res
}
