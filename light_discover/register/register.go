package register

import "time"

const SEND_TIMEOUT = 1000

type Register interface {
	RegisterServer(s *Server)
	OfflineServer(name string)
	DiscoverServer(name string)
}

type LightRegister struct {
	servers     map[string]*Server
	serversList []*Server
	checkTime   int64
	size        int
	onlineSize  int
	offlineSize int
	errorSize   int
}

type netInfo struct {
	ip         string
	port       int
	requestUrl string
	protocol   string
}

func (r *LightRegister) NewRegister() *LightRegister {
	return &LightRegister{
		servers:     make(map[string]*Server),
		serversList: make([]*Server, 0),
	}
}

func (r *LightRegister) CheckHeartTimeout() []*Server {
	if r.size == 0 {
		return nil
	}
	now := time.Now().UnixMilli()
	malfunctions := make([]*Server, 0)
	remain := int64(0)
	for _, v := range r.serversList {
		remain = now - v.lastUpdateAt
		if remain >= v.heartTime+SEND_TIMEOUT {
			switch v.status {
			case ONLINE:
				v.status = CONFIRMATION
				break
			case CONFIRMATION:
				v.status = MALFUNCTION
				malfunctions = append(malfunctions, v)
				break
			}
		}
	}
	return malfunctions
}

func (r *LightRegister) RegisterServer(s *Server) {

}

func (r *LightRegister) OfflineServer(name string) {

}

func (r *LightRegister) DiscoverServer(name string) {

}
