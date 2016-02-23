package tcps

import (
	"fmt"
	"net"
	"time"

	"github.com/ckeyer/go-log"
)

const (
	ServerPort  = 8888
	ServerCount = 1
)

var (
	over    chan int
	servers []*Server
	logger  = log.GetDefaultLogger("ck")
)

func init() {
	over = make(chan int)
	servers = make([]*Server, ServerCount)
}

func Listen() {
	addr := &net.TCPAddr{Port: ServerPort}

	logger.Info("Running...")
	for i := 0; i < ServerCount; i++ {
		go func(index int) {
			listener, err := net.ListenTCP("tcp", addr)
			if err != nil {
				logger.Error(err)
				return
			}
			servers[index] = &Server{
				Name:     fmt.Sprint("Server", index),
				log:      log.GetDefaultLogger(fmt.Sprint("Server", index)),
				addr:     addr,
				listener: listener,
				Status:   SS_READY,
				cstop:    make(chan string),
			}
			go servers[index].ListenAndServe()
		}(i)
	}
}

func Over() {
	over <- 0
	select {
	case <-over:
		logger.Warning("stoped ...")
	case <-time.After(time.Second * 3):
		logger.Warning("stop force...")
	}
}
