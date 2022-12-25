package gossh

import (
	"fmt"
	"log"
	"net"
	"os/user"
	"sync"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

type Host string
type Hosts []Host

func (h Host) Runcmd(cmd string) (string, error) {
	client, err := h.Sshclient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	out, err := client.Run(cmd)
	return string(out), err
}

func (hs Hosts) RunParallel(cmd string) {
	var wg sync.WaitGroup
	for _, h := range hs {
		wg.Add(1)
		h := h
		go func() {
			out, err := h.Runcmd(cmd)
			if err != nil {
				fmt.Print(h, "Error:", err)
			} else {
				fmt.Print(h, ":", string(out))
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func (h Host) Sshclient() (*goph.Client, error) {
	auth, err := goph.UseAgent()
	if err != nil {
		log.Fatal(err)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return goph.NewConn(&goph.Config{
		User:     user.Username,
		Addr:     string(h),
		Auth:     auth,
		Port:     22,
		Callback: VerifyHost,
	})
}

func VerifyHost(host string, remote net.Addr, key ssh.PublicKey) error {
	hostFound, err := goph.CheckKnownHost(host, remote, key, "")
	if hostFound && err != nil {
		return err
	}
	if hostFound && err == nil {
		return nil
	}
	return goph.AddKnownHost(host, remote, key, "")
}
