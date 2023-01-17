package zerolog

import (
	"fmt"
	"net"
)

type Printer struct{}

func (_ Printer) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return 0, nil
}

type UdpSender struct {
	udp net.Conn
}

func (u UdpSender) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return u.udp.Write(p)
}

func NewUdpSender(url string) UdpSender {
	con, err := net.Dial("udp", url)
	if err != nil {
		panic(err)
	}

	return UdpSender{
		udp: con,
	}
}
