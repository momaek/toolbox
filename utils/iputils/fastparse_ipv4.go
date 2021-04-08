package iputils

import (
	"net"
)

// ParseIPv4 ..
func ParseIPv4(s string) net.IP {
	ip, ok := parseIPv4(s)
	if !ok {
		return nil
	}

	return ip[:]
}

func parseIPv4(s string) (ip [4]byte, ok bool) {
	var (
		dots int
		j    uint32
		l    = len(s)
	)

	for i := 0; i < l; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j = j*10 + uint32(s[i]-'0')
		case '.':
			if j > 255 {
				ok = false
				return
			}

			ip[dots] = byte(j)
			j = 0
			dots++
		default:
			ok = false
			return
		}
	}

	if j > 255 || dots != 3 {
		ok = false
		return
	}

	ip[dots] = byte(j)
	ok = true
	return
}
