package register

import (
	"net"
)

type Client struct {
	clientId   string
	clientName string
	netInfo    *netInfo
	connection net.Conn
	servers    []*Server
}

func (c *Client) GetUpdate() {
	size := len(c.servers)
	for i, v := range c.servers {
		if v.GetStatus() == OFFLINE {
			temp := c.servers[i]
			c.servers[i] = c.servers[size-1]
			c.servers[size-1] = temp
			size -= 1
		}
	}
	c.servers = c.servers[0:size]
}
