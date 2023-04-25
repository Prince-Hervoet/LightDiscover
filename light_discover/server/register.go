package server

import "time"

const SEND_TIMEOUT = 1000

type Register struct {
	servers     map[string]*Server
	serversList []*Server
	checkTime   int64
	size        int
	onlineSize  int
	offlineSize int
	errorSize   int
}

func (r *Register) NewRegister() {

}

func (r *Register) CheckHeartTimeout() []*Server {
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

func (r *Register) syncServerStatus() {

}
