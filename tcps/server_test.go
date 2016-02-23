package tcps

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func init() {
	go Listen()
	time.Sleep(time.Second)
}

func TestSendMsg(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := SendMsg("msg")
		if err != nil {
			t.Error(err)
		}
		time.Sleep(1 * time.Second)
	}
	Over()
}

type Client struct {
	cli *net.Conn
}

func SendMsg(msg string) error {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: ServerPort})
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("send err, ", err)
	}
	conn.CloseWrite()
	return err
}
