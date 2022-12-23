package main

import "github.com/jwangace/gossh"

var mmp = map[string]string{
	"disablePuppet": "sudo puppet agent --disable",
}

type host struct {
	IP string
}

func (h *host) getHost() gossh.Host {
	return gossh.Host(h.IP)
}

func (h *host) disablePuppet() {
	h.getHost().Runcmd(mmp["disablePuppet"])
}

func main() {
	myhost := host{
		IP: "192.168.0.1",
	}

	myhost.disablePuppet()
}
