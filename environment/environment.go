package environment

import "os"

func Hostname() string {
	hostNameStr, err := os.Hostname()
	if err != nil {
		return "<unknown>"
	} else {
		return hostNameStr
	}
}
