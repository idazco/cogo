package environment

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

// Hostname returns the host name recognized by the operating system
func Hostname() string {
	hostNameStr, err := os.Hostname()
	if err != nil {
		return "<unknown>"
	} else {
		return hostNameStr
	}
}

// OutboundIP gets the preferred outbound ip of the machine
func OutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// AppPath returns the absolute path to the running golang app
func AppPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return fmt.Sprint(dir), nil
}

// MacAddresses returns all the mac addresses as a map, with the name of the interface as the key
func MacAddresses() (result map[string]string) {
	interfaces, err := net.Interfaces()
	result = make(map[string]string)
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				result[i.Name] = i.HardwareAddr.String()
			}
		}
	}
	return
}
