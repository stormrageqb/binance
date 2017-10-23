package binance

import (
	"github.com/gorilla/websocket"
)

type WsHandler func(message []byte)

type wsConfig struct {
	endpoint string
}

func newWsConfig(endpoint string) *wsConfig {
	return &wsConfig{
		endpoint: endpoint,
	}
}

func wsServe(cfg *wsConfig, handler WsHandler) (done chan struct{}, err error) {
	c, _, err := websocket.DefaultDialer.Dial(cfg.endpoint, nil)
	if err != nil {
		return
	}
	done = make(chan struct{})
	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				return
			}
			go handler(message)
		}
	}()
	return
}
