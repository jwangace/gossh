package main

import "github.com/jwangace/gossh"

func main() {
	myhosts := gossh.Hosts{
		gossh.Host("192.168.0.1"),
		gossh.Host("192.168.0.2"),
		gossh.Host("192.168.0.3"),
	}
	myhosts.RunParallel("date")
}
