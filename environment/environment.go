package environment

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func Hostname() string {
	hostNameStr, err := os.Hostname()
	if err != nil {
		return "<unknown>"
	} else {
		return hostNameStr
	}
}

// Get preferred outbound ip of this machine
func OutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func AppPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return fmt.Sprint(dir), nil
}
