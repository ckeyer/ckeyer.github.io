package tcps

import (
	"net"

	"github.com/ckeyer/go-log"
)

const (
	BufferSize = 64
)
const (
	SS_READY   = "ready"
	SS_RUNNING = "running"
	SS_WAITING = "waiting"
	SS_CLOSED  = "closed"
)

type Controller struct {
	Name         string
	serves       []*Server
	cstop, ckill chan string
}

type Server struct {
	Name                 string
	log                  *log.Logger
	addr                 *net.TCPAddr
	listener             *net.TCPListener
	conns                []*Connect
	Status               string
	cstart, cstop, ckill chan string
}

func (s *Server) ListenAndServe() {
	s.log.Debug("start listenning...")
	if s.listener == nil {
		return
	}
	for {
		select {
		case ss := <-s.cstop:
			s.log.Debug("stoped", ss)
		case <-s.cstart:
			conn, err := s.listener.AcceptTCP()
			if err != nil {
				s.log.Error("accept err,", err)
				continue
			}
			s.log.Debug("connecting from: ", conn.RemoteAddr().String())

			myconn := &Connect{
				Name: s.Name + "--" + conn.RemoteAddr().String(),
				conn: conn,
			}

			s.conns = append(s.conns, myconn)

			go myconn.Serve(s.cstop)
		}
	}
}

type Connect struct {
	Name string
	conn net.Conn
}

func (c *Connect) Serve(sign chan string) {
	for {
		select {
		case <-over:
			over <- 2
			return
		default:
			bs := make([]byte, BufferSize)
			i, err := c.conn.Read(bs)
			if err != nil || i < 1 {
				continue
			}
			sign <- SS_WAITING

			logger.Notice(c.conn.RemoteAddr().String() + ": " + string(bs))
		}
	}
}
