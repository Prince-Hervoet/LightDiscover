package register

import (
	"net"
	"time"
)

const (
	ONLINE       = 1
	OFFLINE      = 0
	MALFUNCTION  = -1
	CONFIRMATION = 2
)

type Server struct {
	serverId     string
	serverName   string
	netInfo      *netInfo
	connection   net.Conn
	version      string
	status       int
	registerAt   int64
	heartTime    int64
	lastUpdateAt int64
	metaInfo     map[string]string
}

func NewServer(id, name string) *Server {
	now := time.Now().UnixMilli()
	return &Server{
		serverId:     id,
		serverName:   name,
		registerAt:   now,
		lastUpdateAt: now,
		status:       CONFIRMATION,
		metaInfo:     nil,
	}
}

func (s *Server) SetNetInfo(ip string, port int, url string, protocol string) {
	n := &netInfo{
		ip:         ip,
		port:       port,
		protocol:   protocol,
		requestUrl: url,
	}
	s.netInfo = n
}

func (s *Server) SetVersion(version string) {
	s.version = version
}

func (s *Server) SetHeartTime(time int64) {
	s.heartTime = time
}

func (s *Server) SetMetaInfo(info map[string]string) {
	s.metaInfo = info
}

func (s *Server) SendHeart() {

}

func (s *Server) GetStatus() int {
	return s.status
}
