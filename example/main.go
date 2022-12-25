package main

import (
	"github.com/jwangace/gossh"
)

func main() {
	myhosts := gossh.Hosts{
		gossh.Host("192.168.0.1"),
		gossh.Host("192.168.0.2"),
		gossh.Host("192.168.0.3"),
	}
	// Run commands parallel on hosts
	myhosts.RunParallel("date")

	// Run commands on single host on condition
	if _, err := myhosts[0].Runcmd("date"); err != nil {
		myhosts[0].Runcmd("uptime")
	}
}
