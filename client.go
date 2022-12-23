package gossh

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"strings"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

type Host string

func (h Host) Sshclient() (*goph.Client, error) {
	auth, err := goph.UseAgent()
	if err != nil {
		log.Fatal(err)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return goph.New(user.Username, string(h), auth)
}

func (h Host) Runcmd(s string) string {
	client, err := h.Sshclient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	out, _ := client.Run(s)
	return string(out)
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
