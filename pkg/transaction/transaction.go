package transaction

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/Wishrem/wuso/client/screen"
	"github.com/Wishrem/wuso/pkg/pool"
	prot "github.com/Wishrem/wuso/pkg/protocol"
	"github.com/gorilla/websocket"
)

type tx struct {
	s        *sender
	recvFunc func(p pool.Param)
	warpData func(data []byte) pool.Param
	recvPool *pool.Pool
}

func New(workers int, recvFunc func(p pool.Param), warpData func([]byte) pool.Param, recvPool *pool.Pool) *tx {
	return &tx{
		s:        newSender(workers),
		recvFunc: recvFunc,
		warpData: warpData,
		recvPool: recvPool,
	}
}

func (tx *tx) Start(ok chan bool) {
	if ws == nil {
		u := &url.URL{
			Scheme: "ws",
			Host:   "localhost:8080",
			Path:   "/wuso/ws",
		}
		ws = &_websocket{
			dialer: &websocket.Dialer{},
			url:    u.String(),
		}
	}

	if ws.conn != nil {
		ok <- reconnect(3)
		return
	}

	if !connect() {
		log.Println("failed to connecting")
		screen.Notice("错误", "暂时无法连接至服务器", screen.NoticeWrong)
		ok <- false
		return
	}
	defer ws.conn.Close()
	go recvMsg()

	intr := make(chan os.Signal, 3)
	signal.Notify(intr, os.Interrupt)

	ok <- true
	for {
		select {
		case <-intr:
			Close()
			return
		case msg := <-ws.recvBytes:
			if len(msg) < 1 {
				// TODO
				continue
			}

			// TODO
			dt, _, msg, err := prot.Decode(msg)
			if err != nil {
				log.Println("decoding received message error:", err)
				// Do Nothing
				continue
			}
			if dt^prot.CMDType == 0 {

			} else if dt^prot.MsgType == 0 {
				tx.recvPool.AddJob(pool.NewJob(tx.warpData(msg), tx.recvFunc))
			} else {
				log.Println("received a msg with the wrong data type, msg:", msg)
				// Do Nothing
			}
		}
	}
}
