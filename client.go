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

	if !askIsHostTrusted(host, key) {
		return errors.New("add key error")
	}

	return goph.AddKnownHost(host, remote, key, "")
}

func askIsHostTrusted(host string, key ssh.PublicKey) bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Unknown Host: %s \nFingerprint: %s \n", host, ssh.FingerprintSHA256(key))
	fmt.Print("Would you like to add it? type yes or no: ")

	a, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.ToLower(strings.TrimSpace(a)) == "yes"
}
